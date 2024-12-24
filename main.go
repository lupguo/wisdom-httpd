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
	svr, err := httpd.NewHttpdServer(*configFile)
	if err != nil {
		log.Fatalf("new httpd server got err %s", err)
	}

	// http server start
	log.Fatalf("http server start fail: %s", svr.Start())
}
