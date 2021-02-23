package repository

import "time"

type RouteTSDBRepository interface {
	Save(t map[string]string, f map[string]interface{}, n time.Time)
}
