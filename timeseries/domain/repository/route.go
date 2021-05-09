package repository

import "time"

type RouteRepository interface {
	Save(t map[string]string, f map[string]interface{}, n time.Time)
}
