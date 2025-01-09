//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/lupguo/wisdom-httpd/app/api"
	"github.com/lupguo/wisdom-httpd/app/application"
	"github.com/lupguo/wisdom-httpd/app/domain/repos"
	"github.com/lupguo/wisdom-httpd/app/domain/service"
	"github.com/lupguo/wisdom-httpd/app/infra/dbs"
)

// api
var apiSet = wire.NewSet(
	api.NewWisdomHandlerImpl,
)

// app
var appSet = wire.NewSet(
	wire.Bind(new(application.IAppWisdom), new(*application.WisdomApp)),
	application.NewWisdomApp,
)

// srv
var srvSet = wire.NewSet(
	wire.Bind(new(service.IServiceWisdom), new(*service.WisdomService)),
	service.NewWisdomService,
)

// repos
var infraSet = wire.NewSet(
	wire.Bind(new(repos.IReposWisdomDB), new(*dbs.WisdomDB)),
	dbs.NewWisdomDB,
)

// NewWisdomAPIHandler 使用 Wire 生成 api.WisdomHandler
func NewWisdomAPIHandler() (*api.WisdomHandler, error) {
	wire.Build(apiSet, appSet, srvSet, infraSet)
	return &api.WisdomHandler{}, nil
}
