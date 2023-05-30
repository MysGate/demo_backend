package blockparser

import (
	"fmt"

	"github.com/MysGate/demo_backend/pubsub"
	"github.com/MysGate/demo_backend/util"
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
	// 1: parse block
	// 2: broastcastLog
	// 3: Upsert db parsed blocknumber
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
