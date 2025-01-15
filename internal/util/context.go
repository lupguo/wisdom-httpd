package util

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

const (
	TraceID = `trace_id`
	EchoCtx = `echo_ctx`
)

// Meta 元信息
type Meta struct {
	TraceId string
}

// Context Wisdom HTTP的上下文
type Context struct {
	echo.Context
	meta map[string]any
	err  error
}

func (c *Context) Value(key any) any {
	if v, ok := c.meta[key.(string)]; ok {
		return v
	}
	return nil
}

func (c *Context) Deadline() (deadline time.Time, ok bool) {
	return
}

func (c *Context) Done() <-chan struct{} {
	return nil
}

func (c *Context) Err() error {
	return c.err
}

// TraceId trace_id
func (c *Context) TraceId() string {
	if v, ok := c.Value(TraceID).(string); ok {
		return v
	}
	return ""
}

// GetHTTPReqEntry 从Ctx获取HTTP的GET或Post参数到
func (c *Context) GetHTTPReqEntry() (reqData []byte, err error) {
	switch c.Request().Method {
	case http.MethodGet:
		qryParams := c.QueryParams()
		if len(qryParams) == 0 {
			return nil, nil
		}

		// QueryParam -> map[string]any
		m := make(map[string]any, len(qryParams))
		for k, _ := range qryParams {
			m[k] = c.QueryParam(k)
		}

		urlData, err := json.Marshal(m)
		if err != nil {
			return nil, errors.Wrap(err, "json marshal got err")
		}

		return urlData, nil
	case http.MethodPost:
		body, err := io.ReadAll(c.Request().Body)
		if err != nil {
			return nil, errors.Wrap(err, "read HTTP request body got err")
		}

		return body, nil
	}

	return nil, errors.Errorf("invalid http method: %s", c.Request().Method)
}

// NewContext 初始化一个HTTPd UtilContext
func NewContext(ctx echo.Context) *Context {
	// 携带元素
	meta := map[string]interface{}{
		TraceID: uuid.New().String(),
	}

	return &Context{
		Context: ctx,
		meta:    meta,
	}
}
