package crp

import (
	"encoding/json"
	"testing"
)

func TestSaveWisdomReq_Marshal(t *testing.T) {
	v := &SaveWisdomReq{
		Sentence: "Example Sentence",
		Speaker:  "Rod",
		ReferURL: "https://x.com",
	}

	b, err := json.Marshal(v)
	if err != nil {
		return
	}
	t.Logf("%s", string(b))

}
