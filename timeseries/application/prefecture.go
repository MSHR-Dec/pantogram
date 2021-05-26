package application

import (
	"time"

	"github.com/MSHR-Dec/pantogram/timeseries/domain/model"
	"github.com/MSHR-Dec/pantogram/timeseries/domain/repository"
	"github.com/MSHR-Dec/pantogram/timeseries/domain/service"
)

type prefectureApplication struct {
	registeredAt time.Time
}

func (a prefectureApplication) save(repo repository.PrefectureRepository, prefectures []Prefecture) {
	var prefModels map[string]*model.Prefecture
	prefModels = map[string]*model.Prefecture{}

	for _, prefecture := range prefectures {
		if _, ok := prefModels[prefecture.Name]; ok {
			prefModels[prefecture.Name].Field.Count++
			continue
		}
		prefModels[prefecture.Name] = &model.Prefecture{
			Tag: model.PrefectureTag{
				Name: prefecture.Name,
			},
			Field: model.PrefectureField{
				Total: prefecture.Count,
				Count: 1,
			},
		}
	}

	for _, prefecture := range prefModels {
		prefecture.Field.CalcRatio()

		repo.Save(service.ToMap(prefecture.Tag),
			map[string]interface{}{"count": prefecture.Field.Count, "ratio": prefecture.Field.Ratio},
			a.registeredAt,
		)
	}
}
