package repos

import (
	"context"

	"github.com/lupguo/wisdom-httpd/app/domain/entity"
)

// IReposWisdomJsonFile 至理名言仓储接口
type IReposWisdomJsonFile interface {
	// ReadWisdom 读取wisdom内容
	ReadWisdom(filename string) ([]*entity.Wisdom, error)
}

// IReposWisdomDB DB仓储接口要实现的能力
type IReposWisdomDB interface {
	// InsertWisdom 批量插入到DB表中
	InsertWisdom(ctx context.Context, wisdoms []*entity.Wisdom) error

	// SelectWisdom 查询指定ID的名言信息
	SelectWisdom(ctx context.Context, qryCond *entity.WisdomQryCond, pageLimit *entity.PageLimit) ([]*entity.Wisdom, error)

	// UpdateWisdom 更新
	UpdateWisdom(ctx context.Context, updEntry *entity.WisdomUpdEntry, qryCond *entity.WisdomQryCond) error

	// DeleteWisdom 删除
	DeleteWisdom(ctx context.Context, qryCond *entity.WisdomQryCond) error
}

// IReposWisdomCache 缓存仓储接口
type IReposWisdomCache interface {
	// ReadWisdom 读取
	ReadWisdom(key string) (*entity.Wisdom, error)

	// WriteWisdom 写入
	WriteWisdom(key string, wisdom *entity.Wisdom) error

	// DelWisdom 删除
	DelWisdom(key string) error
}

// IReposWisdomOpenAI 关于OpenAI仓储层要实现能力
type IReposWisdomOpenAI interface {
	// GenerateAIWisdom 通过AI生成wisdom信息
	GenerateAIWisdom() ([]*entity.Wisdom, error)
}
