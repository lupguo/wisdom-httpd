package httpd

import (
	"bytes"
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
	"github.com/lupguo/wisdom-httpd/config"
	"github.com/pkg/errors"
)

// WisdomTemplate 实现 echo.Renderer 接口
type WisdomTemplate struct {
	template *template.Template
}

// InitParseWisdomTemplate 初始化Wisdom模板
func InitParseWisdomTemplate() (*WisdomTemplate, error) {
	// 创建模板对象并注册自定义模板函数
	tpl := template.New("wisdom").Funcs(
		template.FuncMap{
			"include": func(filename string, data interface{}) (template.HTML, error) {
				tmpl, err := template.ParseFiles(config.GetViewPathList(filename)...)
				if err != nil {
					return "", err
				}
				var result bytes.Buffer
				err = tmpl.ExecuteTemplate(&result, filename, data)
				if err != nil {
					return "", err
				}
				return template.HTML(result.String()), nil
			},
		},
	)

	// 解析&渲染全部模板文件
	views := config.GetViewParseFiles()
	for t, view := range views {
		viewPaths := config.GetViewPathList(view...)
		switch t {
		case "files":
			if _, err := tpl.ParseFiles(viewPaths...); err != nil {
				return nil, errors.Wrapf(err, "tpl.ParseFiles[%v] got err", viewPaths)
			}
		case "glob":
			for _, globViewPath := range viewPaths {
				if _, err := tpl.ParseGlob(globViewPath); err != nil {
					return nil, errors.Wrapf(err, "tpl.ParseGlob[%v] got err", globViewPath)
				}
			}
		}
	}

	return &WisdomTemplate{
		template: tpl,
	}, nil
}

// Render Wisdom渲染模板
func (t *WisdomTemplate) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.template.ExecuteTemplate(w, name, data)
}
