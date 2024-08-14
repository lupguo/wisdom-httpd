package api

import (
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/lupguo/go-shim/shim"
	"github.com/lupguo/wisdom-httpd/app/application"
	"github.com/lupguo/wisdom-httpd/app/domain/entity"
)

// WisdomHandler 名言处理
func (impl *SrvImpl) WisdomHandler(c echo.Context) (rsp *entity.WebPageData, err error) {
	// 预览参数
	preview := c.QueryParam("preview")
	isPreview, _ := strconv.ParseBool(preview)

	// 获取wisdom
	wisdom, err := application.GetRandomWisdom(isPreview)
	if err != nil {
		return nil, shim.LogAndWrapErr(err, "fn[WisdomHandler] get rand wisdom got an err")
	}

	return &entity.WebPageData{
		TemplateName: "wisdom.tmpl",
		PageData:     wisdom,
	}, nil
}
