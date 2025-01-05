package httpd

import (
	"github.com/labstack/echo/v4"
	"github.com/lupguo/go-shim/shim"
	"github.com/lupguo/wisdom-httpd/app/api"
	"github.com/lupguo/wisdom-httpd/app/application"
	"github.com/lupguo/wisdom-httpd/app/infra/conf"
	"github.com/lupguo/wisdom-httpd/internal/log"
	"github.com/pkg/errors"
)

// Server Http Server
type Server struct {
	Cfg    *conf.Config
	Router *Router
	echo   *echo.Echo
}

// NewHttpdServer 创建Httpd服务实例
func NewHttpdServer(configFile string) (*Server, error) {
	// echo 实例
	e := echo.New()
	e.HideBanner = true

	// 应用配置
	cfg, err := conf.ParseConfig(configFile)
	if err != nil {
		return nil, errors.Wrapf(err, "parse config file %s got err", configFile)
	}

	// 日志配置
	err = log.NewServerLog(cfg.LogConfig)
	if err != nil {
		return nil, errors.Wrapf(err, "build logger got err")
	}
	// e.Use(LogrusMiddleware(slog))

	// 配置tmpl渲染器
	render, err := NewWisdomRenderer()
	if err != nil {
		return nil, errors.Wrap(err, "build wisdom template got err")
	}
	e.Renderer = render

	// 路由器配置，通过依赖注入方式实现
	srvImpl := api.NewWisdomImpl(application.NewWisdomApp())
	router, err := RegisterRouterHandler(e, srvImpl)
	if err != nil {
		return nil, errors.Wrap(err, "init router got err")
	}
	svr := &Server{
		Cfg:    cfg,
		echo:   e,
		Router: router,
	}

	return svr, nil
}

// Start Httpd Server 服务期待
func (s *Server) Start() error {
	log.Infof("svr config: %s", shim.ToJsonString(s, true))

	// 监听点
	addr := s.Cfg.Listen
	log.Infof("listen: http://%s", addr)

	// 其他配置
	return s.echo.Start(addr)
}
