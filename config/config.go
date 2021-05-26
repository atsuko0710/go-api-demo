package config

import (
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/lexkong/log"
	"github.com/spf13/viper"
)

type Config struct {
	Name string
}

func Init(cfg string) error {
	c := Config{
		Name: cfg,
	}

	// 初始化配置
	if err := c.initConfig(); err != nil {
		return err
	}

	c.initLog()
	
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
		log.Infof("Config file changed: %s", e.Name)
	})
}

// 初始化日志配置
func (c *Config) initLog() {
	passLagerCfg := log.PassLagerCfg{
		Writers:        viper.GetString("log.writers"),
		LoggerLevel:    viper.GetString("log.logger_level"),
		LoggerFile:     viper.GetString("log.logger_file"),
		LogFormatText:  viper.GetBool("log.log_format_text"),
		RollingPolicy:  viper.GetString("log.rollingPolicy"),
		LogRotateDate:  viper.GetInt("log.log_rotate_date"),
		LogRotateSize:  viper.GetInt("log.log_rotate_size"),
		LogBackupCount: viper.GetInt("log.log_backup_count"),
	}
	log.InitWithConfig(&passLagerCfg)
}
