package handler

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"os"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/lupguo/wisdom-httpd/config"
	"github.com/pkg/errors"
)

// WisdomHandler 名言处理
func WisdomHandler(c echo.Context) error {
	// 随机一条wisdom
	preview, err := strconv.ParseBool(c.QueryParam("preview"))
	if err != nil {
		log.Errorf("wisdom handler get err, parse query param[preview] got err: %v", err)
		return err
	}

	wisdom, err := generateOneRandWisdom(preview)
	if err != nil {
		log.Errorf("wisdom handler get err, generateOneRandWisdom got err: %v", err)
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

// WisdomList 配置列表
type WisdomList struct {
	Preview []string `json:"preview,omitempty"`
	Show    []string `json:"show,omitempty"`
}

// ParseJsonWisdom 从json解析wisdom
func ParseJsonWisdom(filename string) (*WisdomList, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, errors.Wrap(err, "read file `wisdomHandler.json` got err")
	}

	var ws WisdomList
	err = json.Unmarshal(data, &ws)
	if err != nil {
		return nil, errors.Wrap(err, "unmarshal wisdomHandler json data got err")
	}
	return &ws, nil
}

// Wisdom 名言警句
type Wisdom struct {
	Sentence string `json:"sentence,omitempty"` // 句子
	WType    int    `json:"w_type,omitempty"`   // 类型
	Desc     string `json:"desc,omitempty"`     // 描述
	Img      string `json:"img,omitempty"`      // 图片
}

// 随机生成一条名言警句
func generateOneRandWisdom(preview bool) (*Wisdom, error) {
	// 解析wisdoms.json文件
	jsonCont, err := ParseJsonWisdom(config.GetWisdomFilename())
	if err != nil {
		return nil, errors.Wrap(err, "wisdom handler got err")
	}

	// 从json文件获取指定的内容
	wisdomStrs := jsonCont.Show
	if preview == true {
		wisdomStrs = jsonCont.Preview
	}
	if len(wisdomStrs) < 0 {
		return nil, errors.Errorf("get json content for %v is empty", preview)
	}

	// 获取所有的wisdom内容
	var wisdoms []*Wisdom
	for _, w := range wisdomStrs {
		wisdoms = append(wisdoms, &Wisdom{
			Sentence: w,
		})
	}

	// 随机生成一条wisdom内容
	randIdx := rand.Int31n(int32(len(wisdomStrs)))
	randWisdom := wisdoms[randIdx]
	log.Debugf("rand wisdom: %v", randWisdom)

	return randWisdom, nil
}
