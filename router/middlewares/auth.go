package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/lichunqiang/apiserver/handler"
	"github.com/lichunqiang/apiserver/pkg/token"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if _, err := token.ParseRequest(c); err != nil {
			handler.SendResponse(c, err, nil)
			c.Abort()
			return
		}

		c.Next()
	}
}
