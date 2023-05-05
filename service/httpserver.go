package service

import (
	"github.com/MysGate/demo_backend/conf"
	"github.com/gin-gonic/gin"
)

type Server struct {
	cfg *conf.MysGateConfig
	e   *gin.Engine
}

func NewHttpServer(c *conf.MysGateConfig) (s *Server) {
	s = &Server{
		cfg: c,
		e:   gin.Default(),
	}
	s.initRouter()
	return
}

func (s *Server) initRouter() {
	s.e.GET("/ping", s.ping)
	s.e.GET("/coin", s.getSupportCoins)
	s.e.GET("/pair", s.getCrossChainPair)
	s.e.GET("/fee", s.getFee)
	s.e.POST("/cost", s.getCost)
}

func (s *Server) RunHttpService() {
	s.e.Run(s.cfg.Service.ServicePort)
}
