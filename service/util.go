package service

import (
	"github.com/MysGate/demo_backend/conf"
	"github.com/MysGate/demo_backend/module"
	"github.com/MysGate/demo_backend/util"
)

func (s *Server) convertConfChain2ModuleChain(src *conf.Chain) *module.Chain {
	mc := &module.Chain{
		ChainName: src.Name,
		ChainID:   src.ChainID,
	}
	return mc
}

func (s *Server) getSupportChainPair() *module.CrossChainPair {
	ccp := &module.CrossChainPair{}
	for src, dests := range s.cfg.SupportCrossChain {
		cc := s.cfg.FindCrossChain(src)
		if cc == nil {
			util.Logger().Error("getCrossChainPair err")
			continue
		}
		mc := s.convertConfChain2ModuleChain(cc)
		ccp.SrcChain = mc
		for _, dest := range dests {
			d := s.cfg.FindCrossChain(dest)
			if d == nil {
				util.Logger().Error("getCrossChainPair err")
				continue
			}
			mdc := s.convertConfChain2ModuleChain(d)
			ccp.DestChains = append(ccp.DestChains, mdc)
		}
	}
	return ccp
}
