package chain

import (
	"errors"
	"fmt"
	"time"

	"github.com/MysGate/demo_backend/core"
	"github.com/MysGate/demo_backend/model"
	"github.com/MysGate/demo_backend/util"
	"github.com/MysGate/demo_backend/zkp"
)

func (cm *ChainManager) findChainHandler(srcChainId uint64) *ChainHandler {
	handler, ok := cm.handlers[srcChainId]
	if !ok || handler == nil {
		errMsg := fmt.Sprintf("findChainHandler fail, srcChainId:%+v", srcChainId)
		util.Logger().Error(errMsg)
		return nil
	}
	return handler
}

func (cm *ChainManager) findDestHandler(srcChainId, destChainId uint64) *DestChainHandler {
	handler := cm.findChainHandler(srcChainId)
	if handler == nil {
		errMsg := fmt.Sprintf("findDestHandler fail, srcChainId:%+v", srcChainId)
		util.Logger().Error(errMsg)
		return nil
	}

	destHandler, ok := handler.dest[destChainId]
	if !ok || destHandler == nil {
		errMsg := fmt.Sprintf("findDestHandler fail, destChainId:%+v", destChainId)
		util.Logger().Error(errMsg)
		return nil
	}

	return destHandler
}

func (cm *ChainManager) findSrcHandler(srcChainId uint64) *SrcChainHandler {
	handler := cm.findChainHandler(srcChainId)
	if handler == nil {
		errMsg := fmt.Sprintf("findSrcHandler fail, srcChainId:%+v", srcChainId)
		util.Logger().Error(errMsg)
		return nil
	}

	return handler.src
}

func (cm *ChainManager) handlePayForDest(order *model.Order) error {
	destHandler := cm.findDestHandler(order.SrcChainId, order.DestChainId)
	ret, err := destHandler.crossFrom(order)
	if err != nil {
		errMsg := fmt.Sprintf("handlePayForDest crossFrom err:%+v", err)
		util.Logger().Error(errMsg)
		return err
	}

	if !ret {
		errMsg := fmt.Sprintf("handlePayForDest crossFrom trx failed, tx:", order.DestTxHash)
		util.Logger().Error(errMsg)
		return err
	}

	err = model.UpdateOrderStatus(&model.Order{ID: order.ID, DestTxHash: order.DestTxHash, DestTxStatus: 1, Status: core.CrossFrom}, cm.db)
	if err != nil {
		util.Logger().Error(fmt.Sprintf("handlePayForDest update db err:%+v", err))
		return err
	}
	if cm.cfg.ZkVerify.Enable {
		return cm.SendMessage(add_commitment, order)
	}
	return cm.SendMessage(commit_receipt, order)
}

func (cm *ChainManager) handleAddCommitment(order *model.Order) error {
	srcHandler := cm.findSrcHandler(order.SrcChainId)
	if srcHandler == nil {
		errMsg := "handleAddCommitment findSrcHandler nil"
		util.Logger().Error(errMsg)
		return errors.New("handleAddCommitment findSrcHandler nil")
	}

	ret, err := srcHandler.AddCommitment(order)
	if err != nil {
		errMsg := fmt.Sprintf("handleAddCommitment addCommitment fail, err:%+v", err)
		util.Logger().Error(errMsg)
		return err
	}

	if !ret {
		errMsg := fmt.Sprintf("handleAddCommitment addCommitment trx failed, tx:%s", order.AddCommitmentTxHash)
		util.Logger().Error(errMsg)
		return err
	}
	err = model.UpdateOrderStatus(&model.Order{ID: order.ID, AddCommitmentTxStatus: 1, Status: core.AddCommitment, CommitmentIdx: order.CommitmentIdx}, cm.db)
	if err != nil {
		util.Logger().Error(fmt.Sprintf("handleAddCommitment update db err:%+v", err))
		return err
	}

	return cm.SendMessage(generate_zkproof, order)
}

