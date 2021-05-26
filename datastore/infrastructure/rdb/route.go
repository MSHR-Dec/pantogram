package rdb

import (
	"log"

	"gorm.io/gorm"

	"github.com/MSHR-Dec/pantogram/datastore/domain/model"
)

type RouteRDBRepository struct {
	db *gorm.DB
}

func NewRouteRepository(d *gorm.DB) *RouteRDBRepository {
	return &RouteRDBRepository{
		db: d,
	}
}

func (r *RouteRDBRepository) GetRoute(name string) (model.Route, error) {
	var route model.Route
	if err := r.db.Preload("Prefectures").Where("name = ?", name).Take(&route).Error; err != nil {
		log.Println(err)
		return model.Route{}, err
	}
	return route, nil
}
