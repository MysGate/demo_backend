package chain

// event flow:  src_chain received event -> chain manager -> dest_chain

import (
	"fmt"

	"github.com/MysGate/demo_backend/conf"
	"github.com/MysGate/demo_backend/util"
)

type ChainManager struct {
	cfg          *conf.MysGateConfig
	srcHandlers  map[uint64]*SrcChainHandler
	destHandlers map[uint64][]*DestChainHandler
}

func newChainManager(cfg *conf.MysGateConfig) *ChainManager {
	cm := &ChainManager{
		cfg:          cfg,
		srcHandlers:  make(map[uint64]*SrcChainHandler),
		destHandlers: make(map[uint64][]*DestChainHandler),
	}

	return cm
}

func (cm *ChainManager) start() {
	for src, dests := range cm.cfg.SupportCrossChain {
		cc := cm.cfg.FindCrossChain(src)
		if cc == nil {
			util.Logger().Error(fmt.Sprintf("chain manager find crosschain err:%+v ", cc))
			continue
		}

		cch := newSrcChainHandler(cc.SrcClient, cc.SrcAddr)
		cm.srcHandlers[cc.ChainID] = cch
		go cch.runListenEvent()

		for _, dest := range dests {
			cd := cm.cfg.FindCrossChain(dest)
			if cd == nil {
				util.Logger().Error(fmt.Sprintf("chain manager find crosschain err:%+v ", cc))
				continue
			}

			ccd := &DestChainHandler{
				DestClient:      cd.DestClient,
				ContractAddress: cd.DestAddr,
			}
			cm.destHandlers[cc.ChainID] = append(cm.destHandlers[cc.ChainID], ccd)
		}
	}
}

func StartChainManager(cfg *conf.MysGateConfig) *ChainManager {
	cm := newChainManager(cfg)
	go cm.start()
	return cm
}

func (cm *ChainManager) CloseChainManager(cfg *conf.MysGateConfig) *ChainManager {
	for _, sch := range cm.srcHandlers {
		sch.close()
	}
	cm.cfg.CloseClient()
	return cm
}
