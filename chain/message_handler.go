package chain

import (
	"fmt"

	"github.com/MysGate/demo_backend/module"
	"github.com/MysGate/demo_backend/util"
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

func (cm *ChainManager) handlerPayForDest(order *module.Order) error {
	destHandler := cm.findDestHandler(order.SrcChainId, order.DestChainId)
	err := destHandler.crossFrom(order)
	if err != nil {
		errMsg := fmt.Sprintf("handlerPayForDest crossFrom fail, err:%+v", err)
		util.Logger().Error(errMsg)
		return err
	}
	return nil
}

func (cm *ChainManager) handlerGenerateZkproof(order *module.Order) error {
	return nil
}

func (cm *ChainManager) handlerVerifyZkproof(order *module.Order) error {
	return nil
}

func (cm *ChainManager) handlerOrderSucceed(order *module.Order) error {
	return nil
}
