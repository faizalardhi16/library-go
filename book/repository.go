package book

import "gorm.io/gorm"

type Repository interface {
	CreateBook(book Book) (Book, error)
	FindBook() ([]Book, error)
	UpdateBook(book Book) (Book, error)
	FindByID(bookID string) (Book, error)
}

type repository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreateBook(book Book) (Book, error) {
	err := r.db.Create(&book).Error

	if err != nil {
		return book, err
	}

	return book, nil
}

func (r *repository) FindBook() ([]Book, error) {
	var book []Book

	err := r.db.Find(&book).Error

	if err != nil {
		return book, err
	}

	return book, nil
}

func (r *repository) FindByID(bookID string) (Book, error) {
	var book Book

	err := r.db.Where("id = ?", bookID).Find(&book).Error

	if err != nil {
		return book, err
	}

	return book, nil
}

func (r *repository) UpdateBook(book Book) (Book, error) {
	err := r.db.Save(&book).Error

	if err != nil {
		return book, err
	}

	return book, nil
}
