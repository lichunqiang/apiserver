package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/lichunqiang/apiserver/handler/sd"
	w "github.com/lichunqiang/apiserver/router/middlewares"
)

func InitRoute(g *gin.Engine, wm ...gin.HandlerFunc) *gin.Engine  {
	g.Use(gin.Recovery())
	g.Use(gin.Logger())

	g.Use(w.NoCache, w.Options, w.Secure)

	g.Use(wm...)

	g.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})

	//404
	g.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "404",
		})
	})

	sdg := g.Group("/sd")
	{
		sdg.GET("/health", sd.HealthCheck)
		//sdg.GET("/disk", sd.DiskCheck)
		sdg.GET("/cpu", sd.CPUCheck)
	}

	return g
}