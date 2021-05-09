package application

import (
	"github.com/MSHR-Dec/pantogram/datastore/domain/model"
	"github.com/MSHR-Dec/pantogram/datastore/domain/repository"
)

type RouteDetail interface {
	GetRouteDetail(name string) (output RouteDetailOutput, err error)
}

type RouteDetailInteractor struct {
	routeRepository      repository.RouteRepository
	companyRepository    repository.CompanyRepository
	prefectureRepository repository.PrefectureRepository
}

func NewRouteDetailInteractor(route repository.RouteRepository,
	company repository.CompanyRepository,
	prefecture repository.PrefectureRepository,
) *RouteDetailInteractor {
	return &RouteDetailInteractor{
		routeRepository:      route,
		companyRepository:    company,
		prefectureRepository: prefecture,
	}
}

type RouteDetailOutput struct {
	Company     model.Company
	Prefectures []model.Prefecture
}

func (r RouteDetailInteractor) GetRouteDetail(name string) (output RouteDetailOutput, err error) {
	route, err := r.routeRepository.GetRoute(name)
	if err != nil {
		return
	}

	company := r.companyRepository.GetCompany(route.CompanyID)

	prefectureIDs := route.ListPrefectureIDs()
	prefectures := r.prefectureRepository.GetPrefectures(prefectureIDs)

	output.Company = company
	output.Prefectures = prefectures

	return
}
