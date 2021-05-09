package repository

import "time"

type CompanyRepository interface {
	Save(t map[string]string, f map[string]interface{}, n time.Time)
}
