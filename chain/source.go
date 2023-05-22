package chain

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/MysGate/demo_backend/contracts"
	"github.com/MysGate/demo_backend/model"
	"github.com/MysGate/demo_backend/util"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-xorm/xorm"
)

type SrcChainHandler struct {
	Db              *xorm.Engine
	WssClient       *ethclient.Client
	HttpClient      *ethclient.Client
	QuitListen      chan bool
	ContractAddress common.Address
	disp            IDispatcher
}

func NewSrcChainHandler(wssClient *ethclient.Client, httpClient *ethclient.Client, addr common.Address, db *xorm.Engine, disp IDispatcher) *SrcChainHandler {
	cch := &SrcChainHandler{
		WssClient:       wssClient,
		HttpClient:      httpClient,
		ContractAddress: addr,
		QuitListen:      make(chan bool, 10),
		Db:              db,
		disp:            disp,
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
	sub, err := sch.WssClient.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			sch.DispatchEvent(vLog)
		case <-sch.QuitListen:
			return
		}
	}
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

	sch.disp.PayForDest(order)
}

func (sch *SrcChainHandler) parseEvent(vLog types.Log) (*model.Order, bool) {
	contractAbi, err := abi.JSON(strings.NewReader(string(contracts.CrossABI)))
	if err != nil {
		util.Logger().Error(fmt.Sprintf("Not found abi json, err:%+v", err))
		return nil, false
	}

	orderEvent := &contracts.CrossControllerOrder{}
	err = contractAbi.UnpackIntoInterface(orderEvent, "Order", vLog.Data)
	if err != nil {
		util.Logger().Error(fmt.Sprintf("[Order] failed to UnpackIntoInterface: %+v", err))
		return nil, false
	}

	order := &model.Order{
		ID:          orderEvent.OrderId.Int64(),
		SrcAddress:  orderEvent.SrcAddress.Hex(),
		SrcAmount:   util.ConvertTokenAmountToFloat64(orderEvent.SrcAmount.String(), 18),
		SrcToken:    orderEvent.SrcToken.Hex(),
		SrcChainId:  orderEvent.SrcChainId.Uint64(),
		DestAddress: orderEvent.DestAddress.Hex(),
		DestChainId: orderEvent.DestChainId.Uint64(),
		DestToken:   orderEvent.DestAddress.Hex(),
		PoterId:     orderEvent.PorterPool.Hex(),
		Created:     time.Now(),
	}

	srcChainId, _ := sch.HttpClient.NetworkID(context.Background())
	order.SrcChainId = srcChainId.Uint64()
	model.InsertOrder(order, sch.Db)
	return order, true
}

func (sch *SrcChainHandler) commitReceipt(order *model.Order) error {
	// TODO call crossController.sol::commitReceipt
	return nil
}
