package api

import (
	"encoding/json"
	"testing"

	"github.com/lupguo/wisdom-httpd/app/domain/entity"
)

func TestMarshalWisdom(t *testing.T) {
	w := &entity.Wisdom{
		Sentence: "知行合一",
		WType:    0,
		Desc:     "强调理论与实践的统一，即将所学的知识付诸于实际行动中。",
		Img:      "",
	}
	s, _ := json.Marshal(w)
	t.Logf("%s", s)
}

func TestParseJsonWisdom(t *testing.T) {
	data := `{
	"preview": [],
    "show": [
        "不积跬步无以至千里，不积小流无以致千里"
	]}
`

	var ws entity.WisdomList
	err := json.Unmarshal([]byte(data), &ws)
	if err != nil {
		t.Errorf("can not unmrashal data: %v", err)
	}
	t.Logf("%v", data)
}
