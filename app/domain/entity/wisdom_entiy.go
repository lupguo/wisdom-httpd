package entity

import (
	"crypto/rand"
	"fmt"
	"math/big"

	"github.com/lupguo/wisdom-httpd/app/domain/entity/crp"
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
func NewWisdom(p *crp.SaveWisdomReq) *Wisdom {
	return &Wisdom{
		WisdomNo: GenerateRandomHex(),
		Sentence: p.Sentence,
		Speaker:  p.Speaker,
		ReferURL: p.ReferURL,
		Image:    "", // todo AI
	}
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
}

// WisdomUpdEntry 更新内容
type WisdomUpdEntry struct {
	Image string
}

// GenerateRandomHex 生成一个随机数，范围在0到65535之间（即16进制的FFFF）
func GenerateRandomHex() string {
	n, _ := rand.Int(rand.Reader, big.NewInt(1e8))
	return fmt.Sprintf("0x%0X", n)
}

// WisdomList 配置列表
type WisdomList struct {
	Preview   []string `json:"preview,omitempty"`
	Sentences []string `json:"sentences,omitempty"`
}
