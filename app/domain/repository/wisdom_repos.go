package repository

import (
	"github.com/lupguo/wisdom-httpd/app/domain/entity"
)

// WisdomFileRepos 至理名言仓储接口
type WisdomFileRepos interface {
	// ReadWisdom 读取wisdom内容
	ReadWisdom(filename string) ([]*entity.Wisdom, error)
}

// WisdomDBRepos DB仓储接口要实现的能力
type WisdomDBRepos interface {
	// InsertWisdom 批量插入到DB表中
	InsertWisdom(list []*entity.Wisdom) error

	// SelectWisdom 查询指定ID的名言信息
	SelectWisdom(qryCond *entity.Wisdom) (*entity.Wisdom, error)

	// UpdateWisdom 更新
	UpdateWisdom(updEnt *entity.Wisdom, qryCond *entity.Wisdom) error

	// DeleteWisdom 删除
	DeleteWisdom(qryCond *entity.Wisdom) error
}

// WisdomCacheRepos 缓存仓储接口
type WisdomCacheRepos interface {
	// ReadWisdom 读取
	ReadWisdom(key string) (*entity.Wisdom, error)

	// WriteWisdom 写入
	WriteWisdom(key string, wisdom *entity.Wisdom) error

	// DelWisdom 删除
	DelWisdom(key string) error
}

// WisdomOpenAIRepos 关于OpenAI仓储层要实现能力
type WisdomOpenAIRepos interface {
	// GenerateAIWisdom 通过AI生成wisdom信息
	GenerateAIWisdom() ([]*entity.Wisdom, error)
}
