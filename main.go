package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lichunqiang/apiserver/router"
	"github.com/spf13/pflag"
	"github.com/lichunqiang/apiserver/config"
	"github.com/spf13/viper"
	"net/http"
	"log"
	"time"
	"errors"
)

func initServer() *gin.Engine {
	g := gin.New()

	middlewares := []gin.HandlerFunc{}

	router.InitRoute(g, middlewares...)

	return g
}

var (
	filePath = pflag.StringP("config", "c", "", "apiserver config file path.")
)

func main() {
	pflag.Parse()

	if err := config.Init(*filePath); err != nil {
		panic(err)
	}

	//run mode
	gin.SetMode(viper.GetString("runmode"))

	r := initServer()

	go func() {
		if err := pingServer(); err != nil {
			log.Fatal("No response", err)
		}
		log.Print("Apiserver startted successfully.")
	}()

	r.Run(viper.GetString("addr"))
}

func pingServer() error {
	for i := 0; i < viper.GetInt("max_ping_count"); i++ {
		resp, err := http.Get(viper.GetString("url") + "/sd/health")

		if err == nil && resp.StatusCode == http.StatusOK {
			return nil
		}

		log.Print("Waiting for the router, retry in 1 second")
		time.Sleep(time.Second)
	}

	return errors.New("Cannot connect to apiserver")
}
