package httpd

import (
	"bytes"
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
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
			tmpl, err := template.ParseFiles(config.GetViewTmplPath(filename))
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
	indexViews := config.GetSpecialViewPathList("index.tmpl", "wisdom.tmpl")
	log.Infof("index view files: %v", indexViews)

	mainViews := config.GetViewTmplPath("main/*.tmpl")
	log.Infof("main view files: %v", mainViews)

	partialViews := config.GetViewTmplPath("partial/*.tmpl")
	log.Infof("partial view files: %v", partialViews)

	// 模版渲染
	tpl = template.Must(tpl.ParseFiles(indexViews...))
	tpl = template.Must(tpl.ParseGlob(mainViews))
	tpl = template.Must(tpl.ParseGlob(partialViews))
	tplRender := &WisdomTemplate{
		templates: tpl,
	}
	return tplRender
}

func (t *WisdomTemplate) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}
