package rpc

import (
	"fmt"
	"strconv"

	"github.com/MysGate/demo_backend/core"
	"github.com/MysGate/demo_backend/core/errno"
	"github.com/MysGate/demo_backend/model"
	"github.com/MysGate/demo_backend/util"
	"github.com/gin-gonic/gin"
)

func (s *Server) orderSearch(c *gin.Context) {
	orderId := c.Query("orderid")
	id, err := strconv.ParseInt(orderId, 10, 64)
	if err != nil {
		errMsg := fmt.Sprintf("Failed to parse params, reason=[%s]", err)
		util.Logger().Error(errMsg)
		core.SendResponse(c, errno.InvalidRequestParameter, nil)
		return
	}
	has, result := model.GetOrder(id, s.db)
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

	err, orders := model.GetOrderList(requestData.SrcChainId, requestData.DestChainId, s.db)
	if err != nil {
		core.SendResponse(c, errno.InternalServerErr, nil)
	} else {
		core.SendResponse(c, errno.OK, orders)
	}
}
