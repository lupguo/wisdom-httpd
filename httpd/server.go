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
	cfg *config.ServerConfig
	e   *echo.Echo
}

// NewHttpdServer 创建Httpd服务实例
func NewHttpdServer(configFile string) (*Server, error) {
	// 应用配置
	cfg, err := config.ParseConfig(configFile)
	if err != nil {
		return nil, errors.Wrapf(err, "parse config file %s got err", configFile)
	}

	// 日志配置
	e := echo.New()
	log.SetLevel(config.GetLogLevel())

	// 服务
	return &Server{
		cfg: cfg,
		e:   e,
	}, nil
}

// ConfigMiddleware 中间件配置
func (s *Server) ConfigMiddleware() {
	s.e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Skipper:          middleware.DefaultSkipper,
		Format:           config.GetLogFormat(),
		CustomTimeFormat: config.GetLogTimeFormat(),
	}))
}

// ConfigRender 渲染配置
func (s *Server) ConfigRender() {
	s.e.Renderer = GetRenderTemplate()
}

// ConfigRoute 路由配置
func (s *Server) ConfigRoute() {
	routerInit(s.e)
}

// Start Httpd Server 服务期待
func (s *Server) Start() {
	addr := s.cfg.Listen
	log.Infof("listen: %v", addr)
	s.e.Logger.Fatal(s.e.Start(addr))
}
