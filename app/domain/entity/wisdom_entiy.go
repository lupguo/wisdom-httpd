package entity

import (
	"time"
)

// Wisdom 治理名言实体
type Wisdom struct {
	Id          uint32    `json:"id,omitempty"`          // 主键
	WisdomCode  string    `json:"wisdom_code"`           // 名言16进制编码（对外展示）
	WisdomType  int       `json:"wisdom_type,omitempty"` // 名言类型
	Sentence    string    `json:"sentence,omitempty"`    // 名言句子
	Explanation string    `json:"explanation"`           // 名言解读
	FromRegion  int       `json:"from_region"`           // 名言出自国内还是国外（0：国内 1：国外）
	GenSource   int       `json:"gen_source"`            // 名言是通过什么渠道生成的（0：AI生成 1：手动录入）
	WisdomImg   string    `json:"wisdom_img,omitempty"`  // 名言关联图片（由AI生成）
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `json:"delete_at"`
}

// WisdomType 至理名言所属类型
type WisdomType int

const (
	WTypeMotivationAndStruggle    WisdomType = iota // 激励和奋斗
	WTypeSuccessAndGrowth                           // 成功和成长
	WTypeHappinessAndSatisfaction                   // 快乐和满足
	WTypeLoveAndRelationships                       // 爱和人际关系
	WTypeWisdomAndPhilosophy                        // 智慧和哲学
)

// GenSource 生成来源
type GenSource int

const (
	GSourceOpenAI = iota // AI生成
	GSourceManual        // 手动录入
)

// FromRegion 名言出处
type FromRegion int

const (
	FromRegionChina  = iota // 来至中国
	FromRegionAbroad        // 来至国外
)

// WisdomQryCond 查询检索条件
type WisdomQryCond struct {
	Id         []uint32   // 通过Wisdom ID
	Code       []string   // 通过Wisdom Code
	Keywords   string     // 关键字(句子、描述关键字模糊匹配)
	WisdomType WisdomType // 类型
	FromRegion FromRegion // 名言出处
	GenSource  GenSource  // 生成来源(OpenAI、人工)
	Limit      uint32     // 数量
}
