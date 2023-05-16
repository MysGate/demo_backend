package core

import (
	"net/http"

	"github.com/MysGate/demo_backend/core/errno"
	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SendResponse(c *gin.Context, err *errno.Errno, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code:    err.Code,
		Message: err.Message,
		Data:    data,
	})
}
