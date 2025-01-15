package entity

import (
	"crypto/sha256"
	"fmt"

	"github.com/lupguo/wisdom-httpd/app/domain/entity/crp"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

// Wisdom 名言实体
type Wisdom struct {
	gorm.Model `json:"-"`
	WisdomNo   string `json:"wisdom_no" gorm:"column:wisdom_no;unique;not null"` // 名言16进制编码（对外展示）
	Sentence   string `json:"sentence,omitempty" gorm:"column:sentence"`         // 名言句子
	Speaker    string `json:"speaker,omitempty" gorm:"column:speaker"`           // 名人
	ReferURL   string `json:"refer_url,omitempty" gorm:"column:refer_url"`       // 出处URL, URL Reference
	Image      string `json:"image,omitempty" gorm:"column:image"`               // 名言关联图片（由AI生成）
}

// NewWisdom 初始化一个wisdom
func NewWisdom(p *crp.SaveWisdomReq) (*Wisdom, error) {
	if err := p.Validate(); err != nil {
		return nil, errors.Wrap(err, "create wisdom validation fail")
	}

	return &Wisdom{
		WisdomNo: GenerateRandomHex(p.Sentence),
		Sentence: p.Sentence,
		Speaker:  p.Speaker,
		ReferURL: p.ReferURL,
		Image:    "", // todo AI
	}, nil
}

// TableName 名言表
func (w Wisdom) TableName() string {
	return `t_wisdoms`
}

// WisdomQryCond 查询检索条件
type WisdomQryCond struct {
	Ids       []uint32 // 通过Wisdom ID
	WisdomNos []string // 通过Wisdom WisdomNos
	Keywords  string   // 关键字(句子、描述关键字模糊匹配)
	Speaker   string
	Random    bool // 是否随机数据
}

// WisdomUpdEntry 更新内容
type WisdomUpdEntry struct {
	Image string
}

// GenerateRandomHex 返回字符串s压缩对应的16进制字字符串
// 1. 对整个串生成Hash，Hash串转16进制
// 2. 取前N个Hash串内容
func GenerateRandomHex(s string) string {
	// 生成哈希值
	hash := sha256.New()
	hash.Write([]byte(s))
	hashBytes := hash.Sum(nil)

	// 转换为16进制字符串并截取前4个字符
	return fmt.Sprintf("0x%0X", hashBytes[:3])
}

// WisdomList 配置列表
type WisdomList struct {
	Preview   []string `json:"preview,omitempty"`
	Sentences []string `json:"sentences,omitempty"`
}
