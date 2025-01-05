package httpd

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/lupguo/go-shim/shim"
	"github.com/lupguo/wisdom-httpd/app/api"
	"github.com/lupguo/wisdom-httpd/app/infra/conf"
	"github.com/lupguo/wisdom-httpd/internal/log"
	"github.com/lupguo/wisdom-httpd/internal/util"
	"github.com/pkg/errors"
)

// HandleFunc Biz处理方法类型，需要被业务接口去实现
type HandleFunc func(ctx *util.Context, req any) (rsp any, err error)

// RouteHandler 路由处理
type RouteHandler struct {
	Method       string
	URI          string
	HandleFunc   HandleFunc `json:"-"`
	TemplateName string     // 路由模版，纯JSON可以忽略
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

// RegisterRouterHandler 创建一个Web路由
func RegisterRouterHandler(echo *echo.Echo, apiImpl *api.WisdomHandler) (*Router, error) {
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
		r.echo.Router().Add(h.Method, h.URI, warpRouteHandleToEchoHandle(h))
	}

	return nil
}

// 通过注入RouteHandle
func warpRouteHandleToEchoHandle(h *RouteHandler) func(c echo.Context) (err error) {
	return func(c echo.Context) (err error) {
		start := time.Now()
		ctx := util.NewContext(c)

		// RequestBody
		req, reqMap, err := getHTTPBodyEntry(ctx)
		if err != nil {
			return errors.Wrap(err, "getHTTPBodyEntry got err")
		}

		// Biz 请求+响应日志打印
		var rsp any
		defer func(c *util.Context, start time.Time) {
			fields := map[string]any{
				log.FieldError:   err,
				log.FieldElapsed: time.Since(start),
				log.FieldReq:     shim.ToJsonString(reqMap),
				log.FieldRsp:     shim.ToJsonString(rsp),
			}
			log.WithFilesInfoContextf(fields, ctx, "%s", "")
		}(ctx, start)

		// Biz 处理
		rsp, err = h.HandleFunc(ctx, req)
		if err != nil {
			return log.WrapErrorContextf(ctx, err, "h.HandleFunc got err")
		}
		c.Set("rsp", rsp)

		// Biz 结果响应
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
}

// 获取getHTTPBodyEntry请求的信息
func getHTTPBodyEntry(ctx *util.Context) ([]byte, map[string]any, error) {
	if ctx.Request().Method != http.MethodPost {
		return nil, nil, nil
	}

	// reqBody deal
	body, err := io.ReadAll(ctx.Request().Body)
	if err != nil {
		return nil, nil, errors.Wrap(err, "read HTTP request body got err")
	}

	var m map[string]any
	err = json.Unmarshal(body, &m)
	if err != nil {
		return nil, nil, errors.Wrap(err, "router json unmarshal reqbody err")
	}

	return body, m, nil
}
