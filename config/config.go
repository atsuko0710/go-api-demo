package config

import (
	"log"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Config struct {
	Name string
}

func Init(cfg string) error {
	c := Config {
		Name: cfg,
	}

	// 初始化配置
	if err := c.initConfig(); err != nil {
		return err
	}

	c.watchCOnfig()
	return nil
}

// 初始化配置
func (c *Config) initConfig() error {
	if c.Name != "" {
		// 指定文件则解析指定的配置文件
		viper.SetConfigFile(c.Name)
	} else {
		// 没有指定文件就指定默认的配置文件路径
		viper.AddConfigPath("./conf")
		viper.SetConfigName("config")
	}
	// 设置文件格式
	viper.SetConfigType("yaml")
	// 读取配置的环境变量
	viper.AutomaticEnv()
	// 设置环境变量前缀
	// viper.SetEnvPrefix("APISERVER")
	// 如果配置文件中 _ 形式换成 .
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)

	// 读取配置信息
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	return nil
}

// 监控配置文件变化并热加载程序
func (c *Config) watchCOnfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Printf("Config file changed: %s", e.Name)
	})
}
