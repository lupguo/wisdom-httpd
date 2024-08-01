package httpd

import (
	"github.com/labstack/echo/v4"
	"github.com/lupguo/wisdom-httpd/app/api"
	"github.com/lupguo/wisdom-httpd/config"
)

type RouteHandler struct {
	Method  string
	URI     string
	Handler echo.HandlerFunc
}

// 路由配置
func routerInit(e *echo.Echo) {
	// static路由
	e.Static("/", config.PublicPath())

	// web路由
	htmlRouterHandlers := []*RouteHandler{
		// {"GET", "/index.html", api.IndexHandler},
		{"GET", "/wisdom.html", api.WisdomHandler},
	}
	htmlGroups := e.Group("/", []echo.MiddlewareFunc{
		HTMLResponseMiddleware,
	}...)
	for _, h := range htmlRouterHandlers {
		htmlGroups.Add(h.Method, h.URI, h.Handler)
	}

	// api路由
	apiRouterHandlers := []*RouteHandler{
		{"GET", "/wisdom/rand", api.WisdomRandHandler},
		// {"GET", "/code/show", api.CodeHandler},
		// {"GET", "/files/upload", api.UploadHandler},
		// {"GET", "/files/download", api.UploadHandler},
	}
	apiGroups := e.Group("/api", []echo.MiddlewareFunc{
		JSONResponseMiddleware,
	}...)
	for _, h := range apiRouterHandlers {
		apiGroups.Add(h.Method, h.URI, h.Handler)
	}

}
