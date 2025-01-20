package service

import (
	"context"

	"github.com/lupguo/wisdom-httpd/app/domain/entity"
	"github.com/lupguo/wisdom-httpd/app/domain/entity/crp"
	"github.com/lupguo/wisdom-httpd/app/domain/repos"
	"github.com/lupguo/wisdom-httpd/app/infra/conf"
	"github.com/lupguo/wisdom-httpd/app/infra/files"
	"github.com/pkg/errors"
)

// IServiceWisdom 领域服务要提供哪些能
type IServiceWisdom interface {
	// GetWisdoms 按条件获取一批Wisdoms
	GetWisdoms(ctx context.Context, qryCond *entity.WisdomQryCond, pageLimit *entity.PageLimit) ([]*entity.Wisdom, error)

	// SaveWisdoms 生成Wisdom信息，存储到DB中
	SaveWisdoms(ctx context.Context, wisdoms []*entity.Wisdom) error

	// GetWisdomsFromFiles 从Files解析Json
	GetWisdomsFromFiles(ctx context.Context) ([]*entity.Wisdom, error)
}

// WisdomService Wisdom服务依赖的基础设施仓储接口
type WisdomService struct {
	dbsInfra repos.IReposWisdomDB
}

// GetWisdomsFromFiles 从文件解析得到Wisdom服务
func (w *WisdomService) GetWisdomsFromFiles(_ context.Context) ([]*entity.Wisdom, error) {
	// 解析wisdoms.json文件
	data, err := files.ParseJsonWisdom(conf.GetWisdomSentenceFilePath())
	if err != nil {
		return nil, errors.Wrap(err, "wisdom handler got err")
	}

	// 从json文件获取指定的内容
	sentences := data.Sentences
	if len(sentences) <= 0 {
		return nil, errors.New("empty sentences")
	}

	// 获取所有的wisdom内容
	var wisdoms []*entity.Wisdom
	for _, s := range sentences {
		wisdom, err := entity.NewWisdom(&crp.SaveWisdomReq{
			Sentence: s,
			SKey:     conf.GetSKey(),
		})
		if err != nil {
			return nil, err
		}
		wisdoms = append(wisdoms, wisdom)
	}

	return wisdoms, nil
}

// NewWisdomService 初始化WisdomService
func NewWisdomService(dbsInfra repos.IReposWisdomDB) *WisdomService {
	return &WisdomService{dbsInfra: dbsInfra}
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
