package usecase

import (
	"time"

	"github.com/MSHR-Dec/pantogram/src/domain/model"
	"github.com/MSHR-Dec/pantogram/src/domain/repository"
)

type PantogramUsecase interface {
	Create(names []string)
}

type PantogramInteractor struct {
	RouteRDBRepository       repository.RouteRDBRepository
	CompanyRDBRepository     repository.CompanyRDBRepository
	PrefectureRDBRepository  repository.PrefectureRDBRepository
	RouteTSDBRepository      repository.RouteTSDBRepository
	CompanyTSDBRepository    repository.CompanyTSDBRepository
	PrefectureTSDBRepository repository.PrefectureTSDBRepository
}

func NewPantogramInteractor(
	routeRDBRepository repository.RouteRDBRepository,
	companyRDBRepository repository.CompanyRDBRepository,
	prefectureRDBRepository repository.PrefectureRDBRepository,
	routeTSDBRepository repository.RouteTSDBRepository,
	companyTSDBRepository repository.CompanyTSDBRepository,
	prefectureTSDBRepository repository.PrefectureTSDBRepository,
) PantogramUsecase {
	return &PantogramInteractor{
		RouteRDBRepository:       routeRDBRepository,
		CompanyRDBRepository:     companyRDBRepository,
		PrefectureRDBRepository:  prefectureRDBRepository,
		RouteTSDBRepository:      routeTSDBRepository,
		CompanyTSDBRepository:    companyTSDBRepository,
		PrefectureTSDBRepository: prefectureTSDBRepository,
	}
}

func (i *PantogramInteractor) Create(names []string) {
	n := time.Now()

	routeDelayDataIO := newRouteDelayDataIO(names, n)
	routes := routeDelayDataIO.getDelayDatas(i.RouteRDBRepository)
	routeDelayDataIO.createDelayDatas(i.RouteTSDBRepository, routes)

	ids := i.listIDs(routes)

	companyDelayDataIO := newCompanyDelayDataIO(ids.companyIDs, n)
	companies := companyDelayDataIO.getDelayDatas(i.CompanyRDBRepository)
	companyDelayDataIO.createDelayDatas(i.CompanyTSDBRepository, companies)

	prefectureDelayDataIO := newPrefectureDelayDataIO(ids.prefectureIDs, n)
	prefectures := prefectureDelayDataIO.getDelayDatas(i.PrefectureRDBRepository)
	prefectureDelayDataIO.createDelayDatas(i.PrefectureTSDBRepository, prefectures)
}

type IDs struct {
	companyIDs    []int
	prefectureIDs []uint8
}

func (i *PantogramInteractor) listIDs(routes []*model.Route) (ids IDs) {
	for _, route := range routes {
		ids.companyIDs = append(ids.companyIDs, route.GetCompanyID())
		ids.prefectureIDs = route.ListPrefectureIDs(ids.prefectureIDs)
	}

	return
}
