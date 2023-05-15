package service

import (
	"fmt"
	"net/http"

	"github.com/MysGate/demo_backend/module"
	"github.com/MysGate/demo_backend/router"
	"github.com/MysGate/demo_backend/util"
	"github.com/gin-gonic/gin"
)

func (s *Server) ping(c *gin.Context) {
	m := module.GetMessage(module.PONG)
	c.JSON(http.StatusOK, m)
}

func (s *Server) getSupportCoins(c *gin.Context) {
	c.JSON(http.StatusOK, s.cfg.SupportCoins)
}

func (s *Server) getCrossChainPair(c *gin.Context) {
	mccp := s.getSupportChainPair()
	if mccp == nil {
		m := module.GetMessage(module.INTERNAL_ERROR)
		c.JSON(http.StatusInternalServerError, m)
		return
	}

	c.JSON(http.StatusOK, mccp)
}

func (s *Server) getFee(c *gin.Context) {
	coin := c.Query("coin")
	ccf := s.cfg.GetCrossChainFee(coin)
	if ccf == nil {
		m := module.GetMessage(module.PARAM_ERROR)
		c.JSON(http.StatusInternalServerError, m)
		return
	}

	f := &module.Fee{
		Fixed:     ccf.Fixed,
		FloatRate: ccf.FloatRate,
	}
	c.JSON(http.StatusOK, f)
}

func (s *Server) getPorters(c *gin.Context) {
	req := module.RoterReq{}
	err := c.ShouldBind(&req)
	if err != nil {
		m := module.GetMessage(module.PARAM_ERROR)
		c.JSON(http.StatusInternalServerError, m)
		return
	}

	rm := router.GetManager(s.cfg)
	var routers []*module.Router
	porters := rm.SelectPorters(&req)
	for _, p := range porters {
		transfered, completion, err := module.GetPorterTransferedAndCompletion(p.Address, s.db)
		if err != nil {
			util.Logger().Error(fmt.Sprintf("getPorters err: +v", err))
			continue
		}
		p.Transfered = transfered
		p.Completion = completion

		r1 := &module.Router{
			From: req.From,
			To:   p.Address,
		}
		routers = append(routers, r1)

		r2 := &module.Router{
			From: p.Address,
			To:   req.To,
		}
		routers = append(routers, r2)
	}

	resp := &module.RoterResponse{
		Porters: porters,
		Routers: routers,
	}

	c.JSON(http.StatusOK, resp)
}

func (s *Server) getCost(c *gin.Context) {
	req := module.CostReq{}
	err := c.ShouldBind(&req)
	if err != nil {
		m := module.GetMessage(module.PARAM_ERROR)
		c.JSON(http.StatusInternalServerError, m)
		return
	}

	cal := s.cfg.GetCoinLimit(req.Coin)
	if cal == nil {
		m := module.GetMessage(module.PARAM_ERROR)
		c.JSON(http.StatusInternalServerError, m)
		return
	}

	if req.Amount < cal.MinAmount || req.Amount > cal.MaxAmount {
		m := module.GetMessage(module.AMOUNT_ERROR)
		c.JSON(http.StatusInternalServerError, m)
		return
	}

	ccf := s.cfg.GetCrossChainFee(req.Coin)
	if ccf == nil {
		m := module.GetMessage(module.PARAM_ERROR)
		c.JSON(http.StatusInternalServerError, m)
		return
	}

	resp := &module.CostResp{
		Coin:      req.Coin,
		Amount:    req.Amount,
		Fixed:     ccf.Fixed,
		FloatRate: ccf.FloatRate,
		RealSend:  req.Amount + ccf.Fixed,
		Received:  req.Amount - req.Amount*ccf.FloatRate,
		FloatFee:  req.Amount * ccf.FloatRate,
	}

	c.JSON(http.StatusOK, resp)
}
