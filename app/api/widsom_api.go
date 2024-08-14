package api

import (
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/lupguo/go-shim/shim"
	"github.com/lupguo/wisdom-httpd/app/application"
	"github.com/lupguo/wisdom-httpd/app/domain/entity"
)

type ImplAPI struct {
	app application.WisdomApp
}

func NewImplAPI(app application.WisdomApp) *ImplAPI {
	return &ImplAPI{app: app}
}

// WisdomHandler 名言处理
func WisdomHandler(c echo.Context) (rsp *entity.WebPageDataRsp, err error) {
	// 预览参数
	preview := c.QueryParam("preview")
	isPreview, _ := strconv.ParseBool(preview)

	// 获取wisdom
	wisdom, err := application.GetRandomWisdom(isPreview)
	if err != nil {
		return nil, shim.LogAndWrapErr(err, "fn[WisdomHandler] get rand wisdom got an err")
	}

	return &entity.WebPageDataRsp{
		TemplateName: "wisdom.tmpl",
		PageData:     wisdom,
	}, nil
}
