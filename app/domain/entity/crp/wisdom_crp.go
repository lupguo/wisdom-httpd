package crp

// GetOneWisdomReq 请求参数
type GetOneWisdomReq struct {
	Preview bool `json:"preview"`
}

// GetOneWisdomRsp 请求响应
type GetOneWisdomRsp struct {
	Sentence string `json:"sentence"` // 句子
}

// SaveWisdomReq 请求
type SaveWisdomReq struct {
	Sentence string `json:"sentence,omitempty"`  // 名言句子
	Speaker  string `json:"speaker,omitempty"`   // 名人
	ReferURL string `json:"refer_url,omitempty"` // 出处URL, URL Reference
}

type SaveWisdomRsp struct{}
