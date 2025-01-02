package entity

// GetOneWisdomReq 请求参数
type GetOneWisdomReq struct {
	Preview bool `json:"preview"`
}

// GetOneWisdomRsp 请求响应
type GetOneWisdomRsp struct {
	Sentence string `json:"sentence"` // 句子
}

// WisdomList 配置列表
type WisdomList struct {
	Preview   []string `json:"preview,omitempty"`
	Sentences []string `json:"sentences,omitempty"`
}

// Wisdom 名言警句
type Wisdom struct {
	Sentence string `json:"sentence,omitempty"` // 句子
	WType    int    `json:"w_type,omitempty"`   // 类型
	Desc     string `json:"desc,omitempty"`     // 描述
	Img      string `json:"img,omitempty"`      // 图片
}
