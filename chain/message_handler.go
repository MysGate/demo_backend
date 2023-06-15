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
	err := destHandler.crossFrom(order)
	if err != nil {
		errMsg := fmt.Sprintf("handlePayForDest crossFrom fail, err:%+v", err)
		util.Logger().Error(errMsg)
		return err
	}

	err = model.UpdateOrderStatus(&model.Order{ID: order.ID, DestTxHash: order.DestTxHash, DestTxStatus: 1}, cm.db)
	if err != nil {
		util.Logger().Error(fmt.Sprintf("handlePayForDest update db err:%+v", err))
		return err
	}
	cm.AddCommitment(order)
	return nil
}

func (cm *ChainManager) handleAddCommitment(order *model.Order) error {
	err := model.UpdateOrderStatus(&model.Order{ID: order.ID, Status: core.AddCommitment}, cm.db)
	if err != nil {
		util.Logger().Error(fmt.Sprintf("handleAddCommitment update db err:%+v", err))
		return err
	}

	srcHandler := cm.findSrcHandler(order.SrcChainId)
	if srcHandler == nil {
		errMsg := "handleAddCommitment findSrcHandler nil"
		err = errors.New(errMsg)
		util.Logger().Error(errMsg)
		return err
	}

	// TODO Call contract to addCommitment
	err = srcHandler.AddCommitment(order)
	if err != nil {
		errMsg := fmt.Sprintf("handleAddCommitment addCommitment fail, err:%+v", err)
		util.Logger().Error(errMsg)
		return err
	}

	cm.GenerateZkProof(order)
	return nil
}

func (cm *ChainManager) handleGenerateZkproof(order *model.Order) error {
	err := model.UpdateOrderStatus(&model.Order{ID: order.ID, Status: core.Generate}, cm.db)
	if err != nil {
		util.Logger().Error(fmt.Sprintf("handleGenerateZkproof update db err:%+v", err))
		return err
	}

	if cm.cfg.VerifyWithZk {
		pm := zkp.GetProofManager(cm.cfg)
		if pm == nil {
			errMsg := "Init ProofManager err"
			err = errors.New(errMsg)
			util.Logger().Error(errMsg)
			return err
		}

		// TODO call zkp(need modify: pass the params)
		p := pm.GetZKProof()
		if p == nil {
			errMsg := "get zkp err"
			err = errors.New(errMsg)
			util.Logger().Error(errMsg)
			return err
		}
		// TODO  Modify it
		// order.RawProof =
		// order.Proof = p.Proof
		// model.UpdateOrderRawProof(order.ID, p.rawProof, cm.db)
	}

	err = cm.CommitReceipt(order)
	if err != nil {
		util.Logger().Error(fmt.Sprintf("handleGenerateZkproof err:%+v", err))
		return err
	}

	return nil
}

func (cm *ChainManager) handleCommitReceipt(order *model.Order) error {
	err := model.UpdateOrderStatus(&model.Order{ID: order.ID, Status: core.Verify}, cm.db)
	if err != nil {
		util.Logger().Error(fmt.Sprintf("handleVerifyZkproof update db err:%+v", err))
		return err
	}

	srcHandler := cm.findSrcHandler(order.SrcChainId)
	if srcHandler == nil {
		errMsg := "handleVerifyZkproof findSrcHandler nil"
		err = errors.New(errMsg)
		util.Logger().Error(errMsg)
		return err
	}

	if cm.cfg.VerifyWithZk {
		err = srcHandler.commitReceiptWithZk(order)
	} else {
		err = srcHandler.commitReceipt(order)
	}

	if err != nil {
		util.Logger().Error(fmt.Sprintf("handleVerifyZkproof commitReceipt err:%+v", err))
		return err
	}

	err = model.UpdateOrderStatus(&model.Order{ID: order.ID, ReceiptTxHash: order.ReceiptTxHash, Status: core.CommitReceipt}, cm.db)
	if err != nil {
		util.Logger().Error(fmt.Sprintf("handleVerifyZkproof update db err:%+v", err))
		return err
	}

	if cm.cfg.VerifyWithZk {
		err = model.UpdateOrderProof(order.ID, order.Proof, cm.db)
		if err != nil {
			util.Logger().Error(fmt.Sprintf("handleVerifyZkproof update db err:%+v", err))
			return err
		}
	}

	err = cm.OrderSucceed(order)
	if err != nil {
		util.Logger().Error(fmt.Sprintf("handleVerifyZkproof err:%+v", err))
		return err
	}
	return nil
}

func (cm *ChainManager) handleOrderSucceed(order *model.Order) error {
	return model.UpdateOrderStatus(&model.Order{ID: order.ID, Status: core.Success, FinishedTime: time.Now()}, cm.db)
}
