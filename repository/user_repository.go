package repository

import "time"

type User struct {
	ID        uint
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserRepository interface {
	FindUserById(ID int) (User, error)
	CreateUser(user User) (User, error)
}

func (r *DBRepository) FindUserById(ID int) (User, error) {
	var user User

	err := r.db.Select("id", "name", "email", "created_at", "updated_at").First(&user, ID).Error

	return user, err
}

func (r *DBRepository) CreateUser(user User) (User, error) {
	err := r.db.Create(&user).Error

	return user, err
}
