package httpd

import (
	"bytes"
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
	"github.com/lupguo/wisdom-httpd/config"
)

// WisdomTemplate 实现 echo.Renderer 接口
type WisdomTemplate struct {
	template *template.Template
}

// InitParseWisdomTemplate 初始化Wisdom模板
func InitParseWisdomTemplate() *WisdomTemplate {
	// 创建模板对象并注册自定义模板函数
	tpl := template.New("index.tmpl").Funcs(template.FuncMap{
		"include": func(filename string, data interface{}) (template.HTML, error) {
			tmpl, err := template.ParseFiles(config.GetViewRealPath(filename))
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
	})

	// 解析&渲染模板文件
	views := []string{
		"index.tmpl",
		"wisdom.tmpl",
		"main/*.tmpl",
		"partial/*.tmpl",
	}
	viewRealPaths := config.GetViewRealPathList(views...)
	tpl = template.Must(tpl.ParseFiles(viewRealPaths...))
	return &WisdomTemplate{
		template: tpl,
	}
}

// Render Wisdom渲染模板
func (t *WisdomTemplate) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.template.ExecuteTemplate(w, name, data)
}
