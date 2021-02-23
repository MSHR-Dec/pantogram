package usecase

import (
	"math"
	"time"

	"github.com/MSHR-Dec/pantogram/src/domain/model"
	"github.com/MSHR-Dec/pantogram/src/domain/repository"
	"github.com/MSHR-Dec/pantogram/src/domain/service"
)

type prefectureDelayDataIO struct {
	prefectureIDs []uint8
	registeredAt  time.Time
}

func newPrefectureDelayDataIO(input []uint8, n time.Time) *prefectureDelayDataIO {
	return &prefectureDelayDataIO{
		prefectureIDs: input,
		registeredAt:  n,
	}
}

func (i *prefectureDelayDataIO) getDelayDatas(repo repository.PrefectureRDBRepository) (prefectures []*model.Prefecture) {
	repo.GetAllPrefectures()
	return
}

func (i *prefectureDelayDataIO) createDelayDatas(repo repository.PrefectureTSDBRepository, prefectures []*model.Prefecture) {
	for _, prefecture := range prefectures {
		prefecture.AddCount(i.prefectureIDs)
		ratio := float64(prefecture.DelayCount) / float64(prefecture.RouteCount) * 100
		tsdb := model.NewPrefectureTSDB(prefecture.Name)
		repo.Save(service.ToMap(tsdb),
			map[string]interface{}{"count": prefecture.DelayCount, "ratio": math.Round(ratio*100) / 100},
			i.registeredAt)
	}
}
