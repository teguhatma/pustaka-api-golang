package services

import (
	"pustaka-api/repository"
	"pustaka-api/schemas"
)

type BookService interface {
	FindAll() ([]repository.Book, error)
	FindById(ID int) (repository.Book, error)
	Create(book schemas.BookRequest) (repository.Book, error)
	Delete(ID int) (repository.Book, error)
	Update(book schemas.BookRequest, ID int) (repository.Book, error)
}

type service struct {
	repository repository.BookRepository
}

func NewService(repository repository.BookRepository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]repository.Book, error) {
	return s.repository.FindAll()
}

func (s *service) FindById(ID int) (repository.Book, error) {
	return s.repository.FindById(ID)
}

func (s *service) Create(bookRequest schemas.BookRequest) (repository.Book, error) {
	book := repository.Book{
		Title:       bookRequest.Title,
		Price:       bookRequest.Price,
		Description: bookRequest.Description,
		Rating:      bookRequest.Rating,
	}
	newBook, err := s.repository.Create(book)
	return newBook, err
}

func (s *service) Delete(ID int) (repository.Book, error) {
	return s.repository.Delete(ID)
}

func (s *service) Update(bookRequest schemas.BookRequest, ID int) (repository.Book, error) {
	book := repository.Book{
		Title:       bookRequest.Title,
		Price:       bookRequest.Price,
		Description: bookRequest.Description,
		Rating:      bookRequest.Rating,
	}
	newBook, err := s.repository.Update(book, ID)
	return newBook, err
}
