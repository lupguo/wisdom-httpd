package api

import (
	"context"
	"encoding/json"

	"github.com/lupguo/go-shim/shim"
	"github.com/lupguo/wisdom-httpd/app/application"
	"github.com/lupguo/wisdom-httpd/app/domain/entity"
	"github.com/lupguo/wisdom-httpd/app/domain/entity/crp"
	"github.com/lupguo/wisdom-httpd/internal/log"
	"github.com/pkg/errors"
)

// WisdomHandler 接口初始化
type WisdomHandler struct {
	wisdomApp application.IAppWisdom
	appTools  *application.ToolsApp
}

// NewWisdomHandlerImpl 初始化wisdom实现
func NewWisdomHandlerImpl(wisdomApp application.IAppWisdom, toolApp *application.ToolsApp) *WisdomHandler {
	return &WisdomHandler{
		wisdomApp: wisdomApp,
		appTools:  toolApp,
	}
}

// Index 首页渲染
func (h *WisdomHandler) Index(ctx context.Context, req []byte) (rsp any, err error) {
	wisdom, err := h.wisdomApp.GetRandOneWisdomFromJsonFile(nil, false)
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

// GetWisdom 通过序号获取wisdom信息
func (h *WisdomHandler) GetWisdom(ctx context.Context, reqData []byte) (rsp any, err error) {
	req := &crp.GetOneWisdomReq{}
	if err = json.Unmarshal(reqData, &req); err != nil {
		return nil, errors.Wrap(err, "fn[GetWisdom] handle unmarshal req data got err")
	}

	qryCond := &entity.WisdomQryCond{
		WisdomNos: req.GetNos(),
		Keywords:  req.Keywords,
		Speaker:   req.Speaker,
		Random:    req.IsRandom(),
	}

	wisdom, err := h.wisdomApp.GetWisdomByCond(ctx, qryCond)
	if err != nil {
		return nil, errors.Wrap(err, "fn[GetWisdom] got err")
	}

	return wisdom, nil
}

// SaveWisdom 保存
func (h *WisdomHandler) SaveWisdom(ctx context.Context, reqData []byte) (rsp any, err error) {
	req := &crp.SaveWisdomReq{}
	if err = json.Unmarshal(reqData, &req); err != nil {
		return nil, errors.Wrap(err, "fn[SaveWisdom] handle unmarshal req data got err")
	}

	// 创建wisdom
	wisdom, err := entity.NewWisdom(req)
	if err != nil {
		return nil, errors.Wrap(err, "fn[SaveWisdom] new wisdom got got err")
	}
	err = h.wisdomApp.SaveOneWisdom(ctx, wisdom)
	if err != nil {
		return nil, errors.Wrap(err, "fn[SaveWisdom] handle save wisdom got err")
	}

	log.InfoContextf(ctx, "save wisdom=>%s", shim.ToJsonString(wisdom))

	return &crp.SaveWisdomRsp{
		Code:     wisdom.WisdomNo,
		Sentence: wisdom.Sentence,
	}, nil
}
