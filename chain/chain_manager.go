package chain

// event flow:  src_chain received event -> chain manager -> dest_chain

import (
	"fmt"

	"github.com/MysGate/demo_backend/conf"
	"github.com/MysGate/demo_backend/module"
	"github.com/MysGate/demo_backend/util"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-xorm/xorm"
)

type Conn struct {
	SrcClient  *ethclient.Client
	DestClient *ethclient.Client
}

type ChainManager struct {
	cfg          *conf.MysGateConfig
	srcHandlers  map[uint64]*SrcChainHandler
	destHandlers map[uint64][]*DestChainHandler
	clients      map[uint64]*Conn
	db           *xorm.Engine
}

func newChainManager(cfg *conf.MysGateConfig, db *xorm.Engine) *ChainManager {
	cm := &ChainManager{
		cfg:          cfg,
		srcHandlers:  make(map[uint64]*SrcChainHandler),
		destHandlers: make(map[uint64][]*DestChainHandler),
		clients:      make(map[uint64]*Conn),
		db:           db,
	}

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

func (cm *ChainManager) start() {
	for src, dests := range cm.cfg.SupportCrossChain {
		cc := cm.cfg.FindCrossChain(src)
		if cc == nil {
			util.Logger().Error(fmt.Sprintf("chain manager find crosschain err:%+v ", cc))
			continue
		}

		srcClient := cm.createEthClient(cc.SrcRpcUrl)

		cch := NewSrcChainHandler(srcClient, cc.SrcAddr, cm.db, cm)
		cm.srcHandlers[cc.ChainID] = cch
		go cch.runListenEvent()

		for _, dest := range dests {
			cd := cm.cfg.FindCrossChain(dest)
			if cd == nil {
				util.Logger().Error(fmt.Sprintf("chain manager find crosschain err:%+v ", cc))
				continue
			}
			destClient := cm.createEthClient(cd.DestRpcUrl)
			ccd := NewDestChainHandler(destClient, cd.DestAddr, cd.Key)
			cm.destHandlers[cc.ChainID] = append(cm.destHandlers[cc.ChainID], ccd)
		}
	}
}

func StartChainManager(cfg *conf.MysGateConfig, db *xorm.Engine) *ChainManager {
	cm := newChainManager(cfg, db)
	go cm.start()
	return cm
}

func (cm *ChainManager) CloseChainManager() *ChainManager {
	for _, sch := range cm.srcHandlers {
		sch.close()
	}

	for _, conn := range cm.clients {
		conn.SrcClient.Close()
		conn.DestClient.Close()
	}

	return cm
}

func (cm *ChainManager) DispatchCrossChainOrder(*module.Order) error {
	return nil
}
