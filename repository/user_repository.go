package repository

import (
	"pustaka-api/schemas"
	"time"

	"gorm.io/gorm"
)

// This user struct only for creating model
// Not to response model
type User struct {
	gorm.Model
	Books     []Book `gorm:"foreignKey:UserID;references:ID"`
	ID        uint
	Name      string
	Email     string `gorm:"uniqueIndex"`
	Password  string
	Token     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserRepository interface {
	FindUsers() ([]schemas.User, error)
	FindUserById(ID int) (schemas.User, error)
	CreateUser(user User) (User, error)
	DeleteUser(ID int) (User, error)
	UpdateUser(user User, ID int) (User, error)
	FindUserByEmail(Email string) (User, error)
	SaveUserToken(Token, Email string) error
}

func (r *DBRepository) FindUserById(ID int) (schemas.User, error) {
	var user schemas.User

	err := r.db.First(&user, ID).Error

	return user, err
}

func (r *DBRepository) CreateUser(user User) (User, error) {
	err := r.db.Create(&user).Error

	return user, err
}

func (r *DBRepository) FindUsers() ([]schemas.User, error) {
	var users []schemas.User

	err := r.db.Find(&users).Error

	return users, err
}

func (r *DBRepository) DeleteUser(ID int) (User, error) {
	var user User

	err := r.db.First(&user, ID).Delete(ID).Error

	return user, err
}

func (r *DBRepository) UpdateUser(u User, ID int) (User, error) {
	var user User
	err := r.db.First(&user, ID).Updates(&u).Error

	return user, err
}

func (r *DBRepository) FindUserByEmail(Email string) (User, error) {
	var user User

	err := r.db.Where("email=?", Email).First(&user).Error

	return user, err
}

func (r *DBRepository) SaveUserToken(Token, Email string) error {
	var user User

	err := r.db.Where("email=?", Email).First(&user).Update("token", Token).Error

	return err
}
