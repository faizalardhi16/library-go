package book

type CreateBookFormat struct {
	ID          string `json:"id"`
	FileName    string `json:"file_name"`
	Title       string `json:"title"`
	Penulis     string `json:"penulis"`
	TahunTerbit string `json:"tahun_terbit"`
}

func CreateBookFormatter(book Book) CreateBookFormat {
	f := CreateBookFormat{}

	f.ID = book.ID
	f.FileName = book.FileName
	f.Title = book.Title
	f.Penulis = book.Penulis
	f.TahunTerbit = book.TahunTerbit

	return f
}

func BookFormat(book []Book) []CreateBookFormat {
	booksFormats := []CreateBookFormat{}

	for _, r := range book {
		booksFormat := CreateBookFormatter(r)
		booksFormats = append(booksFormats, booksFormat)
	}

	return booksFormats
}
