package httpd

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/lupguo/go-shim/shim"
	"github.com/lupguo/wisdom-httpd/app/api"
	"github.com/lupguo/wisdom-httpd/app/application"
	"github.com/lupguo/wisdom-httpd/app/infra/config"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

// Server Http Server
type Server struct {
	appCfg    *config.AppConfig
	echo      *echo.Echo
	RouterMap RouterMap
}

// NewHttpdServer 创建Httpd服务实例
func NewHttpdServer(configFile string) (*Server, error) {
	// echo 实例
	e := echo.New()
	e.HideBanner = true

	// 应用配置
	appCfg, err := config.ParseConfig(configFile)
	if err != nil {
		return nil, errors.Wrapf(err, "parse config file %s got err", configFile)
	}

	// 日志配置
	stdLog, err := config.InitLogger(e, appCfg.Log)
	if err != nil {
		return nil, errors.Wrapf(err, "init logger got err")
	}

	// 注入到Echo中间件中
	var header []string
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		Skipper:        nil,
		BeforeNextFunc: nil,
		LogValuesFunc: func(c echo.Context, values middleware.RequestLoggerValues) error {
			// req, _ := io.ReadAll(c.Request().Body)
			traceId := uuid.NewString()
			stdLog.WithFields(log.Fields{
				"trace_id":  traceId,
				"uri":       values.URI,
				"status":    values.Status,
				"method":    values.Method,
				"remote_ip": values.RemoteIP,
				"host":      values.Host,
				"rt":        values.Latency,
				"ua":        values.UserAgent,
				"rsp_size":  values.ResponseSize,
			}).Infof("request")
			// .Info("req")
			// Infof("qry=>%s, reqbody=>%s, values=>%s", c.QueryString(), req, shim.ToJsonString(values, false))

			// stdLog.Infof("reqmsg=>%s", shim.ToJsonString(values, false))

			return nil
		},
		HandleError:      true,
		LogLatency:       true,
		LogProtocol:      true,
		LogRemoteIP:      true,
		LogHost:          true,
		LogMethod:        true,
		LogURI:           true,
		LogURIPath:       true,
		LogRoutePath:     true,
		LogRequestID:     true,
		LogReferer:       true,
		LogUserAgent:     true,
		LogStatus:        true,
		LogError:         true,
		LogContentLength: true,
		LogResponseSize:  true,
		LogHeaders:       header,
		LogQueryParams:   nil,
		LogFormValues:    nil,
	}))

	// 服务实例注入
	srvImpl := api.NewImplAPI(application.NewWisdomApp())
	httpdServer := &Server{
		appCfg:    appCfg,
		echo:      e,
		RouterMap: GetRouterConfigMap(srvImpl),
	}
	log.Infof("appConfig: %s", shim.ToJsonString(appCfg, false))

	// 路由
	httpdServer.InitRouteConfig()

	// 提前页面渲染
	if err = httpdServer.InitRenderConfig(); err != nil {
		return nil, err
	}

	return httpdServer, nil
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
	log.Infof("listen: http://%s", addr)

	// 其他配置
	return s.echo.Start(addr)
}
