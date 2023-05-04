package service

import (
	"net/http"

	"github.com/MysGate/demo_backend/module"
	"github.com/gin-gonic/gin"
)

func (s *Server) ping(c *gin.Context) {
	m := module.GetMessage(module.PONG)
	c.JSON(http.StatusOK, m)
}

func (s *Server) getCrossChainPair(c *gin.Context) {
	c.String(200, "pair")
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
