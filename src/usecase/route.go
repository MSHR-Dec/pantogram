package usecase

import (
	"time"

	"github.com/MSHR-Dec/pantogram/src/domain/model"
	"github.com/MSHR-Dec/pantogram/src/domain/repository"
	"github.com/MSHR-Dec/pantogram/src/domain/service"
)

type routeDelayDataIO struct {
	names        []string
	registeredAt time.Time
}

func newRouteDelayDataIO(input []string, n time.Time) *routeDelayDataIO {
	return &routeDelayDataIO{
		names:        input,
		registeredAt: n,
	}
}

func (i *routeDelayDataIO) getDelayDatas(repo repository.RouteRDBRepository) (routes []*model.Route) {
	routes = repo.GetRoutes(i.names)
	return
}

func (i *routeDelayDataIO) createDelayDatas(repo repository.RouteTSDBRepository, routes []*model.Route) {
	for _, route := range routes {
		tsdb := model.NewRouteTSDB(route.Name, route.Company.Name, route.ToString())
		repo.Save(service.ToMap(tsdb),
			map[string]interface{}{"count": 1},
			i.registeredAt)
	}
}
