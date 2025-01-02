package service

import (
	"github.com/lupguo/wisdom-httpd/app/domain/entity"
	"github.com/lupguo/wisdom-httpd/app/domain/repository"
)

// WisdomServiceInf 领域服务要提供哪些能
type WisdomServiceInf interface {

	// GetWisdom 从DB、File获取一条Wisdom内容
	GetWisdom(limit int) ([]*entity.Wisdom, error)

	// PostWisdom 生成Wisdom信息，存储到DB中
	PostWisdom(wisdom ...*entity.Wisdom) error
}

// WisdomService Wisdom服务依赖的基础设施仓储接口
type WisdomService struct {
	fileInfra   repository.WisdomFileRepos
	dbsInfra    repository.WisdomDBRepos
	openAIInfra repository.WisdomOpenAIRepos
	cacheInfra  repository.WisdomCacheRepos
}

//
// func (wisSrv *WisdomService) GetWisdom(limit int) ([]*entity.GetOneWisdom, error) {
// 	var wisdoms []*entity.GetOneWisdom
//
// 	return wisdoms, nil
// }
//
// func (wisSrv *WisdomService) PostWisdom(ents ...*entity.GetOneWisdom) error {
// 	var wisdoms []*entity.GetOneWisdom
//
// 	for _, wis := range ents {
// 		wisdoms = append(wisdoms, wis)
// 	}
//
// 	err := wisSrv.dbsInfra.InsertWisdom(wisdoms)
// 	if err != nil {
// 		return errors.Wrap(err, "post wisdom got err")
// 	}
//
// 	return nil
// }
