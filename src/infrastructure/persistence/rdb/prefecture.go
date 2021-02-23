package rdb

import (
	"github.com/MSHR-Dec/pantogram/src/domain/model"
	"github.com/MSHR-Dec/pantogram/src/domain/repository"
	"gorm.io/gorm"
)

type PrefectureRDBRepository struct {
	db *gorm.DB
}

func NewPrefectureRDBRepository(d *gorm.DB) repository.PrefectureRDBRepository {
	return &PrefectureRDBRepository{
		db: d,
	}
}

func (r *PrefectureRDBRepository) GetAllPrefectures() (prefectures []*model.Prefecture) {
	r.db.Find(&prefectures)
	return
}
