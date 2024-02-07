package application

import (
	"time"

	"github.com/lupguo/wisdom-httpd/app/domain/entity"
	"github.com/lupguo/wisdom-httpd/app/domain/repository"
	"github.com/lupguo/wisdom-httpd/app/domain/service"
)

// WisdomAppInf wisdom应用接口
type WisdomAppInf interface {

	// GetRandOneWisdom 随机获取一条至理名言
	GetRandOneWisdom() (*entity.Wisdom, error)

	// CronAIWisdomGenerate Open定时请求获取至理名言存储到DB中
	CronAIWisdomGenerate(intervalTime time.Duration) error
}

type WisdomApp struct {
	// 生成Ai Repost
	openAIInfra repository.WisdomOpenAIRepos

	// 保存&读取wisdom
	wisdomSrv service.WisdomServiceInf
}

func NewWisdomApp(openAIInfra repository.WisdomOpenAIRepos, wisdomSrv service.WisdomServiceInf) *WisdomApp {
	return &WisdomApp{openAIInfra: openAIInfra, wisdomSrv: wisdomSrv}
}

func (wisApp *WisdomApp) GetRandOneWisdom() (*entity.Wisdom, error) {
	// TODO implement me

	return &entity.Wisdom{}, nil
}

func (wisApp *WisdomApp) CronAIWisdomGenerate(intervalTime time.Duration) error {

	return nil
}
