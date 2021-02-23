package rdb

import (
	"github.com/MSHR-Dec/pantogram/src/domain/model"
	"github.com/MSHR-Dec/pantogram/src/domain/repository"
	"gorm.io/gorm"
)

type CompanyRDBRepository struct {
	db *gorm.DB
}

func NewCompanyRDBRepository(d *gorm.DB) repository.CompanyRDBRepository {
	return &CompanyRDBRepository{
		db: d,
	}
}

func (r *CompanyRDBRepository) GetCompanies(companyIDs []int) (companies []*model.Company) {
	r.db.Find(&companies, companyIDs)
	return
}
