package service

import (
	"fmt"

	"github.com/MysGate/demo_backend/core"
	"github.com/MysGate/demo_backend/core/errno"
	"github.com/MysGate/demo_backend/module"
	"github.com/MysGate/demo_backend/util"
	"github.com/gin-gonic/gin"
)

func (s *Server) orderSearch(c *gin.Context) {
	var requestData core.OrderSearchRequest
	if err := c.BindJSON(&requestData); err != nil {
		errMsg := fmt.Sprintf("Failed to bind request params, reason=[%s]", err)
		util.Logger().Error(errMsg)
		core.SendResponse(c, errno.BindErr, nil)
		return
	}

	has, result := module.GetOrder(requestData.OrderId, s.db)
	if has {
		core.SendResponse(c, errno.OK, result)
	} else {
		core.SendResponse(c, errno.InternalServerErr, nil)
	}
}

func (s *Server) orderList(c *gin.Context) {
	var requestData core.OrderListRequest
	if err := c.BindJSON(&requestData); err != nil {
		errMsg := fmt.Sprintf("Failed to bind request params, reason=[%s]", err)
		util.Logger().Error(errMsg)
		core.SendResponse(c, errno.BindErr, nil)
		return
	}

	err, orders := module.GetOrderList(requestData.SrcChainId, requestData.DestChainId, s.db)
	if err != nil {
		core.SendResponse(c, errno.InternalServerErr, nil)
	} else {
		core.SendResponse(c, errno.OK, orders)
	}
}
