package api

import (
	"github.com/labstack/echo/v4"
	"github.com/lupguo/wisdom-httpd/app/application"
	"github.com/lupguo/wisdom-httpd/app/domain/entity"
)

// SrvImpl 接口初始化
type SrvImpl struct {
	c   echo.Context
	app application.WisdomAppInf
}

// NewWisdomImpl 初始化wisdom实现
func NewWisdomImpl(app application.WisdomAppInf) *SrvImpl {
	return &SrvImpl{app: app}
}

// IndexHandler 首页渲染
func (impl *SrvImpl) IndexHandler(c echo.Context, req any) (rsp any, err error) {
	wisdom, err := impl.app.GetRandOneWisdom(false)
	if err != nil {
		return nil, err
	}

	rsp = &entity.IndexPageData{
		User:    &entity.User{Name: "TerryRod"},
		Wisdom:  wisdom.Sentence,
		Content: "wisdom page index content",
	}

	return rsp, nil
}
