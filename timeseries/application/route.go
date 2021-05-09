package application

import (
	"strings"
	"time"

	"github.com/MSHR-Dec/pantogram/timeseries/domain/model"
	"github.com/MSHR-Dec/pantogram/timeseries/domain/repository"
	"github.com/MSHR-Dec/pantogram/timeseries/domain/service"
)

type routeApplication struct {
	registeredAt time.Time
}

func (a routeApplication) save(repo repository.RouteRepository, routes []Route) {
	for _, route := range routes {
		prefectures := strings.Join(route.PrefName, ", ")

		routeTag := &model.RouteTag{
			Name:       route.Name,
			Company:    route.ComName,
			Prefecture: prefectures,
		}
		routeField := &model.RouteField{Count: 1}

		repo.Save(service.ToMap(routeTag),
			map[string]interface{}{"count": routeField.Count},
			a.registeredAt)
	}
}
