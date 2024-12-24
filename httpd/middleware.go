package httpd

import (
	"time"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

// LogrusMiddleware 是一个自定义的日志中间件
func LogrusMiddleware(logger *logrus.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// 记录请求信息
			start := time.Now()
			method := c.Request().Method
			path := c.Request().URL.Path

			// 调用下一个中间件或处理程序
			err := next(c)

			// 记录响应信息
			duration := time.Since(start)
			status := c.Response().Status

			// 打印日志
			logger.WithFields(logrus.Fields{
				"method":   method,
				"path":     path,
				"status":   status,
				"duration": duration,
			}).Info("Request processed")

			return err
		}
	}
}

//
// // Response 统一错误响应格式
// type Response struct {
// 	Status  string `json:"status"`
// 	Message string `json:"message"`
// 	Data    any    `json:"data,omitempty"`
// }
//
// // JSONResponseMiddleware JSON渲染
// func JSONResponseMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		// 调用下一个处理器
// 		err := next(c)
// 		if err != nil {
// 			// 处理错误响应
// 			return c.JSON(http.StatusInternalServerError, Response{
// 				Status:  "error",
// 				Message: err.Error(),
// 			})
// 		}
//
// 		// 获取响应状态码
// 		status := c.Response().Status
//
// 		// 处理成功响应
// 		return c.JSON(status, Response{
// 			Status:  "success",
// 			Message: "Request processed succ",
// 			Data:    c.Get("data"), // 在处理器中设置的数据
// 		})
// 	}
// }
//
// // WebResponseMiddleware Web渲染
// func WebResponseMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		// 调用下一个处理器
// 		err := next(c)
// 		if err != nil {
// 			return c.Render(http.StatusInternalServerError, "error.tmpl", err.Error())
// 		}
//
// 		// 获取响应状态码
// 		data := c.Get("data")
// 		rsp, ok := data.(*entity.Response)
// 		if !ok {
// 			err = errors.Errorf("fn[WebResponseMiddleware] web page data[%s] assert fail", shim.ToJsonString(data, false))
// 			return c.Render(http.StatusInternalServerError, "error.tmpl", err.Error())
// 		}
//
// 		return c.Render(http.StatusOK, rsp.GetTemplateName(), rsp.GetPageData())
// 	}
// }
