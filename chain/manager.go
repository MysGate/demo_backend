package chain

// event flow:  src_chain received event -> chain manager -> dest_chain

import (
	"fmt"

	"github.com/MysGate/demo_backend/conf"
	"github.com/MysGate/demo_backend/util"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-xorm/xorm"
)

type Connection struct {
	WssClient  *ethclient.Client
	HttpClient *ethclient.Client
}

type ChainHandler struct {
	src  *SrcChainHandler
	dest map[uint64]*DestChainHandler
}

type ChainManager struct {
	cfg      *conf.MysGateConfig
	db       *xorm.Engine
	clients  map[uint64]*Connection
	handlers map[uint64]*ChainHandler
	msgChan  chan *message
}

func InitChainManager(cfg *conf.MysGateConfig, db *xorm.Engine) *ChainManager {
	cm := &ChainManager{
		cfg:      cfg,
		db:       db,
		handlers: make(map[uint64]*ChainHandler),
		clients:  make(map[uint64]*Connection),
		msgChan:  make(chan *message, 1000000),
	}

	cm.startChainManager()
	return cm
}

func (cm *ChainManager) createEthClient(rpc string) *ethclient.Client {
	conn, err := ethclient.Dial(rpc)
	if err != nil {
		errMsg := fmt.Sprintf("createEthClient err is: %+v", err)
		util.Logger().Error(errMsg)
		return nil
	}
	return conn
}

func (cm *ChainManager) initEthClient(chanid uint64, wss, http string) *Connection {
	conn, ok := cm.clients[chanid]
	if !ok || conn == nil {
		co := &Connection{
			HttpClient: cm.createEthClient(http),
			WssClient:  cm.createEthClient(wss),
		}
		cm.clients[chanid] = co
		return co
	}

	if conn.HttpClient == nil {
		conn.HttpClient = cm.createEthClient(http)
	}

	if conn.WssClient == nil {
		conn.WssClient = cm.createEthClient(wss)
	}

	return conn
}

func (cm *ChainManager) start() {
	for src, dests := range cm.cfg.SupportCrossChain {
		cc := cm.cfg.FindCrossChain(src)
		if cc == nil {
			util.Logger().Error(fmt.Sprintf("chain manager find crosschain err:%+v ", cc))
			continue
		}

		conn := cm.initEthClient(cc.ChainID, cc.WssRpcUrl, cc.HttpRpcUrl)
		cch := NewSrcChainHandler(conn.WssClient, conn.HttpClient, cc.SrcAddr, cc.Key, cm.db, cm)
		if _, ok := cm.handlers[cc.ChainID]; !ok {
			ch := &ChainHandler{
				src:  cch,
				dest: make(map[uint64]*DestChainHandler),
			}
			cm.handlers[cc.ChainID] = ch
		}

		for _, dest := range dests {
			cd := cm.cfg.FindCrossChain(dest)
			if cd == nil {
				util.Logger().Error(fmt.Sprintf("chain manager find crosschain err:%+v ", cc))
				continue
			}

			conn := cm.initEthClient(cd.ChainID, cd.WssRpcUrl, cd.HttpRpcUrl)
			ccd := NewDestChainHandler(conn.HttpClient, cd.DestAddr, cd.Key)
			cm.handlers[cc.ChainID].dest[cd.ChainID] = ccd
		}

		go cch.runListenEvent()
	}
}

func (cm *ChainManager) startChainManager() {
	cm.start()
	go cm.messageLoop()
}

func (cm *ChainManager) CloseChainManager() {
	for _, sch := range cm.handlers {
		sch.src.close()
	}

	for _, conn := range cm.clients {
		conn.WssClient.Close()
		conn.HttpClient.Close()
	}

	cm.closeMessageLoop()
}
