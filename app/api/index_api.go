package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/lupguo/go-shim/shim"
	"github.com/lupguo/wisdom-httpd/app/application"
	"github.com/lupguo/wisdom-httpd/app/domain/entity"
)

// SrvImpl 接口初始化
type SrvImpl struct {
	c   echo.Context
	app application.WisdomAppInf
}

func NewImplAPI(app application.WisdomAppInf) *SrvImpl {
	return &SrvImpl{app: app}
}

// IndexHandler 首页渲染
func (impl *SrvImpl) IndexHandler(c echo.Context) (rsp *entity.WebPageData, err error) {
	wisdom, err := impl.app.GetRandOneWisdom(false)
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
