package application

import (
	"context"
	"math/rand"

	"github.com/labstack/gommon/log"
	"github.com/lupguo/wisdom-httpd/app/domain/entity"
	"github.com/lupguo/wisdom-httpd/app/domain/service"
	"github.com/lupguo/wisdom-httpd/app/infra/conf"
	"github.com/lupguo/wisdom-httpd/app/infra/files"
	"github.com/pkg/errors"
)

// IAppWisdom wisdom应用接口
type IAppWisdom interface {
	// GetRandOneWisdomFromJsonFile 随机获取一条至理名言，如果是预览，默认给ID=1的首条
	GetRandOneWisdomFromJsonFile(ctx context.Context, isPreview bool) (*entity.Wisdom, error)

	// GetWisdomByCond 按条件查询wisdom
	GetWisdomByCond(ctx context.Context, qryCond *entity.WisdomQryCond) (*entity.Wisdom, error)

	// SaveOneWisdom 保存一条Wisdom记录到DB
	SaveOneWisdom(ctx context.Context, wisdom *entity.Wisdom) error
}

// WisdomApp Wisdom应用
type WisdomApp struct {
	wsrv service.IServiceWisdom
}

// GetWisdomByCond 按条件查询wisdom
func (app *WisdomApp) GetWisdomByCond(ctx context.Context, qryCond *entity.WisdomQryCond) (*entity.Wisdom, error) {
	pageLimit := &entity.PageLimit{
		Page:     1,
		PageSize: 1,
	}
	wisdoms, err := app.wsrv.GetWisdoms(ctx, qryCond, pageLimit)
	if err != nil {
		return nil, err
	}
	if len(wisdoms) >= 1 {
		return wisdoms[0], nil
	}

	return nil, errors.New("no Wisdom found")
}

// NewWisdomApp 初始化Wisdom的APP
func NewWisdomApp(wisdomSvr service.IServiceWisdom) *WisdomApp {
	return &WisdomApp{wsrv: wisdomSvr}
}

// SaveOneWisdom 保存一条wisdom记录到DB中
func (app *WisdomApp) SaveOneWisdom(ctx context.Context, wisdom *entity.Wisdom) error {
	return app.wsrv.SaveWisdoms(ctx, []*entity.Wisdom{wisdom})
}

// GetRandOneWisdomFromJsonFile Randomly obtain and generate a famous aphorism
func (app *WisdomApp) GetRandOneWisdomFromJsonFile(ctx context.Context, isPreview bool) (*entity.Wisdom, error) {
	// 解析wisdoms.json文件
	list, err := files.ParseJsonWisdom(conf.GetWisdomSentenceFilePath())
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
