package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// 定义模板内容
const codeTpl = `
	<!DOCTYPE html>
	<html>
	<head>
		<title>User Profile</title>
	</head>
	<body>
		<code>{{.}}</code>
	</body>
	</html>
`

// CodeHandler 代码处理器
func CodeHandler(c echo.Context) error {
	// tpl, _ := template.New("code").Parse(codeTpl)
	return c.Render(http.StatusOK, "code", nil)
	//
	// // 执行模板
	// data := &bytes.Buffer{}
	// err := tpl.Execute(data, `Hello World`)
	// if err != nil {
	// 	return c.HTML(http.StatusOK, err.Error())
	// }
	//
	// log.Infof("data code: %v", data)
	// return c.HTMLBlob(http.StatusOK, data.Bytes())
}
