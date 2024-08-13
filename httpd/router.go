package httpd

import (
	"github.com/labstack/echo/v4"
	"github.com/lupguo/wisdom-httpd/app/api"
	"github.com/lupguo/wisdom-httpd/app/domain/entity"
	"github.com/lupguo/wisdom-httpd/config"
)

type HandlerFunc func(c echo.Context) (data *entity.WebPageData, err error)

type RouteHandler struct {
	Method      string
	URI         string
	HandlerFunc HandlerFunc
}

// 路由配置
func routerInit(e *echo.Echo) {
	// static路由
	e.Static("/", config.PublicPath())

	// web路由
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

	// api路由
	apiRouterHandlers := []*RouteHandler{
		{"GET", "/wisdom", api.WisdomHandler},
		// {"GET", "/code/show", api.CodeHandler},
		// {"GET", "/files/upload", api.UploadHandler},
		// {"GET", "/files/download", api.UploadHandler},
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
