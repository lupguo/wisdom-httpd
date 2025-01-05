package api

import (
	"github.com/lupguo/go-shim/shim"
	"github.com/lupguo/wisdom-httpd/app/application"
	"github.com/lupguo/wisdom-httpd/app/domain/entity"
	"github.com/lupguo/wisdom-httpd/internal/log"
	"github.com/lupguo/wisdom-httpd/internal/util"
)

// WisdomHandler 接口初始化
type WisdomHandler struct {
	app application.WisdomAppInf
}

// NewWisdomImpl 初始化wisdom实现
func NewWisdomImpl(app application.WisdomAppInf) *WisdomHandler {
	return &WisdomHandler{app: app}
}

// Index 首页渲染
func (h *WisdomHandler) Index(ctx *util.Context, req any) (rsp any, err error) {
	wisdom, err := h.app.GetRandOneWisdom(nil, false)
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

// GetOneWisdom 名言处理
func (h *WisdomHandler) GetOneWisdom(ctx *util.Context, _ any) (rsp any, err error) {
	req := &entity.GetOneWisdomReq{}
	if err = ctx.Bind(&req); err != nil {
		return nil, err
	}

	// 获取wisdom
	log.Infof("req => %s", shim.ToJsonString(req))
	wisdom, err := h.app.GetRandOneWisdom(nil, req.Preview)
	if err != nil {
		return nil, log.WrapErrorContextf(ctx, err, "fn[GetOneWisdom] get rand wisdom got an err")
	}
	rsp = &entity.GetOneWisdomRsp{
		Sentence: wisdom.Sentence,
	}

	log.Infof("rsp <= %s", shim.ToJsonString(rsp))
	return rsp, nil
}

// SaveWisdom 保存
func (h *WisdomHandler) SaveWisdom(ctx *util.Context, _ any) (rsp any, err error) {

	return nil, nil
}
