package repository

import "github.com/MSHR-Dec/pantogram/src/domain/model"

type CompanyRDBRepository interface {
	GetCompanies(companyIDs []int) (companies []*model.Company)
}
