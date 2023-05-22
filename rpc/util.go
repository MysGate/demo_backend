package rpc

import (
	"github.com/MysGate/demo_backend/conf"
	"github.com/MysGate/demo_backend/model"
	"github.com/MysGate/demo_backend/util"
)

func (s *Server) convertConfChain2ModelChain(src *conf.Chain) *model.Chain {
	mc := &model.Chain{
		ChainName:    src.Name,
		ChainID:      src.ChainID,
		SupportCoins: src.SuppirtCoins,
	}
	return mc
}

func (s *Server) getSupportChainPair() *model.CrossChainPair {
	ccp := &model.CrossChainPair{}
	for src, dests := range s.cfg.SupportCrossChain {
		cc := s.cfg.FindCrossChain(src)
		if cc == nil {
			util.Logger().Error("getCrossChainPair err")
			continue
		}
		mc := s.convertConfChain2ModelChain(cc)
		ccp.SrcChain = mc
		for _, dest := range dests {
			d := s.cfg.FindCrossChain(dest)
			if d == nil {
				util.Logger().Error("getCrossChainPair err")
				continue
			}
			mdc := s.convertConfChain2ModelChain(d)
			ccp.DestChains = append(ccp.DestChains, mdc)
		}
	}
	return ccp
}
