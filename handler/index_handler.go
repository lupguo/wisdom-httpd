package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type PageData struct {
	Error  error
	User   *User
	Wisdom string
	Data   interface{}
}

type IndexData struct {
	Content string
}

type User struct {
	Name string
}

// IndexHandler 首页渲染
func IndexHandler(c echo.Context) error {
	wisdom, err := generateOneRandWisdom()
	if err != nil {
		return c.Render(http.StatusInternalServerError, "error.tmpl", err.Error())
	}

	return c.Render(http.StatusOK, "index.tmpl", &PageData{
		Error:  nil,
		User:   &User{Name: "TerryRod"},
		Wisdom: wisdom.Sentence,
		Data: IndexData{
			Content: "附带Body内容",
		},
	})
}

func ErrorHandler(c echo.Context) error {
	return c.Render(http.StatusOK, "error", "mock error!")
}
