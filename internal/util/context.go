package util

import (
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

const (
	TraceID = `trace_id`
	BizData = `biz_data`
)

// Context Wisdom HTTP的上下文
type Context struct {
	echo.Context
	meta map[string]string
}

func (c *Context) Deadline() (deadline time.Time, ok bool) {
	return
}

func (c *Context) Done() <-chan struct{} {
	return nil
}

func (c *Context) Err() error {
	return nil
}

func (c *Context) Value(key any) any {
	return c.meta[key.(string)]
}

// NewContext 初始化一个HTTPd Context
func NewContext(ctx echo.Context) *Context {
	// 携带元素
	meta := make(map[string]string)
	meta[TraceID] = uuid.New().String()

	return &Context{
		Context: ctx,
		meta:    meta,
	}
}
