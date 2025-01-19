package httpd

import (
	"github.com/labstack/echo/v4"
	"github.com/lupguo/go-shim/shim"
	"github.com/lupguo/wisdom-httpd/app/api"
	"github.com/lupguo/wisdom-httpd/app/infra/conf"
	"github.com/lupguo/wisdom-httpd/internal/log"
	"github.com/pkg/errors"
)

// HttpdServer Http HttpdServer
type HttpdServer struct {
	Cfg    *conf.Config
	Router *Router
	echo   *echo.Echo
}

// NewHttpdServer 创建Httpd服务实例
func NewHttpdServer(cfg *conf.Config, apiHandler *api.WisdomHandler) (*HttpdServer, error) {
	// Echo框架渲染器
	render, err := NewHTMLRenderer()
	if err != nil {
		return nil, errors.Wrap(err, "build wisdom template got err")
	}

	// Echo 实例
	e := echo.New()
	e.HideBanner = true
	e.Renderer = render

	// Echo中间件
	InitMiddlewareConfig(e)

	// 路由器配置，通过依赖注入方式实现
	router, err := RegisterRouterHandler(e, apiHandler)
	if err != nil {
		return nil, errors.Wrap(err, "init router got err")
	}
	svr := &HttpdServer{
		Cfg:    cfg,
		echo:   e,
		Router: router,
	}

	return svr, nil
}

// Start Httpd HttpdServer 服务期待
func (s *HttpdServer) Start() error {
	log.Infof("svr config: %s", shim.ToJsonString(s, true))

	// 监听点
	addr := s.Cfg.Listen
	log.Infof("listen: http://%s", addr)

	// 其他配置
	return s.echo.Start(addr)
}
