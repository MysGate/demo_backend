package chain

import (
	"errors"

	"github.com/MysGate/demo_backend/module"
	"github.com/MysGate/demo_backend/util"
)

type msgType int

const (
	pay_for_dest msgType = iota
	generate_zkproof
	verify_zkproof
	order_succeed
	close
)

type message struct {
	op    msgType
	param interface{}
}

func (cm *ChainManager) messageLoop() {
	for {
		select {
		case msg, open := <-cm.msgChan:
			if !open {
				return
			}
			switch msg.op {
			case pay_for_dest:
				if msg.param == nil {
					util.Logger().Info("generate_zkproof param nil")
					continue
				}

				order, ok := msg.param.(*module.Order)
				if !ok {
					util.Logger().Info("pay_for_dest param err")
					continue
				}
				cm.handlerPayForDest(order)
			case generate_zkproof:
				util.Logger().Info("message loop msg: generate_zkproof")
				if msg.param == nil {
					util.Logger().Info("generate_zkproof param nil")
					continue
				}

				ch, ok := msg.param.(*module.Order)
				if !ok {
					util.Logger().Info("generate_zkproof param err")
					continue
				}

				cm.handlerGenerateZkproof(ch)

			case verify_zkproof:
				util.Logger().Info("verify_zkproof loop msg: upload")
				if msg.param == nil {
					util.Logger().Info("verify_zkproof param nil")
					continue
				}

				order, ok := msg.param.(*module.Order)
				if !ok {
					util.Logger().Info("verify_zkproof param err")
					continue
				}
				cm.handlerVerifyZkproof(order)
			case order_succeed:
				util.Logger().Info("message loop msg: order_succeed")
				if msg.param == nil {
					util.Logger().Info("verify_zkproof param nil")
					continue
				}

				order, ok := msg.param.(*module.Order)
				if !ok {
					util.Logger().Info("verify_zkproof param err")
					continue
				}
				cm.handlerOrderSucceed(order)
			case close:
				util.Logger().Info("message loop msg: close")
				return
			}

		}
	}
}

func (m *ChainManager) isLoopExit() bool {
	return m.msgChan == nil
}

func (cm *ChainManager) closeMessageLoop() {
	if cm.isLoopExit() {
		util.Logger().Error("channel is close, close msg not implement")
		return
	}

	cm.msgChan <- &message{op: close}
	return
}

func (cm *ChainManager) PayForDest(order *module.Order) error {
	if cm.isLoopExit() {
		errMsg := "PayForDest: channel is close"
		err := errors.New(errMsg)
		util.Logger().Error(errMsg)
		return err
	}

	cm.msgChan <- &message{op: pay_for_dest, param: order}
	return nil
}

func (cm *ChainManager) GenerateZkProof(order *module.Order) error {
	if cm.isLoopExit() {
		errMsg := "GenerateZkProof: channel is close"
		err := errors.New(errMsg)
		util.Logger().Error(errMsg)
		return err
	}

	cm.msgChan <- &message{op: generate_zkproof, param: order}
	return nil
}

func (cm *ChainManager) VerifyZkProof(order *module.Order) error {
	if cm.isLoopExit() {
		errMsg := "VerifyZkProof: channel is close"
		err := errors.New(errMsg)
		util.Logger().Error(errMsg)
		return err
	}

	cm.msgChan <- &message{op: verify_zkproof, param: order}
	return nil
}

func (cm *ChainManager) OrderSucceed(order *module.Order) error {
	if cm.isLoopExit() {
		errMsg := "OrderSucceed: channel is close"
		err := errors.New(errMsg)
		util.Logger().Error(errMsg)
		return err
	}

	cm.msgChan <- &message{op: order_succeed, param: order}
	return nil
}
