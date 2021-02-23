package model

import "time"

type DailyRouteSummary struct {
	ID             int `gorm:"primary_key;not null"`
	RouteID        Route
	DelayCount     int       `gorm:"not null"`
	DelayDuration  time.Time `gorm:"not null"`
	RegisteredDate time.Time `sql:"not null;type:date"`
}
