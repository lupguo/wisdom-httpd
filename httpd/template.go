package httpd

import (
	"bytes"
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
	"github.com/lupguo/wisdom-httpd/app/infra/conf"
	"github.com/pkg/errors"
)

// HTMLRenderer 实现 echo.Renderer 接口
type HTMLRenderer struct {
	tpl *template.Template
}

// NewHTMLRenderer 初始化Wisdom模板渲染器
// 1. 初始化模版
// 2. 注册模版函数
// 3. 编译所有模版
func NewHTMLRenderer() (*HTMLRenderer, error) {
	tpl := template.New("wisdom")

	// 注册自定义模板函数
	tpl.Funcs(
		template.FuncMap{
			"include": func(filename string, data interface{}) (template.HTML, error) {
				tmpl, err := template.ParseFiles(conf.GetViewPathList(filename)...)
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

	// 渲染ALL模板文件（新增了模版需要重新编译）
	views := conf.GetViewParseFiles()
	for t, view := range views {
		viewPaths := conf.GetViewPathList(view...)
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

	return &HTMLRenderer{
		tpl: tpl,
	}, nil
}

// Render Wisdom渲染模板
func (t *HTMLRenderer) Render(w io.Writer, tplName string, data interface{}, _ echo.Context) error {
	if t.tpl.Lookup(tplName) == nil {
		return errors.Errorf("HTMLRenderer render got err, template[%s] not found", tplName)
	}

	return t.tpl.ExecuteTemplate(w, tplName, data)
}

// DefinedTemplates 获取已定义的模版信息
func (t *HTMLRenderer) DefinedTemplates() string {
	return t.tpl.DefinedTemplates()
}
