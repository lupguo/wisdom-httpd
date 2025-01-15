package application

import (
	"context"

	"github.com/lupguo/wisdom-httpd/app/domain/service"
	"github.com/lupguo/wisdom-httpd/internal/log"
	"github.com/pkg/errors"
)

type ToolsApp struct {
	wsrv service.IServiceWisdom
}

func NewToolsApp(wsrv service.IServiceWisdom) *ToolsApp {
	return &ToolsApp{wsrv: wsrv}
}

// SaveWisdomToDB 保存wisdom到db
func (t *ToolsApp) SaveWisdomToDB(ctx context.Context) error {
	// wisdoms 从文件获取信息
	wisdoms, err := t.wsrv.GetWisdomsFromFiles(ctx)
	if err != nil {
		return errors.Wrap(err, "fn[GetWisdomsFromFiles] get Wisdoms failed")
	}

	// db
	err = t.wsrv.SaveWisdoms(ctx, wisdoms)
	if err != nil {
		return errors.Wrap(err, "fn[SaveWisdoms] failed to save wisdoms")
	}

	log.InfoContextf(ctx, "save done")
	return nil
}
