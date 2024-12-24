package httpd

import (
	"github.com/labstack/echo/v4"
	"github.com/lupguo/go-shim/shim"
	"github.com/lupguo/wisdom-httpd/app/api"
	"github.com/lupguo/wisdom-httpd/app/application"
	"github.com/lupguo/wisdom-httpd/app/infra/conf"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

// Server Http Server
type Server struct {
	cfg    *conf.Config
	echo   *echo.Echo
	stdLog *log.Logger
	router *Router
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
	stdLog, err := conf.InitLogger(cfg.Log)
	if err != nil {
		return nil, errors.Wrapf(err, "init logger got err")
	}
	e.Use(LogrusMiddleware(stdLog))

	// 配置tmpl渲染器
	render, err := NewWisdomRenderer()
	if err != nil {
		return nil, errors.Wrap(err, "init wisdom template got err")
	}
	e.Renderer = render

	// 路由器配置，通过依赖注入方式实现
	srvImpl := api.NewWisdomImpl(application.NewWisdomApp())
	svr := &Server{
		cfg:    cfg,
		router: InitRouter(e, srvImpl),
		stdLog: stdLog,
	}

	return svr, nil
}

// Start Httpd Server 服务期待
func (s *Server) Start() error {
	log.Infof("svr config: %s", shim.ToJsonString(s, true))

	// 监听点
	addr := s.cfg.Listen
	log.Infof("listen: http://%s", addr)

	// 其他配置
	return s.echo.Start(addr)
}
