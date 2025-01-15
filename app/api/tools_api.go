package api

import (
	"context"

	"github.com/lupguo/wisdom-httpd/app/domain/entity/crp"
	"github.com/lupguo/wisdom-httpd/app/infra/conf"
	"github.com/pkg/errors"
)

// ToolRefreshToDB 工具保存到DB
func (h *WisdomHandler) ToolRefreshToDB(ctx context.Context, reqData []byte) (rsp any, err error) {
	if conf.GetRefreshToDBFlag() == false {
		return &crp.RspBody{Data: "refresh to db is disabled"}, nil
	}

	err = h.appTools.SaveWisdomToDB(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to save wisdom to db")
	}

	return nil, err
}
