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
	templates *template.Template
}

// GetRenderTemplate 获取APP的渲染模板
func GetRenderTemplate() *WisdomTemplate {
	// 创建模板对象并注册自定义模板函数
	tpl := template.New("index.tmpl").Funcs(template.FuncMap{
		"include": func(filename string, data interface{}) (template.HTML, error) {
			tmpl, err := template.ParseFiles(config.GetSpecialViewPath(filename))
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
	// 解析模板文件
	tpl = template.Must(tpl.ParseFiles(config.GetSpecialViewPathList("index.tmpl", "wisdom.tmpl")...))
	tpl = template.Must(tpl.ParseGlob(config.GetSpecialViewPath("main/*.tmpl")))
	tpl = template.Must(tpl.ParseGlob(config.GetSpecialViewPath("partial/*.tmpl")))
	tplRender := &WisdomTemplate{
		templates: tpl,
	}
	return tplRender
}

func (t *WisdomTemplate) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}