func (cm *ChainManager) handleGenerateZkproof(order *model.Order) error {
	err := model.UpdateOrderStatus(&model.Order{ID: order.ID, Status: core.Generate}, cm.db)
	if err != nil {
		util.Logger().Error(fmt.Sprintf("handleGenerateZkproof update db err:%+v", err))
		return err
	}

	pm := zkp.GetProofManager(cm.cfg)
	if pm == nil {
		errMsg := "Init ProofManager err"
		util.Logger().Error(errMsg)
		return errors.New(errMsg)
	}

	sch := cm.findSrcHandler(order.SrcChainId)
	if sch == nil {
		errMsg := "handleGenerateZkproof:findSrcHandler nil"
		util.Logger().Error(errMsg)
		return errors.New(errMsg)
	}

	zkVerifier, err := sch.getZkVerifier()
	util.Logger().Info(fmt.Sprintf("zkVerifier:%s", zkVerifier.Hex()))
	if err != nil {
		util.Logger().Error(fmt.Sprintf("handleGenerateZkproof:getZkVerifier err:%+v", err))
		return err
	}

	req := &model.ProofReq{
		Addr: zkVerifier.Hex(),
		// Addr:   "0xED6BBe1286FAE2b21152915B0731303F8C6dEd06",
		Url:    sch.Rpc,
		TxHash: order.DestTxHash,
		CmtIdx: int(order.CommitmentIdx),
	}
	rp, raw := pm.GetZkRawProof(req)
	if rp == nil || raw == "" {
		errMsg := "get raw proof err"
		err = errors.New(errMsg)
		util.Logger().Error(errMsg)
		return err
	}

	order.RZKP = rp
	order.RawProof = raw

	zp := rp.RawProofToZkProof()
	if zp == nil {
		errMsg := "get zkp err"
		err = errors.New(errMsg)
		util.Logger().Error(errMsg)
		return err
	}
	order.ZKP = zp

	proofHash := zp.Keccak256EncodePackedZkProof()
	order.Proof = string(proofHash[:])
	model.UpdateOrderStatus(&model.Order{ID: order.ID, Proof: order.Proof, RawProof: raw}, cm.db)

	return cm.SendMessage(commit_receipt, order)
}

func (cm *ChainManager) handleCommitReceipt(order *model.Order) error {
	var err error
	var ret bool
	srcHandler := cm.findSrcHandler(order.SrcChainId)
	if srcHandler == nil {
		errMsg := "handleCommitReceipt findSrcHandler nil"
		err = errors.New(errMsg)
		util.Logger().Error(errMsg)
		return err
	}

	if cm.cfg.ZkVerify.Enable {
		ret, err = srcHandler.commitReceiptWithZk(order)
	} else {
		ret, err = srcHandler.commitReceipt(order)
	}

	if err != nil {
		util.Logger().Error(fmt.Sprintf("handleCommitReceipt commitReceipt err:%+v", err))
		return err
	}

	if !ret {
		util.Logger().Error(fmt.Sprintf("handleCommitReceipt commitReceipt failed, tx:%+v", order.AddCommitmentTxHash))
		return err
	}

	err = model.UpdateOrderStatus(&model.Order{ID: order.ID, ReceiptTxHash: order.ReceiptTxHash, ReceiptTxStatus: 1, Status: core.CommitReceipt}, cm.db)
	if err != nil {
		util.Logger().Error(fmt.Sprintf("handleCommitReceipt update db err:%+v", err))
		return err
	}

	if cm.cfg.ZkVerify.Enable {
		err = model.UpdateOrderProof(order.ID, order.Proof, cm.db)
		if err != nil {
			util.Logger().Error(fmt.Sprintf("handleCommitReceipt update db err:%+v", err))
			return err
		}
	}

	return cm.SendMessage(order_succeed, order)
}

func (cm *ChainManager) handleOrderSucceed(order *model.Order) error {
	return model.UpdateOrderStatus(&model.Order{ID: order.ID, Status: core.Success, FinishedTime: time.Now()}, cm.db)
}
