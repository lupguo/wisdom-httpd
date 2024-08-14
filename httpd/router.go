package httpd

import (
	"github.com/labstack/echo/v4"
	"github.com/lupguo/wisdom-httpd/app/api"
	"github.com/lupguo/wisdom-httpd/app/domain/entity"
	"github.com/lupguo/wisdom-httpd/app/infra/config"
)

type HandlerFunc func(c echo.Context) (data *entity.WebPageDataRsp, err error)

type RouteHandler struct {
	Method      string
	URI         string
	HandlerFunc HandlerFunc
}

// 路由配置
func routerInit(e *echo.Echo) {
	// 静态路由
	initStaticRouter(e)

	// web路由
	initWebRouter(e)

	// api路由
	initAPIRouter(e)

}

func initAPIRouter(e *echo.Echo) {
	apiRouterHandlers := []*RouteHandler{
		{"GET", "/wisdom", api.WisdomHandler},
	}
	apiGroups := e.Group("/api", []echo.MiddlewareFunc{
		JSONResponseMiddleware,
	}...)
	for _, h := range apiRouterHandlers {
		// wrap函数
		eFn := func(c echo.Context) error {
			got, err := h.HandlerFunc(c)
			if err != nil {
				return err
			}
			c.Set("data", got)
			return nil
		}

		apiGroups.Add(h.Method, h.URI, eFn)
	}
}

// web路由
func initWebRouter(e *echo.Echo) {
	htmlRouterHandlers := []*RouteHandler{
		{"GET", "/", api.IndexHandler},
		{"GET", "/wisdom", api.WisdomHandler},
	}
	htmlGroups := e.Group("", []echo.MiddlewareFunc{
		HTMLResponseMiddleware,
	}...)
	for _, h := range htmlRouterHandlers {
		// wrap函数
		eFn := func(c echo.Context) error {
			got, err := h.HandlerFunc(c)
			if err != nil {
				return err
			}
			c.Set("data", got)
			return nil
		}
		htmlGroups.Add(h.Method, h.URI, eFn)
	}
}

// static路由
func initStaticRouter(e *echo.Echo) {
	e.Static("/", config.PublicPath())
}
