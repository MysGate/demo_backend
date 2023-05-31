package blockparser

import (
	"context"
	"fmt"
	"math/big"
	"strconv"
	"strings"
	"time"

	"github.com/MysGate/demo_backend/model"
	"github.com/MysGate/demo_backend/pubsub"
	"github.com/MysGate/demo_backend/util"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-xorm/xorm"
)

type Parser struct {
	keys   map[string]string
	e      *xorm.Engine
	client *ethclient.Client
	quit   chan bool
	work   chan bool
}

func NewParser(rpc string, keys []string, e *xorm.Engine) *Parser {
	conn, err := ethclient.Dial(rpc)
	if err != nil {
		errMsg := fmt.Sprintf("createEthClient err is: %+v", err)
		util.Logger().Error(errMsg)
		return nil
	}

	p := &Parser{
		e:      e,
		client: conn,
		keys:   make(map[string]string),
		quit:   make(chan bool, 1),
		work:   make(chan bool, 1),
	}

	for _, v := range keys {
		p.keys[v] = v
	}

	return p
}

func (p *Parser) parse() {
	for {
		select {
		case <-p.work:
			p.parseImpl()
		case <-p.quit:
			return
		}
	}
}

func (p *Parser) parseImpl() {
	util.Logger().Info("parseImpl")
	header, _ := p.client.HeaderByNumber(context.Background(), nil)
	for _, topic := range p.keys {
		// 1: parse block
		key := strings.Split(topic, ":")
		chanId, _ := strconv.ParseInt(key[0], 10, 64)
		has, block := model.GetBlock(chanId, key[1], p.e)
		if !has {
			block = &model.Block{
				ChainId:     int(chanId),
				Contract:    key[1],
				BlockNumber: header.Number.Int64() - 10000}
			model.InsertBlock(block, p.e)
		}
		query := ethereum.FilterQuery{
			FromBlock: big.NewInt(block.BlockNumber + 1),
			ToBlock:   header.Number,
			Addresses: []common.Address{
				common.HexToAddress(key[1]),
			},
		}
		logs, err := p.client.FilterLogs(context.Background(), query)
		if err != nil {
			util.Logger().Error(fmt.Sprintf("chainId %d failed to GetEthLogs: %v", chanId, err))
			return
		}
		// 2: broastcastLog
		for _, log := range logs {
			p.broastcastLog(topic, &log)
		}

		// 3: Upsert db parsed blocknumber
		block.BlockNumber = header.Number.Int64()
		block.Updated = time.Now()
		model.UpdateBlock(block, p.e)
	}

}

func (p *Parser) broastcastLog(key string, vLog *types.Log) {
	m := pubsub.GetSubscribeManager()
	m.TryPublish(key, vLog)
}

func (p *Parser) closeParse() {
	p.quit <- true
}

func (p *Parser) doWork() {
	p.work <- true
}