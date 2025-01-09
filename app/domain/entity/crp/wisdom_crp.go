package crp

import (
	"github.com/pkg/errors"
)

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

// Validate 请求检测
// todo 可以转到PB到validate能力
func (s *SaveWisdomReq) Validate() error {
	if s == nil {
		return errors.New("SaveWisdomReq is nil")
	}
	if s.Sentence == "" {
		return errors.New("Sentence is empty")
	}

	return nil
}

type SaveWisdomRsp struct{}
