package repository

import "time"

type PrefectureRepository interface {
	Save(t map[string]string, f map[string]interface{}, n time.Time)
}
