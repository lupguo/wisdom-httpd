package httpd

import (
	"context"
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

// APIHandleFunc Biz处理方法类型，需要被业务接口去实现
type APIHandleFunc func(ctx context.Context, reqData []byte) (rsp any, err error)

// RouteHandler 路由处理
type RouteHandler struct {
	Method        string
	URI           string
	APIHandleFunc APIHandleFunc `json:"-"`
	TemplateName  string        // 路由模版，纯JSON可以忽略
}

// registerRoutes 路由配置，通过apiImpl实例注入
func registerRoutes(api *api.WisdomHandler) []*RouteHandler {
	return []*RouteHandler{
		{"GET", "/", api.Index, "index.tmpl"},

		// 2024.12.31
		{"GET", "/wisdom", api.GetOneWisdom, "wisdom.tmpl"},

		{"GET", "/api/wisdom", api.GetOneWisdom, ""},
		{"POST", "/api/wisdom", api.SaveWisdom, ""},
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
	r.echo.Static("/dist", conf.PublicPath())

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
		reqData, reqMap, err := getHTTPReqEntry(ctx)
		if err != nil {
			return errors.Wrap(err, "getHTTPReqEntry got err")
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
			log.WithContext(ctx).WithFields(fields).Print()
		}(ctx, start)

		// Biz 处理
		rsp, err = h.APIHandleFunc(ctx, reqData)
		if err != nil {
			return errors.Wrapf(err, "api handler[%s] got err", h.URI)
		}

		// header log
		// log.InfoContextf(ctx, "header => %s ", shim.ToJsonString(ctx.Request().Header))

		// Biz 结果响应
		if c.Request().Header.Get("Content-Type") == "application/json" || h.TemplateName == "" {
			return c.JSON(http.StatusOK, rsp)
		} else {
			// text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7
			return c.Render(http.StatusOK, h.TemplateName, rsp)
		}

	}
}

// 获取getHTTPBodyEntry请求的信息
func getHTTPReqEntry(ctx *util.Context) (reqData []byte, reqMap map[string]any, err error) {
	var m map[string]any
	switch ctx.Request().Method {
	case http.MethodGet:
		qryParams := ctx.QueryParams()
		if len(qryParams) == 0 {
			return nil, nil, nil
		}

		m = make(map[string]any, len(qryParams))
		for k, v := range qryParams {
			m[k] = v
		}

		mar, err := json.Marshal(m)
		if err != nil {
			return nil, nil, errors.Wrap(err, "json marshal got err")
		}

		return mar, m, nil
	case http.MethodPost:
		reqBody, err := io.ReadAll(ctx.Request().Body)
		if err != nil {
			return nil, nil, errors.Wrap(err, "read HTTP request body got err")
		}

		if reqBody != nil {
			err = json.Unmarshal(reqBody, &m)
			if err != nil {
				return nil, nil, errors.Wrap(err, "router json unmarshal reqbody err")
			}
		}

		return reqBody, m, nil
	}

	return nil, nil, errors.Errorf("invalid http method: %s", ctx.Request().Method)
}
