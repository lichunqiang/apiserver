package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/lichunqiang/apiserver/pkg/constvar"
	"github.com/satori/go.uuid"
)

func RequestId() gin.HandlerFunc {
	return func(c *gin.Context) {
		//check request header
		reqId := c.Request.Header.Get(constvar.XRequestId)

		if reqId == "" {
			u4 := uuid.NewV4()
			reqId = u4.String()
		}

		//used for application
		c.Set(constvar.XRequestId, reqId)

		c.Writer.Header().Set(constvar.XRequestId, reqId)
		c.Next()
	}
}
