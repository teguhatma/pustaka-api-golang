package schemas

import "time"

type APIResponseUser struct {
	Code    int    `example:"200"`
	Message string `example:"message"`
	Data    UserResponse
}

type UserRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"email,required"`
	Password string `json:"password" binding:"required"`
}

type UserResponse struct {
	Id        int
	Name      string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
