package service

import (
	"context"

	"github.com/lupguo/wisdom-httpd/app/domain/entity"
	"github.com/lupguo/wisdom-httpd/app/domain/repository"
	"github.com/pkg/errors"
)

// IServiceWisdom 领域服务要提供哪些能
type IServiceWisdom interface {
	// GetOneRandomWisdom 获取随机一条Wisdom
	// GetOneRandomWisdom(ctx context.Context) (*entity.Wisdom, error)

	// GetWisdoms 按条件获取一批Wisdoms
	GetWisdoms(ctx context.Context, qryCond *entity.WisdomQryCond, pageLimit *entity.PageLimit) ([]*entity.Wisdom, error)

	// SaveWisdoms 生成Wisdom信息，存储到DB中
	SaveWisdoms(ctx context.Context, wisdoms []*entity.Wisdom) error
}

// WisdomService Wisdom服务依赖的基础设施仓储接口
type WisdomService struct {
	dbsInfra repository.IReposWisdomDB
}

// GetWisdoms 按条件获取一批Wisdoms
func (w *WisdomService) GetWisdoms(ctx context.Context, qryCond *entity.WisdomQryCond, pageLimit *entity.PageLimit) ([]*entity.Wisdom, error) {
	// 验证输入参数
	if qryCond == nil || pageLimit == nil {
		return nil, errors.New("invalid query condition or page limit")
	}

	// 调用数据库查询
	wisdoms, err := w.dbsInfra.SelectWisdom(ctx, qryCond, pageLimit)
	if err != nil {
		return nil, err
	}

	return wisdoms, nil
}

// SaveWisdoms 生成Wisdom信息，存储到DB中
func (w *WisdomService) SaveWisdoms(ctx context.Context, wisdoms []*entity.Wisdom) error {
	// 验证输入参数
	if wisdoms == nil || len(wisdoms) == 0 {
		return errors.New("no wisdoms to save")
	}

	// 调用数据库插入
	if err := w.dbsInfra.InsertWisdom(ctx, wisdoms); err != nil {
		return err
	}

	return nil
}
