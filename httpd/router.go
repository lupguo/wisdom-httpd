package httpd

import (
	"github.com/labstack/echo/v4"
	"github.com/lupguo/wisdom-httpd/app/api"
	"github.com/lupguo/wisdom-httpd/app/domain/entity"
	"github.com/lupguo/wisdom-httpd/app/infra/config"
	"github.com/pkg/errors"
)

type HandlerFunc func(c echo.Context) (data *entity.WebPageData, err error)

type RouteHandler struct {
	Method      string
	URI         string
	HandlerFunc HandlerFunc
}

type RouterMap map[string][]*RouteHandler

// GetRouterConfigMap 路由配置，通过apiImpl实例注入
func GetRouterConfigMap(apiImpl *api.SrvImpl) RouterMap {
	return RouterMap{
		"web": {
			{"GET", "/", apiImpl.IndexHandler},
			{"GET", "/wisdom", apiImpl.WisdomHandler},
		},
		"json": {
			{"GET", "/wisdom", apiImpl.WisdomHandler},
		},
	}
}

// 路由配置
func routerInit(e *echo.Echo, routerMap RouterMap) {
	// 静态路由
	e.Static("/", config.PublicPath())

	// 动态路由
	var prefix string
	for key, handlers := range routerMap {
		rg := &echo.Group{}
		switch key {
		case "web":
			rg = e.Group(prefix, []echo.MiddlewareFunc{WebResponseMiddleware}...)
		case "json":
			prefix = "api"
			rg = e.Group(prefix, []echo.MiddlewareFunc{JSONResponseMiddleware}...)
		}

		// router group处理
		for _, h := range handlers {
			eFn := func(c echo.Context) error {
				// 使用注册的apiImpl实例方法处理
				rsp, err := h.HandlerFunc(c)
				if err != nil {
					return errors.Wrap(err, "h.HandlerFunc got err")
				}
				c.Set("data", rsp)
				return nil
			}
			rg.Add(h.Method, h.URI, eFn)
		}

	}
}
