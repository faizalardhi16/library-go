package models

import "time"

type Transaction struct {
	ID            string `gorm:"size:36;not null;uniqueIndex;primary_key"`
	UserID        string `gorm:"size:36"`
	BookID        string `gorm:"size:36"`
	ReservationID string `gorm:"size:36"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
