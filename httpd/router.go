package httpd

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/lupguo/wisdom-httpd/app/api"
	"github.com/lupguo/wisdom-httpd/app/infra/conf"
	"github.com/pkg/errors"
)

// HandleFunc 处理方法
type HandleFunc func(c echo.Context, req any) (rsp any, err error)

// RouteHandler 路由处理
type RouteHandler struct {
	Method       string
	URI          string
	TemplateName string
	HandleFunc   HandleFunc `json:"-"`
}

// registerRoutes 路由配置，通过apiImpl实例注入
func registerRoutes(apiImpl *api.SrvImpl) []*RouteHandler {
	return []*RouteHandler{
		{"GET", "/", "index.tmpl", apiImpl.IndexHandler},
		{"GET", "/wisdom", "wisdom.tmpl", apiImpl.WisdomHandler},
	}
}

// InitRouter 创建一个Web路由
func InitRouter(echo *echo.Echo, apiImpl *api.SrvImpl) *Router {
	return &Router{
		echo:          echo,
		routeHandlers: registerRoutes(apiImpl),
	}
}

// Router Web路由
type Router struct {
	echo          *echo.Echo
	routeHandlers []*RouteHandler
}

// HandleConfig 处理路由注册
func (r *Router) HandleConfig() error {
	// 静态路由
	r.echo.Static("/", conf.PublicPath())

	// 动态路由
	for _, h := range r.routeHandlers {
		// warp成使用注册的apiImpl实例方法处理
		eHandleFn := func(c echo.Context) error {
			rsp, err := h.HandleFunc(c, c.Request().Body)
			if err != nil {
				return errors.Wrap(err, "h.HandleFunc got err")
			}

			// json or html render
			switch c.Request().Header.Get("Content-Type") {
			case "application/json":
				return c.JSON(http.StatusOK, rsp)
			case "text/html":
				return c.Render(http.StatusOK, h.TemplateName, rsp)
			}

			return nil
		}

		// 分组添加
		r.echo.Router().Add(h.Method, h.URI, eHandleFn)
	}

	return nil
}
