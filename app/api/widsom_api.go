package api

import (
	"encoding/json"
	"math/rand"
	"os"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/lupguo/go-shim/shim"
	"github.com/lupguo/wisdom-httpd/config"
	"github.com/pkg/errors"
)

// WisdomHandler 名言处理
func WisdomHandler(c echo.Context) error {
	// 预览参数
	preview := c.QueryParam("preview")
	isPreview, _ := strconv.ParseBool(preview)

	// 获取wisdom
	wisdom, err := GetRandomWisdom(isPreview)
	if err != nil {
		return shim.LogAndWrapErr(err, "fn[WisdomHandler] get rand wisdom got an err")
	}

	// 数据设置
	c.Set("data", wisdom)

	return nil
}

// WisdomRandHandler 名言处理
func WisdomRandHandler(c echo.Context) (err error) {
	// 预览参数
	preview := c.QueryParam("preview")
	isPreview, _ := strconv.ParseBool(preview)

	// 获取wisdom
	wisdom, err := GetRandomWisdom(isPreview)
	if err != nil {
		return shim.LogAndWrapErr(err, "fn[WisdomRandHandler] get rand wisdom got an err")
	}

	// 数据设置
	c.Set("data", wisdom)

	return nil
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

// GetRandomWisdom Randomly obtain and generate a famous aphorism
func GetRandomWisdom(isPreview bool) (*Wisdom, error) {
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
