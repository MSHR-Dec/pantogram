package application

import (
	"time"

	"github.com/MSHR-Dec/pantogram/timeseries/domain/model"
	"github.com/MSHR-Dec/pantogram/timeseries/domain/repository"
	"github.com/MSHR-Dec/pantogram/timeseries/domain/service"
)

type companyApplication struct {
	registeredAt time.Time
}

func (a companyApplication) save(repo repository.CompanyRepository, companies []Company) {
	var comModels map[string]*model.Company
	comModels = map[string]*model.Company{}

	for _, company := range companies {
		if _, ok := comModels[company.Name]; ok {
			comModels[company.Name].Field.Count++
			continue
		}
		comModels[company.Name] = &model.Company{
			Tag: model.CompanyTag{
				Name: company.Name,
			},
			Field: model.CompanyField{
				Count: 1,
			},
		}
	}

	for _, company := range comModels {
		repo.Save(service.ToMap(company.Tag),
			map[string]interface{}{"count": company.Field.Count},
			a.registeredAt)
	}
}
