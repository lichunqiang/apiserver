package util

import (
	"github.com/gin-gonic/gin"
	"github.com/lichunqiang/apiserver/pkg/constvar"
	"github.com/teris-io/shortid"
)

func GenShortId() (string, error) {
	return shortid.Generate()
}

func GetReqId(c *gin.Context) string {
	v, ok := c.Get(constvar.XRequestId)
	if !ok {
		return ""
	}

	if requestId, ok := v.(string); ok {
		return requestId
	}

	return ""
}
