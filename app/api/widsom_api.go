package api

import (
	"context"
	"encoding/json"

	"github.com/lupguo/wisdom-httpd/app/application"
	"github.com/lupguo/wisdom-httpd/app/domain/entity"
	"github.com/lupguo/wisdom-httpd/app/domain/entity/crp"
	"github.com/lupguo/wisdom-httpd/internal/log"
	"github.com/lupguo/wisdom-httpd/internal/util"
	"github.com/pkg/errors"
)

// WisdomHandler 接口初始化
type WisdomHandler struct {
	app application.IAppWisdom
}

// NewWisdomHandlerImpl 初始化wisdom实现
func NewWisdomHandlerImpl(app application.IAppWisdom) *WisdomHandler {
	return &WisdomHandler{app: app}
}

// Index 首页渲染
func (h *WisdomHandler) Index(ctx context.Context, req []byte) (rsp any, err error) {
	wisdom, err := h.app.GetRandOneWisdom(nil, false)
	if err != nil {
		return nil, err
	}

	rsp = &crp.PageDataIndexRsp{
		User:    &crp.User{Name: "Rod"},
		Wisdom:  wisdom.Sentence,
		Content: "wisdom page index content",
	}

	return rsp, nil
}

// GetOneWisdom 获取一条名言金句
func (h *WisdomHandler) GetOneWisdom(ctx context.Context, _ []byte) (rsp any, err error) {
	utilCtx := ctx.(*util.Context)
	log.InfoContextf(ctx, "req => %v", utilCtx.Get("preview"))

	// 获取wisdom
	wisdom, err := h.app.GetRandOneWisdom(ctx, false)
	if err != nil {
		return nil, log.WrapAndLogErrorf(ctx, err, "fn[GetOneWisdom] get rand wisdom got an err")
	}

	rsp = &entity.Wisdom{
		WisdomNo: "0x3FBA",
		Sentence: wisdom.Sentence,
		Speaker:  "鲁迅",
		ReferURL: "https://tkstorm.com",
		Image:    "https://localhost:1666/imgs/code.png",
	}

	return rsp, nil
}

// SaveWisdom 保存
func (h *WisdomHandler) SaveWisdom(ctx context.Context, reqData []byte) (rsp any, err error) {
	req := &crp.SaveWisdomReq{}
	if err = json.Unmarshal(reqData, &req); err != nil {
		return nil, errors.Wrap(err, "fn[SaveWisdom] handle unmarshal req data got err")
	}

	// 基本检测
	if err = req.Validate(); err != nil {
		return nil, errors.Wrap(err, "fn[SaveWisdom] handle validate req data got err")
	}

	err = h.app.SaveOneWisdom(ctx, entity.NewWisdom(req))
	if err != nil {
		return nil, errors.Wrap(err, "fn[SaveWisdom] handle save wisdom got err")
	}

	return &crp.SaveWisdomRsp{}, nil
}
