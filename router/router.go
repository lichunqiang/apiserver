package router

import (
	"github.com/gin-gonic/gin"
	"apiserver/handler/sd"
	"apiserver/handler/user"
	w "apiserver/router/middlewares"
	"net/http"
)

func InitRoute(g *gin.Engine, wm ...gin.HandlerFunc) *gin.Engine {
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

	g.POST("/login", user.Login)

	u := g.Group("/v1/users")
	//u.Use(w.AuthMiddleware())
	{
		//u.GET("", user.List)
		u.POST("", user.Create)
		//u.PUT("/:id", user.Update)
		u.DELETE("/:id", user.Delete)
		u.GET("/:id", user.Get)
	}

	sdg := g.Group("/sd")
	{
		sdg.GET("/health", sd.HealthCheck)
		//sdg.GET("/disk", sd.DiskCheck)
		sdg.GET("/cpu", sd.CPUCheck)
	}

	return g
}
