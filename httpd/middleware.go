package httpd

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitMiddlewareConfig(e *echo.Echo) {
	// CORS 中间件配置
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},                                       // 允许所有域名，或指定特定域名
		AllowMethods:     []string{http.MethodGet, http.MethodPost},           // 允许的 HTTP 方法
		AllowHeaders:     []string{echo.HeaderContentType, echo.HeaderAccept}, // 允许的请求头
		ExposeHeaders:    []string{"X-My-Custom-Header"},                      // 允许暴露的响应头
		AllowCredentials: true,                                                // 是否允许发送 cookie
	}))
}
