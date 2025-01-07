package util

import (
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
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
	// TODO implement me
	panic("implement me")
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
