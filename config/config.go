package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/lexkong/log"
	"github.com/spf13/viper"
	"strings"
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
	viper.SetEnvPrefix("github.com/lichunqiang/apiserver")
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)

	//解析配置
	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	return nil
}

func (c *Config) initLog() {
	cfg := log.PassLagerCfg{
		Writers:        viper.GetString("log.writers"),
		LoggerLevel:    viper.GetString("log.logger_level"),
		LoggerFile:     viper.GetString("log.logger_file"),
		LogFormatText:  viper.GetBool("log.log_format_text"),
		RollingPolicy:  viper.GetString("log.rollingPolicy"),
		LogRotateDate:  viper.GetInt("log.log_rotate_date"),
		LogRotateSize:  viper.GetInt("log.log_rotate_size"),
		LogBackupCount: viper.GetInt("log.log_backup_count"),
	}

	log.InitWithConfig(&cfg)
}

//热加载
func (c *Config) watch() {
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		log.Infof("Config file changed: %s", in.Name)
	})
}

func Init(path string) error {
	c := Config{
		FilePath: path,
	}

	if err := c.initConfig(); err != nil {
		return err
	}

	//init log
	c.initLog()

	//监听配置文件的修改
	c.watch()

	return nil
}
