package schemas

import "time"

type APIResponseUser struct {
	Code    int    `example:"200"`
	Message string `example:"success"`
	Data    UserResponse
}

type UserRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"email,required"`
	Password string `json:"password" binding:"required"`
}

type UserResponse struct {
	ID        int       `json:"ID"`
	Name      string    `json:"Name"`
	Email     string    `json:"Email"`
	CreatedAt time.Time `json:"CreatedAt"`
	UpdatedAt time.Time `json:"UpdatedAt"`
}

type APIResponseUsers struct {
	Code    int    `example:"200"`
	Message string `example:"success"`
	Data    []UserResponse
}

type User struct {
	ID        uint
	Name      string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
