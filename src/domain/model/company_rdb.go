package model

type Company struct {
	ID    int    `gorm:"primary_key;not null"`
	Name  string `gorm:"type:varchar(255);default:null"`
	Count int
}

func (m *Company) GetName() string {
	return m.Name
}

func (m *Company) AddCount(companyIDs []int) {
	for _, companyID := range companyIDs {
		if m.ID == companyID {
			m.Count++
		}
	}
}
