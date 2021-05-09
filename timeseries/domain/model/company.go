package model

type Company struct {
	Tag CompanyTag
	Field CompanyField
}

type CompanyTag struct {
	Name string
}

type CompanyField struct {
	Count int
}
