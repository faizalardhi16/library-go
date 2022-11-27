package models

import "time"

type Transaction struct {
	ID        string `gorm:"size:36;not null;uniqueIndex;primary_key"`
	UserID    string `gorm:"size:36"`
	Status    string `gorm:"size:100"`
	CreateAt  time.Time
	UpdatedAt time.Time
}
