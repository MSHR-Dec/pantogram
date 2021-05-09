package rdb

import (
	"github.com/MSHR-Dec/pantogram/datastore/domain/model"
	"gorm.io/gorm"
)

type PrefectureRDBRepository struct {
	db *gorm.DB
}

func NewPrefectureRepository(d *gorm.DB) *PrefectureRDBRepository {
	return &PrefectureRDBRepository{
		db: d,
	}
}

func (r *PrefectureRDBRepository) GetPrefectures(prefectureIDs []uint8) []model.Prefecture {
	var prefectures []model.Prefecture
	r.db.Find(&prefectures, prefectureIDs)
	return prefectures
}
