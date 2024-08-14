package application

import (
	"math/rand"

	"github.com/labstack/gommon/log"
	"github.com/lupguo/wisdom-httpd/app/domain/entity"
	"github.com/lupguo/wisdom-httpd/app/infra/config"
	"github.com/lupguo/wisdom-httpd/app/infra/files"
	"github.com/pkg/errors"
)

// WisdomAppInf wisdom应用接口
type WisdomAppInf interface {

	// GetRandOneWisdom 随机获取一条至理名言
	GetRandOneWisdom() (*entity.Wisdom, error)

	// CronAIWisdomGenerate Open定时请求获取至理名言存储到DB中
	// CronAIWisdomGenerate(intervalTime time.Duration) error
}

type WisdomApp struct {
}

// GetRandomWisdom Randomly obtain and generate a famous aphorism
func GetRandomWisdom(isPreview bool) (*entity.Wisdom, error) {
	// 解析wisdoms.json文件
	list, err := files.ParseJsonWisdom(config.GetWisdomFilePath())
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
	var wisdoms []*entity.Wisdom
	for _, s := range sentences {
		wisdoms = append(wisdoms, &entity.Wisdom{
			Sentence: s,
		})
	}

	// 随机生成一条wisdom内容
	randIdx := rand.Int31n(int32(len(wisdoms)))
	randWisdom := wisdoms[randIdx]
	log.Debugf("rand wisdom: %v", randWisdom)

	return randWisdom, nil
}
