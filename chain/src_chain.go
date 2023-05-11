package chain

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum"
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
			fmt.Println(vLog)
		case <-sch.QuitListen:
			break FOR
		}
	}

	fmt.Println("aaa")
}
