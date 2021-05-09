package model

type Prefecture struct {
	ID         uint8  `gorm:"primary_key;not null"`
	Name       string `gorm:"type:varchar(255);default:null"`
	RegionID   uint8  `gorm:"not null"`
	RouteCount uint32 `gorm:"default:null"`
	DelayCount int
}

func (m *Prefecture) AddCount(prefectureIDs []uint8) {
	for _, prefectureID := range prefectureIDs {
		if m.ID == prefectureID {
			m.DelayCount++
		}
	}
}
