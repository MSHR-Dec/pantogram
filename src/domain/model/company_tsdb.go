package model

type CompanyTSDB struct {
	Name string
}

func NewCompanyTSDB(name string) *CompanyTSDB {
	return &CompanyTSDB{
		Name: name,
	}
}
