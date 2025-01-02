package httpd

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/lupguo/wisdom-httpd/app/api"
	"github.com/lupguo/wisdom-httpd/app/infra/conf"
	"github.com/lupguo/wisdom-httpd/internal/util"
	"github.com/pkg/errors"
)

// HandleFunc Biz处理方法类型，需要被业务接口去实现
type HandleFunc func(ctx *util.Context, req any) (rsp any, err error)

// RouteHandler 路由处理
type RouteHandler struct {
	Method       string
	URI          string
	HandleFunc   HandleFunc
	TemplateName string // 路由模版，纯JSON可以忽略
}

// registerRoutes 路由配置，通过apiImpl实例注入
func registerRoutes(api *api.WisdomHandler) []*RouteHandler {
	return []*RouteHandler{
		{"GET", "/", api.Index, "index.tmpl"},
		{"GET", "/wisdom", api.GetOneWisdom, "wisdom.tmpl"},

		// 2024.12.31
		{"POST", "/wisdom", api.SaveWisdom, ""},
	}
}

// InitRouter 创建一个Web路由
func InitRouter(echo *echo.Echo, apiImpl *api.WisdomHandler) (*Router, error) {
	r := &Router{
		echo:          echo,
		RouteHandlers: registerRoutes(apiImpl),
	}

	if err := r.build(); err != nil {
		return nil, errors.Wrap(err, "build router got err")
	}

	return r, nil
}

// Router Web路由
type Router struct {
	RouteHandlers []*RouteHandler
	echo          *echo.Echo
}

// build 处理路由注册
func (r *Router) build() error {
	// 静态路由
	r.echo.Static("/", conf.PublicPath())

	// 动态路由
	for _, h := range r.RouteHandlers {
		// warp成Echo处理处理
		echoHttpHandleFunc := func(c echo.Context) error {
			// todo set req as the second param to h.HandleFun
			req := map[string]any{}
			_ = c.Bind(&req)

			rsp, err := h.HandleFunc(util.NewContext(c), req)
			if err != nil {
				return errors.Wrap(err, "h.HandleFunc got err")
			}

			// json or html render
			switch c.Request().Header.Get("Content-Type") {
			case "application/json":
				return c.JSON(http.StatusOK, rsp)
			case "text/html":
				if h.TemplateName == "" {
					return errors.Errorf("URI[%s] html template name is empty", h.URI)
				}
				return c.Render(http.StatusOK, h.TemplateName, rsp)
			}

			return nil
		}

		// 分组添加
		r.echo.Router().Add(h.Method, h.URI, echoHttpHandleFunc)
	}

	return nil
}
