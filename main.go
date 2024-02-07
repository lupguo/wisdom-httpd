package main

import (
	"log"

	"github.com/lupguo/wisdom-httpd/httpd"
	"github.com/spf13/pflag"
)

// flag
var (
	configFile = pflag.StringP("conf", "c", "./config.yaml", "Application configuration YAML file name")
)

func main() {
	pflag.Parse()
	server, err := httpd.NewHttpdServer(*configFile)
	if err != nil {
		log.Fatalf("new httpd server got err %s", err)
	}

	// httpd server config
	server.ConfigMiddleware()
	server.ConfigRender()
	server.ConfigRoute()

	// http server start
	server.Start()
}
