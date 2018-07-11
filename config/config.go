package config

import (
	"github.com/spf13/viper"
	"strings"
	"github.com/fsnotify/fsnotify"
	"log"
)

type Config struct {
	FilePath string
}

//初始化配置信息
func (c *Config) initConfig() error {
	if c.FilePath != "" {
		viper.SetConfigFile(c.FilePath)
	} else {
		viper.AddConfigPath("conf")
		viper.SetConfigName("config")
	}

	viper.SetConfigType("yaml")
	//读取匹配的环境变量
	viper.AutomaticEnv()
	viper.SetEnvPrefix("apiserver")
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)

	//解析配置
	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	return nil
}

//热加载	
func (c *Config) watch() {
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		log.Printf("Config file changed: %s", in.Name)
	})
}

func Init(path string) error {
	c := Config{
		FilePath: path,
	}

	if err := c.initConfig(); err != nil {
		return err
	}

	//监听配置文件的修改
	c.watch()

	return nil
}
