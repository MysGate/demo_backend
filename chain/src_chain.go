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
)

type SrcChainHandler struct {
	SrcClient       *ethclient.Client
	QuitListen      chan bool
	ContractAddress common.Address
}

func newSrcChainHandler(client *ethclient.Client, addr common.Address) *SrcChainHandler {
	cch := &SrcChainHandler{
		SrcClient:       client,
		ContractAddress: addr,
		QuitListen:      make(chan bool, 1),
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
	logs := make(chan types.Log, 100)
	sub, err := sch.SrcClient.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}

FOR:
	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			if vLog.Topics[0].Hex() == orderTopic {
				abiJson := ""
				contractAbi, _ := abi.JSON(strings.NewReader(abiJson))
				orderEvent := util.OrderEvent{}
				err = contractAbi.UnpackIntoInterface(&orderEvent, "Order", vLog.Data)
				if err != nil {
					util.Logger().Error(fmt.Sprintf("[Order] failed to UnpackIntoInterface: %v", err))
				}
				order := module.Order{}
				order.SrcAddress = orderEvent.SrcAddress.Hex()
				order.SrcAmount = orderEvent.SrcAmount
				order.SrcToken = orderEvent.SrcToken.Hex()
				srcChainId, _ := sch.SrcClient.NetworkID(context.Background())
				order.SrcChainId = int(srcChainId.Int64())
				order.DestAddress = orderEvent.DestAddress.Hex()
				order.FinishedTime = time.Now()
				e := module.GetMySql()
				orderModel := module.Order{}
				orderModel.InsertOrder(order, e)
				fmt.Println(vLog)
			}
		case <-sch.QuitListen:
			break FOR
		}
	}

	fmt.Println("aaa")
}
