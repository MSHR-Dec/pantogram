package repository

import (
	"github.com/MSHR-Dec/pantogram/datastore/domain/model"
)

type CompanyRepository interface {
	GetCompany(companyID int) model.Company
}
