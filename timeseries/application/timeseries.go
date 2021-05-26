package application

import (
	"time"

	"github.com/MSHR-Dec/pantogram/timeseries/domain/repository"
)

type Route struct {
	Name     string
	ComName  string
	PrefName []string
}

type Company struct {
	Name string
}

type Prefecture struct {
	Name  string
	Count int
}

type Timeseries interface {
	SaveTimeseries(routes []Route, companies []Company, prefectures []Prefecture)
}

type TimeseriesInteractor struct {
	routeRepository      repository.RouteRepository
	companyRepository    repository.CompanyRepository
	prefectureRepository repository.PrefectureRepository
}

func NewTimeseriesInteractor(route repository.RouteRepository,
	company repository.CompanyRepository,
	prefecture repository.PrefectureRepository,
) *TimeseriesInteractor {
	return &TimeseriesInteractor{
		routeRepository:      route,
		companyRepository:    company,
		prefectureRepository: prefecture,
	}
}

func (t TimeseriesInteractor) SaveTimeseries(routes []Route,
	companies []Company,
	prefectures []Prefecture) {
	route := routeApplication{registeredAt: time.Now()}
	company := companyApplication{registeredAt: time.Now()}
	prefecture := prefectureApplication{registeredAt: time.Now()}

	route.save(t.routeRepository, routes)
	company.save(t.companyRepository, companies)
	prefecture.save(t.prefectureRepository, prefectures)
}
