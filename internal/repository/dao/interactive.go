package dao

import (
	"context"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

type InteractiveDao interface {
	IncrReadCnt(ctx context.Context, bizID int64, biz string) error
}

var _ InteractiveDao = (*GormInteractiveDao)(nil)

type GormInteractiveDao struct {
	db *gorm.DB
}

func NewInteractiveDao(db *gorm.DB) InteractiveDao {
	return &GormInteractiveDao{db: db}
}

func (dao *GormInteractiveDao) IncrReadCnt(ctx context.Context, bizID int64, biz string) error {
	now := time.Now().UnixMilli()
	return dao.db.WithContext(ctx).Clauses(clause.OnConflict{
		Columns: []clause.Column{
			{Name: "id"},
		},
		DoUpdates: clause.Assignments(map[string]interface{}{
			"updated_at": now,
			"read_cnt":   gorm.Expr("read_cnt + ?", 1),
		}),
	}).Create(&Interactive{
		BizID:     bizID,
		Biz:       biz,
		ReadCnt:   1,
		CreatedAt: now,
		UpdatedAt: now,
	}).Error
}

type Interactive struct {
	ID         int64  `gorm:"primary_key;AUTO_INCREMENT"`
	BizID      int64  `gorm:"uniqueIndex:biz_type_id"`
	Biz        string `gorm:"type:varchar(128);uniqueIndex:biz_type_id"`
	ReadCnt    int
	LikeCnt    int
	CollectCnt int
	CreatedAt  int64
	UpdatedAt  int64
}
