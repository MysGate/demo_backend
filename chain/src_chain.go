package chain

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/MysGate/demo_backend/module"
	"github.com/MysGate/demo_backend/util"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-xorm/xorm"
)

type SrcChainHandler struct {
	Client          *ethclient.Client
	QuitListen      chan bool
	ContractAddress common.Address
	Db              *xorm.Engine
	dispatcher      IDispatcher
}

func NewSrcChainHandler(client *ethclient.Client, addr common.Address, db *xorm.Engine, disp IDispatcher) *SrcChainHandler {
	cch := &SrcChainHandler{
		Client:          client,
		ContractAddress: addr,
		QuitListen:      make(chan bool, 10),
		Db:              db,
		dispatcher:      disp,
	}
	return cch
}

func (sch *SrcChainHandler) close() {
	sch.QuitListen <- true
}

var orderTopic string = util.GetOrderTopic()

func (sch *SrcChainHandler) runListenEvent() {
	query := ethereum.FilterQuery{
		Addresses: []common.Address{sch.ContractAddress},
	}

	logs := make(chan types.Log, 10000000)
	sub, err := sch.Client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}

FOR:
	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			sch.DispatchEvent(vLog)
		case <-sch.QuitListen:
			break FOR
		}
	}

	util.Logger().Info("SrcChainHandler: exit listen event")
}
func (sch *SrcChainHandler) DispatchEvent(vLog types.Log) {
	if vLog.Topics[0].Hex() != orderTopic {
		return
	}
	order, succeed := sch.parseEvent(vLog)
	if !succeed || order == nil {
		util.Logger().Error(fmt.Sprintf("DispatchEvent parseEvent failed: %+v", vLog))
		return
	}

	sch.dispatcher.DispatchCrossChainOrder(order)
}

func (sch *SrcChainHandler) parseEvent(vLog types.Log) (*module.Order, bool) {
	abiJson := ""
	contractAbi, _ := abi.JSON(strings.NewReader(abiJson))
	orderEvent := util.OrderEvent{}
	err := contractAbi.UnpackIntoInterface(&orderEvent, "Order", vLog.Data)
	if err != nil {
		util.Logger().Error(fmt.Sprintf("[Order] failed to UnpackIntoInterface: %+v", err))
		return nil, false
	}

	order := &module.Order{
		SrcAddress:   orderEvent.SrcAddress.Hex(),
		SrcAmount:    orderEvent.SrcAmount,
		SrcToken:     orderEvent.SrcToken.Hex(),
		DestAddress:  orderEvent.DestAddress.Hex(),
		FinishedTime: time.Now(),
	}

	srcChainId, _ := sch.Client.NetworkID(context.Background())
	order.SrcChainId = int(srcChainId.Int64())
	module.InsertOrder(order, sch.Db)
	fmt.Println(vLog)
	return order, true
}
