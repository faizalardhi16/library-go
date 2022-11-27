package book

type CreateBookInput struct {
	FileName    string `json:"file_name"`
	Title       string `json:"title"`
	Penulis     string `json:"penulis"`
	TahunTerbit string `json:"tahun_terbit"`
}

type UpdateInput struct {
	ID string `uri:"id" binding:"required"`
}
