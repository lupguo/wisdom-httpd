package httpd

import (
	"github.com/labstack/echo/v4"
	"github.com/lupguo/wisdom-httpd/app/api"
	"github.com/lupguo/wisdom-httpd/config"
)

// 路由配置
func routerInit(e *echo.Echo) {
	// 静态资源配置
	e.Static("/", config.DistPath())

	// 首页处理
	e.GET("/", api.IndexHandler)
	e.GET("/error", api.ErrorHandler)

	// wisdom接口处理
	e.GET("/wisdom", api.WisdomHandler)

	// 调试statusCode处理
	e.GET("/code", api.CodeHandler)

	// handler - 文件下载mock
	e.GET("/files/upload", api.UploadHandler)
	e.GET("/files/download", api.DownloadHandler)
}
