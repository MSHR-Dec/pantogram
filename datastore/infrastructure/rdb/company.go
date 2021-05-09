package rdb

import (
	"github.com/MSHR-Dec/pantogram/datastore/domain/model"
	"gorm.io/gorm"
)

type CompanyRDBRepository struct {
	db *gorm.DB
}

func NewCompanyRepository(d *gorm.DB) *CompanyRDBRepository {
	return &CompanyRDBRepository{
		db: d,
	}
}

func (r *CompanyRDBRepository) GetCompany(companyID int) model.Company {
	var company model.Company
	r.db.First(&company, companyID)
	return company
}
