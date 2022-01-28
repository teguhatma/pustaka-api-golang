package services

import (
	"pustaka-api/repository"
	"pustaka-api/schemas"
)

type BookService interface {
	FindBookAll() ([]repository.Book, error)
	FindBookById(ID int) (repository.Book, error)
	CreateBook(book schemas.BookRequest) (repository.Book, error)
	DeleteBook(ID int) (repository.Book, error)
	UpdateBook(book schemas.BookRequest, ID int) (repository.Book, error)
}

type bookService struct {
	repository repository.BookRepository
}

func BookNewService(repository repository.BookRepository) *bookService {
	return &bookService{repository}
}

func (s *bookService) FindBookAll() ([]repository.Book, error) {
	return s.repository.FindBookAll()
}

func (s *bookService) FindBookById(ID int) (repository.Book, error) {
	return s.repository.FindBookById(ID)
}

func (s *bookService) CreateBook(bookRequest schemas.BookRequest) (repository.Book, error) {
	book := repository.Book{
		Title:       bookRequest.Title,
		Price:       bookRequest.Price,
		Description: bookRequest.Description,
		Rating:      bookRequest.Rating,
		UserID:      bookRequest.UserID,
	}
	newBook, err := s.repository.CreateBook(book)
	return newBook, err
}

func (s *bookService) DeleteBook(ID int) (repository.Book, error) {
	return s.repository.DeleteBook(ID)
}

func (s *bookService) UpdateBook(bookRequest schemas.BookRequest, ID int) (repository.Book, error) {
	book := repository.Book{
		Title:       bookRequest.Title,
		Price:       bookRequest.Price,
		Description: bookRequest.Description,
		Rating:      bookRequest.Rating,
		UserID:      bookRequest.UserID,
	}
	newBook, err := s.repository.UpdateBook(book, ID)
	return newBook, err
}
