package httpd

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/lupguo/go-shim/shim"
	"github.com/lupguo/wisdom-httpd/app/domain/entity"
	"github.com/pkg/errors"
)

// Response 统一错误响应格式
type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func JSONResponseMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// 调用下一个处理器
		err := next(c)
		if err != nil {
			// 处理错误响应
			return c.JSON(http.StatusInternalServerError, Response{
				Status:  "error",
				Message: err.Error(),
			})
		}

		// 获取响应状态码
		status := c.Response().Status

		// 处理成功响应
		return c.JSON(status, Response{
			Status:  "success",
			Message: "Request processed successfully",
			Data:    c.Get("data"), // 假设您在处理器中设置了数据
		})
	}
}

func HTMLResponseMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// 调用下一个处理器
		err := next(c)
		if err != nil {
			return c.Render(http.StatusInternalServerError, "error.tmpl", err.Error())
		}

		// 获取响应状态码
		data := c.Get("data")
		pageData, ok := data.(*entity.WebPageData)
		if !ok {
			err = errors.Errorf("fn[HTMLResponseMiddleware] web page data[%s] assert fail", shim.ToJsonString(data, false))
			return c.Render(http.StatusInternalServerError, "error.tmpl", err.Error())
		}

		return c.Render(http.StatusOK, pageData.GetTemplateName(), pageData.GetPageData())
	}
}
