package middlewares

import (
	"github.com/gin-gonic/gin"
	"time"
	"net/http"
)

func NoCache(c *gin.Context) {
	c.Writer.Header().Add("Cache-Control", "no-cache, no-store, max-age=0, must-revalidate, value")
	c.Writer.Header().Add("Expires", "Thu, 01 Jan 1970 00:00:00 GMT")
	c.Writer.Header().Add("Last-Modified", time.Now().UTC().Format(http.TimeFormat))

	c.Next()
}

func Options(c *gin.Context) {
	if c.Request.Method != "OPTIONS" {
		c.Next()
	} else {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Add("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS")
		c.Writer.Header().Add("Allow", "GET,POST,PUT,PATCH,DELETE,OPTIONS")
		c.Writer.Header().Add("Access-Control-Allow-Headers", "authorization, origin, content-type, accept")
		c.Writer.Header().Add("Content-Type", "application/json")
		c.AbortWithStatus(http.StatusOK)
	}
}

func Secure(c *gin.Context) {
	c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Add("X-Frame-Options", "DENY")
	c.Writer.Header().Add("X-Content-Type-Options", "nosniff")
	c.Writer.Header().Add("X-XSS-Protection", "1; mode=block")

	if c.Request.TLS != nil {
		c.Writer.Header().Add("Strict-Transport-Security", "max-age=31536000")
	}
}