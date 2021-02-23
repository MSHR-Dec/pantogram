package model

type Region struct {
	ID   uint8  `gorm:"primary_key;not null"`
	Name string `gorm:"type:varchar(255);default:null"`
}
