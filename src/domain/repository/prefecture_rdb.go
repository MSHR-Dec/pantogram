package repository

import (
	"github.com/MSHR-Dec/pantogram/src/domain/model"
)

type PrefectureRDBRepository interface {
	GetAllPrefectures() (prefectures []*model.Prefecture)
}
