package repository

import (
	"time"

	"gorm.io/gorm"
)

type BookRepository interface {
	FindBookAll() ([]Book, error)
	FindBookById(ID int) (Book, error)
	CreateBook(book Book) (Book, error)
	DeleteBook(ID int) (Book, error)
	UpdateBook(book Book, ID int) (Book, error)
}

type Book struct {
	gorm.Model
	UserID      int
	ID          uint
	Title       string
	Description string
	Price       int
	Rating      int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (r *DBRepository) FindBookAll() ([]Book, error) {
	var books []Book

	err := r.db.Find(&books).Error

	return books, err
}

func (r *DBRepository) FindBookById(ID int) (Book, error) {
	var book Book

	err := r.db.First(&book, ID).Error

	return book, err
}

func (r *DBRepository) CreateBook(book Book) (Book, error) {
	err := r.db.Create(&book).Error

	return book, err
}

func (r *DBRepository) DeleteBook(ID int) (Book, error) {
	var book Book

	err := r.db.First(&book, ID).Delete(ID).Error

	return book, err
}

func (r *DBRepository) UpdateBook(b Book, ID int) (Book, error) {
	var book Book

	err := r.db.First(&book, ID).Updates(&b).Error

	return book, err
}
