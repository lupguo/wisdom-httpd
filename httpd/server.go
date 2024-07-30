package httpd

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/lupguo/wisdom-httpd/app/api"
	"github.com/lupguo/wisdom-httpd/config"
	"github.com/pkg/errors"
)

// Server Http Server
type Server struct {
	cfg  *config.Config
	echo *echo.Echo
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
		cfg:  cfg,
		echo: e,
	}, nil
}

// ConfigMiddleware 中间件配置
func (s *Server) ConfigMiddleware() {
	s.echo.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Skipper:          middleware.DefaultSkipper,
		Format:           config.GetLogFormat(),
		CustomTimeFormat: config.GetLogTimeFormat(),
	}))
}

// ConfigRender 渲染配置
func (s *Server) ConfigRender() {
	s.echo.Renderer = InitParseWisdomTemplate()
}

// ConfigRoute 路由配置
func (s *Server) ConfigRoute() {
	routerInit(s.echo)
}

// Start Httpd Server 服务期待
func (s *Server) Start() {
	addr := s.cfg.Listen
	log.Infof("listen: %v", addr)
	s.echo.Logger.Fatal(s.echo.Start(addr))
}

// 路由配置
func routerInit(e *echo.Echo) {
	// 静态资源配置
	e.Static("/", config.GetAssetPath())

	// 首页处理
	e.GET("/", api.IndexHandler)
	e.GET("/error", api.ErrorHandler)

	// wisdom接口处理
	e.GET("/wisdom", api.WisdomHandler)

	// 调试statusCode处理
	e.GET("/code", api.CodeHandler)

	// handler - 文件下载mock
	e.GET("/files/upload", api.UploadHandler)
	e.GET("/files/download", api.DownloadHandler)
}
