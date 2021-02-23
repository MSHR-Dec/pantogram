package model

import "strings"

type Route struct {
	ID          int    `gorm:"primary_key;not null"`
	Name        string `gorm:"type:varchar(255);default:null"`
	CompanyID   int
	IsDelay     bool         `gorm:"type:bool;default:false"`
	Prefectures []Prefecture `gorm:"many2many:routes_prefectures"`
	Company     Company
}

func (m *Route) ToString() string {
	var prefArr []string
	for _, prefecture := range m.Prefectures {
		prefArr = append(prefArr, prefecture.Name)
	}

	return strings.Join(prefArr, ", ")
}

func (m *Route) GetCompanyID() int {
	return m.CompanyID
}

func (m *Route) ListPrefectureIDs(prefectures []uint8) []uint8 {
	for _, prefecture := range m.Prefectures {
		prefectures = append(prefectures, prefecture.ID)
	}
	return prefectures
}
