package chain

import (
	"errors"
	"fmt"

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

func (cm *ChainManager) handlerPayForDest(order *model.Order) error {
	destHandler := cm.findDestHandler(order.SrcChainId, order.DestChainId)
	err := destHandler.crossFrom(order)
	if err != nil {
		errMsg := fmt.Sprintf("handlerPayForDest crossFrom fail, err:%+v", err)
		util.Logger().Error(errMsg)
		return err
	}

	err = model.UpdateOrderDestTxHash(order, cm.db)
	if err != nil {
		util.Logger().Error(fmt.Sprintf("handlerPayForDest update db err:%+v", err))
		return err
	}

	return nil
}

func (cm *ChainManager) handlerGenerateZkproof(order *model.Order) error {
	err := model.UpdateOrderStatus(order.ID, core.Generate, cm.db)
	if err != nil {
		util.Logger().Error(fmt.Sprintf("handlerGenerateZkproof update db err:%+v", err))
		return err
	}

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
	order.Proof = p.Proof
	err = cm.VerifyZkProof(order)
	if err != nil {
		util.Logger().Error(fmt.Sprintf("handlerGenerateZkproof err:%+v", err))
		return err
	}

	return nil
}

func (cm *ChainManager) handlerVerifyZkproof(order *model.Order) error {
	err := model.UpdateOrderStatus(order.ID, core.Verify, cm.db)
	if err != nil {
		util.Logger().Error(fmt.Sprintf("handlerVerifyZkproof update db err:%+v", err))
		return err
	}

	srcHandler := cm.findSrcHandler(order.SrcChainId)
	if srcHandler == nil {
		errMsg := "handlerVerifyZkproof findSrcHandler nil"
		err = errors.New(errMsg)
		util.Logger().Error(errMsg)
		return err
	}

	err = srcHandler.commitReceipt(order)
	if err != nil {
		util.Logger().Error(fmt.Sprintf("handlerVerifyZkproof commitReceipt err:%+v", err))
		return err
	}

	err = model.UpdateOrderProof(order.ID, order.Proof, cm.db)
	if err != nil {
		util.Logger().Error(fmt.Sprintf("handlerVerifyZkproof update db err:%+v", err))
		return err
	}

	err = cm.OrderSucceed(order)
	if err != nil {
		util.Logger().Error(fmt.Sprintf("handlerVerifyZkproof err:%+v", err))
		return err
	}
	return nil
}

func (cm *ChainManager) handlerOrderSucceed(order *model.Order) error {
	return model.UpdateOrderStatus(order.ID, core.Success, cm.db)
}
