package main

import (
	"fmt"
	"net/http"
	"webapi/route"
	"webapi/util"
	"webapi/version"

	"github.com/facebookgo/grace/gracehttp"

	_ "github.com/go-sql-driver/mysql"
)

const (
	// DefaultconfigFile 配置文件位置
	DefaultconfigFile = "config/webapi/webapi.ini"
	// DefauleHTTPPort 默认端口
	DefauleHTTPPort = ":9090"
	// DefaultLogConfigFile 默认日志配置文件路径
	DefaultLogConfigFile = "config/webapi/webapi_seelog.xml"
)

func initConfig() {
	err := util.ProcessConfigFile(DefaultconfigFile)
	if err != nil {
		fmt.Println(err)
	}
}

func initMysql() {
	err := util.InitMysql()
	if err != nil {
		fmt.Println(err)
	}
}
func initLog() {
	err := util.SetLogConfig(DefaultLogConfigFile)
	util.Logger.Flush()
	if err != nil {
		fmt.Println(err)
	}
}
func init() {
	initLog()
	initConfig()
	initMysql()
}

func main() {
	util.Logger.Info("Starting user page server, version: ", version.Version)
	gracehttp.Serve(
		&http.Server{Addr: DefauleHTTPPort, Handler: route.NewRouter()},
	)
}
