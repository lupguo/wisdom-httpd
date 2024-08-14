package httpd

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/lupguo/wisdom-httpd/app/api"
	"github.com/lupguo/wisdom-httpd/app/application"
	"github.com/lupguo/wisdom-httpd/app/infra/config"
	"github.com/pkg/errors"
)

// Server Http Server
type Server struct {
	appCfg    *config.AppConfig
	logCfg    *config.LogConfig
	echo      *echo.Echo
	RouterMap RouterMap
}

// NewHttpdServer 创建Httpd服务实例
func NewHttpdServer(configFile string) (*Server, error) {
	// 应用配置
	appCfg, err := config.ParseConfig(configFile)
	if err != nil {
		return nil, errors.Wrapf(err, "parse config file %s got err", configFile)
	}

	// 日志配置
	logCfg, err := config.AppLogConfig()
	if err != nil {
		return nil, errors.Wrapf(err, "get app log config got err")
	}
	if err := logCfg.InitLogger(); err != nil {
		return nil, errors.Wrapf(err, "init logger got err")
	}
	log.SetLevel(logCfg.GetLogLevel())

	// echo 实例
	e := echo.New()
	e.HideBanner = true

	// 服务实例注入
	srvImpl := api.NewImplAPI(application.NewWisdomApp())
	httpdServer := &Server{
		appCfg:    appCfg,
		logCfg:    logCfg,
		echo:      e,
		RouterMap: GetRouterConfigMap(srvImpl),
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
		Format:           s.logCfg.GetLogFormat(),
		CustomTimeFormat: s.logCfg.GetLogTimeFormat(),
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
	routerInit(s.echo, s.RouterMap)
}

// Start Httpd Server 服务期待
func (s *Server) Start() error {
	addr := s.appCfg.Listen
	log.Infof("listen: %v", addr)

	// 其他配置
	return s.echo.Start(addr)
}
