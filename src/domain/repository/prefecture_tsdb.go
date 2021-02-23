package repository

import "time"

type PrefectureTSDBRepository interface {
	Save(t map[string]string, f map[string]interface{}, n time.Time)
}
