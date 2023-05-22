package rpc

import (
	"github.com/MysGate/demo_backend/conf"
	"github.com/gin-gonic/gin"
	"github.com/go-xorm/xorm"
)

type Server struct {
	cfg *conf.MysGateConfig
	e   *gin.Engine
	db  *xorm.Engine
}

func NewHttpServer(c *conf.MysGateConfig, db *xorm.Engine) (s *Server) {
	s = &Server{
		cfg: c,
		e:   gin.Default(),
		db:  db,
	}
	s.initRouter()
	return
}

func (s *Server) initRouter() {
	s.e.GET("/ping", s.ping)
	s.e.GET("/coin", s.getSupportCoins)
	s.e.GET("/pair", s.getCrossChainPair)
	s.e.GET("/fee", s.getFee)
	s.e.GET("/order/nextid", s.getNextOrderID)
	s.e.POST("/porter", s.getPorters)
	s.e.POST("/cost", s.getCost)
	s.e.POST("/order/search", s.orderSearch)
	s.e.POST("/order/list", s.orderList)
}

func (s *Server) RunHttpService() {
	s.e.Run(s.cfg.Service.ServicePort)
}
