package repository

import "time"

type User struct {
	ID        uint
	Name      string 	
	Email     string	`gorm:"uniqueIndex"`
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserRepository interface {
	FindUsers() ([]User, error)
	FindUserById(ID int) (User, error)
	CreateUser(user User) (User, error)
	DeleteUser(ID int) (User, error)
	UpdateUser(user User, ID int) (User, error)
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

func (r *DBRepository) FindUsers() ([]User, error) {
	var users []User

	err := r.db.Select("id", "name", "email", "created_at", "updated_at").Find(&users).Error

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
