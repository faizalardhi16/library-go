package models

import "time"

type Reservation struct {
	ID          string `gorm:"size:36;not null;uniqueIndex;primary_key"`
	UserID      string `gorm:"size:36;not null"`
	BookingDate string `gorm:"size:100;not null"`
	CreateAt    time.Time
	UpdatedAt   time.Time
}
