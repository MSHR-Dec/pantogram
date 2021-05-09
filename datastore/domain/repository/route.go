package repository

import (
	"github.com/MSHR-Dec/pantogram/datastore/domain/model"
)

type RouteRepository interface {
	GetRoute(name string) (model.Route, error)
}
