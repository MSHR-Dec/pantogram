package repository

import (
	"github.com/MSHR-Dec/pantogram/datastore/domain/model"
)

type PrefectureRepository interface {
	GetPrefectures(prefectureIDs []uint8) []model.Prefecture
}
