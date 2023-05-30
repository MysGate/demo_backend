package blockparser

import (
	"time"

	"github.com/MysGate/demo_backend/conf"
	"github.com/MysGate/demo_backend/util"
	"github.com/go-xorm/xorm"
)

type Manager struct {
	e       *xorm.Engine
	cfg     *conf.MysGateConfig
	parsers []*Parser
	ticker  *time.Ticker
	quit    chan bool
}

func InitParserManager(cfg *conf.MysGateConfig, e *xorm.Engine) *Manager {
	m := &Manager{
		cfg:  cfg,
		e:    e,
		quit: make(chan bool, 1),
	}

	for _, chain := range m.cfg.Chains {
		keys := m.cfg.GetChainKey(chain)
		m.runParsers(chain.HttpRpcUrl, keys)
	}

	go m.triggerParse()

	return m
}

func (m *Manager) triggerParse() {
	m.ticker = time.NewTicker(1500 * time.Millisecond)
	for {
		select {
		case <-m.ticker.C:
			m.parse()
		case <-m.quit:
			m.ticker.Stop()
			return
		}
	}
}

func (m *Manager) parse() {
	for _, p := range m.parsers {
		p.Parse()
	}
}

func (m *Manager) runParsers(rpc string, keys []string) {
	p := NewParser(rpc, keys, m.e)
	if p == nil {
		util.Logger().Error("RunParser:NewParser nil ")
		return
	}

	m.parsers = append(m.parsers, p)
	go p.Parse()
}

func (m *Manager) CloseParser() {
	m.quit <- true

	for _, p := range m.parsers {
		p.CloseParse()
	}
}
