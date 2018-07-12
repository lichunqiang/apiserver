package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/lichunqiang/apiserver/pkg/errno"
	"net/http"
)

type ApiResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func SendResponse(c *gin.Context, err error, data interface{}) {
	code, message := errno.DecodeErr(err)

	c.JSON(http.StatusOK, ApiResponse{
		Code:    code,
		Message: message,
		Data:    data,
	})
}
