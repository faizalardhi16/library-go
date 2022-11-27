package models

import "time"

type Book struct {
	ID          string `gorm:"size:36;not null;uniqueIndex;primary_key"`
	FileName    string `gorm:"size:100;not null"`
	Title       string `gorm:"size:100;not null"`
	Penulis     string `gorm:"size:100;not null"`
	TahunTerbit string `gorm:"size:100;not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
