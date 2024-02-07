package handler

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/lupguo/wisdom-httpd/config"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

// WisdomHandler 名言处理
func WisdomHandler(c echo.Context) error {
	// 随机一条wisdom
	wisdom, err := generateOneRandWisdom()
	if err != nil {
		return c.String(http.StatusOK, err.Error())
	}

	// json 响应
	reqType := c.QueryParam("type")
	if reqType == "json" {
		data, err := json.Marshal(wisdom)
		if err != nil {
			return c.String(http.StatusOK, err.Error())
		}

		return c.JSONBlob(http.StatusOK, data)
	}

	// web html 响应
	return c.Render(http.StatusOK, "wisdom.tmpl", wisdom)
}

// ParseJsonWisdom 从json解析wisdom
func ParseJsonWisdom(filename string) ([]*Wisdom, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, errors.Wrap(err, "read file `wisdomHandler.json` got err")
	}
	var ws []string
	err = json.Unmarshal(data, &ws)
	if err != nil {
		return nil, errors.Wrap(err, "unmarshal wisdomHandler json data got err")
	}

	var wisdoms []*Wisdom
	for _, w := range ws {
		wisdoms = append(wisdoms, &Wisdom{
			Sentence: w,
		})
	}

	return wisdoms, nil
}

// Wisdom 名言警句
type Wisdom struct {
	Sentence string `json:"sentence,omitempty"` // 句子
	WType    int    `json:"w_type,omitempty"`   // 类型
	Desc     string `json:"desc,omitempty"`     // 描述
	Img      string `json:"img,omitempty"`      // 图片
}

func generateOneRandWisdom() (*Wisdom, error) {
	// 解析wisdoms
	wisdoms, err := ParseJsonWisdom(config.GetWisdomFilename())
	if err != nil {
		log.Debugf("generate wisdom got err: %v", err)
		return nil, errors.Wrap(err, "wisdom handler got err")
	}
	if len(wisdoms) < 0 {
		return nil, errors.Wrap(err, "empty wisdom")
	}

	// 随机一条
	randIdx := rand.Int31n(int32(len(wisdoms)))
	randWisdom := wisdoms[randIdx]
	logrus.Debugf("rand wisdom: %v", randWisdom)

	return randWisdom, nil
}
