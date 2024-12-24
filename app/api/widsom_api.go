package api

import (
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/lupguo/go-shim/shim"
)

// WisdomHandler 名言处理
func (impl *SrvImpl) WisdomHandler(c echo.Context, req any) (rsp any, err error) {
	// 预览参数
	preview := c.QueryParam("preview")
	isPreview, _ := strconv.ParseBool(preview)

	// 获取wisdom
	wisdom, err := impl.app.GetRandOneWisdom(isPreview)
	if err != nil {
		return nil, shim.LogAndWrapErr(err, "fn[WisdomHandler] get rand wisdom got an err")
	}

	return wisdom, nil
}
