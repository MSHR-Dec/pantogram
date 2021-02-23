package model

import "time"

type DelayStart struct {
	ID        int       `gorm:"primary_key;not null"`
	StartTime time.Time `gorm:"not null"`
	RouteID   Route
}
