package book

import "time"

type Book struct {
	ID          string `json:"id"`
	FileName    string `json:"file_name"`
	Title       string `json:"title"`
	Penulis     string `json:"penulis"`
	TahunTerbit string `json:"tahun_terbit"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
