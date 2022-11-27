package book

import (
	"errors"

	"github.com/google/uuid"
)

type Service interface {
	CreateBook(input CreateBookInput) (Book, error)
	GetBook() ([]Book, error)
	UpdateBook(bookID string, fileLocation string) (Book, error)
	GetBookByID(bookID string) (Book, error)
}

type service struct {
	repository Repository
}

func NewBookService(repository Repository) *service {
	return &service{repository}
}

func (s *service) CreateBook(input CreateBookInput) (Book, error) {

	book := Book{}

	book.ID = uuid.New().String()
	book.FileName = input.FileName
	book.Penulis = input.Penulis
	book.TahunTerbit = input.TahunTerbit
	book.Title = input.Title

	books, err := s.repository.CreateBook(book)

	if err != nil {
		return books, err
	}

	return books, nil

}

func (s *service) GetBook() ([]Book, error) {
	book, err := s.repository.FindBook()

	if err != nil {
		return book, err
	}

	return book, nil
}

func (s *service) UpdateBook(bookID string, fileLocation string) (Book, error) {
	getBook, err := s.repository.FindByID(bookID)

	if err != nil {
		return getBook, err
	}

	getBook.FileName = fileLocation

	if getBook.ID == "" {
		return getBook, errors.New("Books Not Found")
	}

	updatedBook, err := s.repository.UpdateBook(getBook)

	return updatedBook, err

}

func (s *service) GetBookByID(bookID string) (Book, error) {
	getBook, err := s.repository.FindByID(bookID)

	if err != nil {
		return getBook, err
	}

	return getBook, nil
}
