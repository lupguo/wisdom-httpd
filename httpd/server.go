package httpd

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/lupguo/wisdom-httpd/config"
	"github.com/pkg/errors"
)

// Server Http Server
type Server struct {
	cfg      *config.AppConfig
	echo     *echo.Echo
	logLevel log.Lvl
}

// NewHttpdServer 创建Httpd服务实例
func NewHttpdServer(configFile string) (*Server, error) {
	// 应用配置
	cfg, err := config.ParseConfig(configFile)
	if err != nil {
		return nil, errors.Wrapf(err, "parse config file %s got err", configFile)
	}

	// 服务
	httpdServer := &Server{
		cfg:      cfg,
		echo:     echo.New(),
		logLevel: config.GetLogLevel(),
	}
	// 渲染
	if err = httpdServer.InitRenderConfig(); err != nil {
		return nil, err
	}

	// 中间件
	httpdServer.InitMiddlewareConfig()

	// 路由
	httpdServer.InitRouteConfig()

	return httpdServer, nil
}

// InitMiddlewareConfig 中间件配置
func (s *Server) InitMiddlewareConfig() {
	s.echo.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Skipper:          middleware.DefaultSkipper,
		Format:           config.GetLogFormat(),
		CustomTimeFormat: config.GetLogTimeFormat(),
	}))
}

// InitRenderConfig 渲染配置
func (s *Server) InitRenderConfig() error {
	templateRender, err := InitParseWisdomTemplate()
	if err != nil {
		return errors.Wrap(err, "init wisdom template got err")
	}
	s.echo.Renderer = templateRender
	return nil
}

// InitRouteConfig 路由配置
func (s *Server) InitRouteConfig() {
	routerInit(s.echo)
}

// Start Httpd Server 服务期待
func (s *Server) Start() {
	addr := s.cfg.Listen

	// 日志配置
	log.SetLevel(s.logLevel)
	log.Infof("listen: %v", addr)
	s.echo.HideBanner = true
	s.echo.Logger.Fatal(s.echo.Start(addr))
}
