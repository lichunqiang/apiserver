package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/lichunqiang/apiserver/config"
	"github.com/lichunqiang/apiserver/model"
	"github.com/lichunqiang/apiserver/pkg/version"
	"github.com/lichunqiang/apiserver/router"
	"github.com/lichunqiang/apiserver/router/middlewares"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"net/http"
	"os"
	"time"
)

func initServer() *gin.Engine {
	g := gin.New()

	router.InitRoute(
		g,
		middlewares.Logging(),
		middlewares.RequestId(),
	)

	return g
}

var (
	filePath = pflag.StringP("config", "c", "", "apiserver config file path.")

	showVersion = pflag.BoolP("version", "v", false, "show version info")
)

func main() {
	pflag.Parse()

	if *showVersion {
		versionInfo := version.Get()
		marshalled, err := json.MarshalIndent(&versionInfo, "", "	")
		if err != nil {
			fmt.Printf("%v\n", err)
			os.Exit(1)
		}
		fmt.Println(string(marshalled))
		return
	}

	if err := config.Init(*filePath); err != nil {
		panic(err)
	}

	//run mode
	gin.SetMode(viper.GetString("runmode"))

	model.DB.Init()
	defer model.DB.Close()

	r := initServer()

	go func() {
		if err := pingServer(); err != nil {
			log.Fatal("No response", err)
		}
		log.Debug("Apiserver startted successfully.")
	}()

	//https server
	cert := viper.GetString("tls.cert")
	key := viper.GetString("tls.key")
	if cert != "" && key != "" {
		go func() {
			log.Infof("Start to listening the incoming requests on https address %s", viper.GetString("tls.addr"))
			log.Info(http.ListenAndServeTLS(viper.GetString("tls.addr"), cert, key, r).Error())
		}()
	}

	r.Run(viper.GetString("addr"))
}

func pingServer() error {
	for i := 0; i < viper.GetInt("max_ping_count"); i++ {
		resp, err := http.Get(viper.GetString("url") + "/sd/health")

		if err == nil && resp.StatusCode == http.StatusOK {
			return nil
		}

		log.Info("Waiting for the router, retry in 1 second")
		time.Sleep(time.Second)
	}

	return errors.New("Cannot connect to apiserver")
}
