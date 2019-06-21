package util

import (
	"github.com/Unknwon/goconfig"
)

var (
	cfg *Config
)

// NewConfig 新建配置对象
func NewConfig() *Config {
	c := new(Config)
	return c
}

// Config config
type Config struct {
	Mysql MysqlOpts `json:"mysql,omitempty"`
}

// Configini 配置文初始化对象
var Configini *goconfig.ConfigFile

// MysqlOpts mysql 的结构
type MysqlOpts struct {
	Host     string `json:"host,omitempty"`
	User     string `json:"user,omitempty"`
	Password string `json:"password,omitempty"`
	Port     int    `json:"port,omitempty"`
}

// ProcessConfigFile 传入配置文件
func ProcessConfigFile(configFile string) error {
	if configFile == "" {
		return nil
	}
	config, err := goconfig.LoadConfigFile(configFile)
	if err != nil {
		return err
	}

	cfg = new(Config)
	cfg.Mysql.Host = config.MustValue("mysql", "host", "127.0.0.1")
	cfg.Mysql.User = config.MustValue("mysql", "user", "root")
	cfg.Mysql.Password = config.MustValue("mysql", "password", "")
	cfg.Mysql.Port = config.MustInt("mysql", "port")

	Configini = config

	return nil
}
