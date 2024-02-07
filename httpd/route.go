package httpd

import (
	"github.com/labstack/echo/v4"
	"github.com/lupguo/wisdom-httpd/config"
	"github.com/lupguo/wisdom-httpd/handler"
)

func routerInit(e *echo.Echo) {
	e.Static("/", config.GetAssetPath())

	// handler
	e.GET("/", handler.IndexHandler)
	e.GET("/error", handler.ErrorHandler)

	// handler - wisdom
	e.GET("/wisdom", handler.WisdomHandler)

	// handler - code
	e.GET("/code", handler.CodeHandler)

	// handler - 文件下载
	e.GET("/files/upload", handler.UploadHandler)
	e.GET("/files/download", handler.DownloadHandler)
}
