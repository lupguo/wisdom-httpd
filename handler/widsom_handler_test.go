package handler

import (
	"encoding/json"
	"testing"
)

func TestMarshalWisdom(t *testing.T) {
	w := &Wisdom{
		Sentence: "知行合一",
		WType:    0,
		Desc:     "强调理论与实践的统一，即将所学的知识付诸于实际行动中。",
		Img:      "",
	}
	s, _ := json.Marshal(w)
	t.Logf("%s", s)
}
