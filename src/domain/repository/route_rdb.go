package repository

import (
	"github.com/MSHR-Dec/pantogram/src/domain/model"
)

type RouteRDBRepository interface {
	GetRoutes(routeNames []string) (routes []*model.Route)
}
