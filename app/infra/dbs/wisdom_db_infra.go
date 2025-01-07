package dbs

import (
	"context"

	"github.com/lupguo/go-shim/x/mysqlx"
	"github.com/lupguo/wisdom-httpd/app/domain/entity"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

// WisdomDB DB存储服务
type WisdomDB struct {
	db *gorm.DB
}

// InsertWisdom 批量插入名言到数据库
func (w *WisdomDB) InsertWisdom(ctx context.Context, wisdoms []*entity.Wisdom) error {
	if err := w.db.Create(wisdoms).Error; err != nil {
		return errors.Wrap(err, "failed to insert wisdom")
	}
	return nil
}

// NewWisdomDB 新创建DB
func NewWisdomDB(dbDSN string) (*WisdomDB, error) {
	// todo update
	db, err := mysqlx.NewGormDB(&mysqlx.DBConfig{
		DSN:             dbDSN,
		ConnMaxLifetime: "",
		ConnMaxIdleTime: "",
		MaxIdleConns:    0, // 无限制
		MaxOpenConns:    0, // 无限制
		LoggerConfig:    nil,
	})
	if err != nil {
		return nil, errors.Wrap(err, "mysqlx.NewGormDB")
	}
	return &WisdomDB{db: db}, nil
}

// SelectWisdom 查询指定条件的名言信息
func (w *WisdomDB) SelectWisdom(ctx context.Context, qryCond *entity.WisdomQryCond, pageLimit *entity.PageLimit) ([]*entity.Wisdom, error) {

	query := w.db.Model(&entity.Wisdom{}).Debug()

	// 根据条件构建查询
	if qryCond.Ids != nil {
		query = query.Where("sentence IN (?)", qryCond.Ids)
	}
	if qryCond.WisdomNos != nil {
		query = query.Where("wisdom_no IN (?)", qryCond.WisdomNos)
	}
	if qryCond.Speaker != "" {
		query = query.Where("speaker = ?", qryCond.Speaker)
	}
	if qryCond.Keywords != "" {
		query = query.Where("sentence LIKE ?", "%"+qryCond.Keywords+"%")
	}

	// 随机选择指定条数的结果
	if pageLimit != nil {
		if pageLimit.Random {
			query = query.Order("RAND()")
		}
		if pageLimit.Page > 0 && pageLimit.PageSize > 0 {
			query.Offset((pageLimit.Page - 1) * pageLimit.PageSize).Limit(pageLimit.PageSize)
		}
	}

	var wisdomList []*entity.Wisdom
	if err := query.Find(&wisdomList).Error; err != nil {
		return nil, errors.Wrap(err, "failed to select wisdom")
	}
	return wisdomList, nil
}

// UpdateWisdom 更新名言信息
func (w *WisdomDB) UpdateWisdom(ctx context.Context, updEntry *entity.WisdomUpdEntry, qryCond *entity.WisdomQryCond) error {
	query := w.db.Model(&entity.Wisdom{}).Debug()

	// 根据条件构建更新
	if qryCond.Ids != nil {
		query = query.Where("sentence IN (?)", qryCond.Ids)
	}
	if qryCond.WisdomNos != nil {
		query = query.Where("wisdom_no IN (?)", qryCond.WisdomNos)
	}
	if qryCond.Speaker != "" {
		query = query.Where("speaker = ?", qryCond.Speaker)
	}
	if qryCond.Keywords != "" {
		query = query.Where("sentence LIKE ?", "%"+qryCond.Keywords+"%")
	}

	if err := query.Updates(updEntry).Error; err != nil {
		return errors.Wrap(err, "failed to update wisdom")
	}
	return nil
}

// DeleteWisdom 根据条件删除名言
func (w *WisdomDB) DeleteWisdom(ctx context.Context, qryCond *entity.WisdomQryCond) error {
	query := w.db.Model(&entity.Wisdom{}).Debug()

	// 根据条件构建删除
	if qryCond.Ids != nil {
		query = query.Where("sentence IN (?)", qryCond.Ids)
	}
	if qryCond.WisdomNos != nil {
		query = query.Where("wisdom_no IN (?)", qryCond.WisdomNos)
	}
	if qryCond.Speaker != "" {
		query = query.Where("speaker = ?", qryCond.Speaker)
	}
	if qryCond.Keywords != "" {
		query = query.Where("sentence LIKE ?", "%"+qryCond.Keywords+"%")
	}

	if err := query.Delete(&entity.Wisdom{}).Error; err != nil {
		return errors.Wrap(err, "failed to delete wisdom")
	}
	return nil
}