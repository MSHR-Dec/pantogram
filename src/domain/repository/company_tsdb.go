package repository

import "time"

type CompanyTSDBRepository interface {
	Save(t map[string]string, f map[string]interface{}, n time.Time)
}
