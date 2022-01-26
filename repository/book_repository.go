package repository

import (
	"time"

	"gorm.io/gorm"
)

type dbRepository struct {
	db *gorm.DB
}

type BookRepository interface {
	FindAll() ([]Book, error)
	FindById(ID int) (Book, error)
	Create(book Book) (Book, error)
	Delete(ID int) (Book, error)
	Update(book Book, ID int) (Book, error)
}

type Book struct {
	ID          uint
	Title       string
	Description string
	Price       int
	Rating      int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewRepository(db *gorm.DB) *dbRepository {
	return &dbRepository{db}
}

func (r *dbRepository) FindAll() ([]Book, error) {
	var books []Book

	err := r.db.Find(&books).Error

	return books, err
}

func (r *dbRepository) FindById(ID int) (Book, error) {
	var book Book

	err := r.db.First(&book, ID).Error

	return book, err
}

func (r *dbRepository) Create(book Book) (Book, error) {
	err := r.db.Create(&book).Error

	return book, err
}

func (r *dbRepository) Delete(ID int) (Book, error) {
	var book Book

	err := r.db.First(&book, ID).Delete(ID).Error

	return book, err
}

func (r *dbRepository) Update(b Book, ID int) (Book, error) {
	var book Book

	err := r.db.First(&book, ID).Updates(&b).Error

	return book, err
}
