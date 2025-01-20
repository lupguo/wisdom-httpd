package crp

import (
	"github.com/lupguo/wisdom-httpd/app/infra/conf"
	"github.com/pkg/errors"
)

// GetOneWisdomReq 请求参数
type GetOneWisdomReq struct {
	No       string `json:"no,omitempty"`       // 格言的No
	Keywords string `json:"keywords,omitempty"` // 格言的部分文字
	Speaker  string `json:"speaker,omitempty"`  // 格言者
	Random   string `json:"random,omitempty"`   // 是否随机
}

// IsRandom 是否随机
func (r *GetOneWisdomReq) IsRandom() bool {
	return r.Random == "true" || r.Random == "1"
}

// GetNos 获取格言Nos
func (r *GetOneWisdomReq) GetNos() []string {
	var ss []string
	if r.No != "" {
		ss = append(ss, r.No)
	}
	return ss
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
	SKey     string `json:"skey,omitempty"`      // 密钥, yaml配置
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
	if s.SKey != conf.GetSKey() {
		return errors.New("Error Secret Key")
	}

	return nil
}

// SaveWisdomRsp 保存wisdom rsp响应格式
type SaveWisdomRsp struct {
	Code     string `json:"code"`
	Sentence string `json:"sentence,omitempty"` // 名言句子
}
