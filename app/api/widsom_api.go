package api

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
func WisdomHandler(c echo.Context) (err error) {
	// 预览参数
	isPreview := false
	preview := c.QueryParam("preview")
	if preview != "" {
		isPreview, err = strconv.ParseBool(preview)
		if err != nil {
			log.Errorf("wisdom handler get err, parse query param[preview] got err: %v", err)
		}
	}

	// 获取wisdom
	wisdom, err := GetAnRandomWisdom(isPreview)
	if err != nil {
		log.Errorf("wisdom handler get err, getAnRandomWisdom got err: %v", err)
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
	Preview   []string `json:"preview,omitempty"`
	Sentences []string `json:"sentences,omitempty"`
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

// GetAnRandomWisdom Randomly obtain and generate a famous aphorism
func GetAnRandomWisdom(isPreview bool) (*Wisdom, error) {
	// 解析wisdoms.json文件
	list, err := ParseJsonWisdom(config.GetWisdomFilePath())
	if err != nil {
		return nil, errors.Wrap(err, "wisdom handler got err")
	}

	// 从json文件获取指定的内容
	sentences := list.Sentences
	if isPreview == true {
		sentences = list.Preview
	}
	if len(sentences) <= 0 {
		return nil, errors.Errorf("get json content for preview[%v] is empty", isPreview)
	}

	// 获取所有的wisdom内容
	var wisdoms []*Wisdom
	for _, s := range sentences {
		wisdoms = append(wisdoms, &Wisdom{
			Sentence: s,
		})
	}

	// 随机生成一条wisdom内容
	randIdx := rand.Int31n(int32(len(wisdoms)))
	randWisdom := wisdoms[randIdx]
	log.Debugf("rand wisdom: %v", randWisdom)

	return randWisdom, nil
}
