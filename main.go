package main

import (
	"github.com/lupguo/wisdom-httpd/app/infra/conf"
	"github.com/lupguo/wisdom-httpd/httpd"
	"github.com/lupguo/wisdom-httpd/internal/log"
	"github.com/spf13/pflag"
)

// flag
var (
	configFile = pflag.StringP("conf", "c", "./config.yaml", "Application configuration YAML file name")
)

func main() {
	pflag.Parse()

	// 应用配置
	cfg, err := conf.ParseConfig(*configFile)
	if err != nil {
		log.Fatalf("parse config file %s got err: %s", configFile, err)
	}

	// 日志配置
	if err = log.InitServerLog(cfg.LogConfig); err != nil {
		log.Fatalf("init server log got err: %s", err)
	}

	// API服务初始化
	apiHandler, err := NewWisdomAPIHandler()
	if err != nil {
		log.Fatalf("create api handler got err: %s", err)
	}

	svr, err := httpd.NewHttpdServer(cfg, apiHandler)
	if err != nil {
		log.Fatalf("new httpd server got err, %s", err)
	}

	// http server start
	log.Fatalf("http server start fail: %s", svr.Start())
}
