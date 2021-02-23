package model

type RouteTSDB struct {
	Name       string
	Company    string
	Prefecture string
}

func NewRouteTSDB(name string, company string, prefecture string) *RouteTSDB {
	return &RouteTSDB{
		Name:       name,
		Company:    company,
		Prefecture: prefecture,
	}
}
