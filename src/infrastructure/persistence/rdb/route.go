package rdb

import (
	"github.com/MSHR-Dec/pantogram/src/domain/model"
	"github.com/MSHR-Dec/pantogram/src/domain/repository"
	"gorm.io/gorm"
)

type RouteRDBRepository struct {
	db *gorm.DB
}

func NewRouteRDBRepository(d *gorm.DB) repository.RouteRDBRepository {
	return &RouteRDBRepository{
		db: d,
	}
}

func (r *RouteRDBRepository) GetRoutes(routeNames []string) (routes []*model.Route) {
	r.db.Preload("Prefectures").Preload("Company").Where("name IN (?)", routeNames).Find(&routes)
	return
}
