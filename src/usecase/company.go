package usecase

import (
	"time"

	"github.com/MSHR-Dec/pantogram/src/domain/model"
	"github.com/MSHR-Dec/pantogram/src/domain/repository"
	"github.com/MSHR-Dec/pantogram/src/domain/service"
)

type companyDelayDataIO struct {
	companyIDs   []int
	registeredAt time.Time
}

func newCompanyDelayDataIO(input []int, n time.Time) *companyDelayDataIO {
	return &companyDelayDataIO{
		companyIDs:   input,
		registeredAt: n,
	}
}

func (i *companyDelayDataIO) getDelayDatas(repo repository.CompanyRDBRepository) (companies []*model.Company) {
	companies = repo.GetCompanies(i.companyIDs)
	return
}

func (i *companyDelayDataIO) createDelayDatas(repo repository.CompanyTSDBRepository, companies []*model.Company) {
	for _, company := range companies {
		company.AddCount(i.companyIDs)
		tsdb := model.NewCompanyTSDB(company.Name)
		repo.Save(service.ToMap(tsdb),
			map[string]interface{}{"count": company.Count},
			i.registeredAt)
	}
}
