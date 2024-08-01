package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/lupguo/go-shim/shim"
	"github.com/lupguo/wisdom-httpd/app/domain/entity"
)

// IndexHandler 首页渲染
func IndexHandler(c echo.Context) (rsp *entity.WebPageData, err error) {
	wisdom, err := GetRandomWisdom(false)
	if err != nil {
		return nil, err
	}

	rsp = &entity.WebPageData{
		TemplateName: "index.tmpl",
		PageData: &entity.IndexPageData{
			User:    &entity.User{Name: "TerryRod"},
			Wisdom:  wisdom.Sentence,
			Content: "附带Body内容",
		},
	}

	log.Infof("wisdom rsp <= %s", shim.ToJsonString(rsp, false))
	return rsp, nil
}

func ErrorHandler(c echo.Context) error {
	return c.Render(http.StatusOK, "error", "mock error!")
}
