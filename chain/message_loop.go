package chain

import (
	"errors"

	"github.com/MysGate/demo_backend/model"
	"github.com/MysGate/demo_backend/util"
)

type msgType int

const (
	pay_for_dest msgType = iota
	add_commitment
	generate_zkproof
	commit_receipt
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

				order, ok := msg.param.(*model.Order)
				if !ok {
					util.Logger().Info("pay_for_dest param err")
					continue
				}
				cm.handlePayForDest(order)
			case add_commitment:
				if msg.param == nil {
					util.Logger().Info("add_commitment param nil")
					continue
				}

				order, ok := msg.param.(*model.Order)
				if !ok {
					util.Logger().Info("add_commitment param err")
					continue
				}
				cm.handleAddCommitment(order)
			case generate_zkproof:
				util.Logger().Info("message loop msg: generate_zkproof")
				if msg.param == nil {
					util.Logger().Info("generate_zkproof param nil")
					continue
				}

				ch, ok := msg.param.(*model.Order)
				if !ok {
					util.Logger().Info("generate_zkproof param err")
					continue
				}

				cm.handleGenerateZkproof(ch)
			case commit_receipt:
				util.Logger().Info("commit_receipt loop msg: upload")
				if msg.param == nil {
					util.Logger().Info("commit_receipt param nil")
					continue
				}

				order, ok := msg.param.(*model.Order)
				if !ok {
					util.Logger().Info("commit_receipt param err")
					continue
				}
				cm.handleCommitReceipt(order)
			case order_succeed:
				util.Logger().Info("message loop msg: order_succeed")
				if msg.param == nil {
					util.Logger().Info("verify_zkproof param nil")
					continue
				}

				order, ok := msg.param.(*model.Order)
				if !ok {
					util.Logger().Info("verify_zkproof param err")
					continue
				}
				cm.handleOrderSucceed(order)
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

func (cm *ChainManager) SendMessage(msg msgType, order *model.Order) error {
	if cm.isLoopExit() {
		errMsg := "SendMessage: channel is close"
		err := errors.New(errMsg)
		util.Logger().Error(errMsg)
		return err
	}

	cm.msgChan <- &message{op: msg, param: order}
	return nil
}

func (cm *ChainManager) PayForDest(order *model.Order) error {
	return cm.SendMessage(pay_for_dest, order)
}
