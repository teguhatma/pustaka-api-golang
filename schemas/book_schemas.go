package schemas

import "time"

type APIResponseBooks struct {
	Code    int    `example:"200"`
	Message string `example:"success"`
	Data    []BookResponse
}

type APIResponseBook struct {
	Code    int    `example:"200"`
	Message string `example:"success"`
	Data    BookResponse
}

type BookRequest struct {
	Title       string `json:"title" binding:"required"`
	Price       int    `json:"price" binding:"required,number"`
	Description string `json:"description"`
	Rating      int    `json:"rating"`
	UserID      int    `json:"userId"`
}

type BookResponse struct {
	Id          int
	Price       int
	Description string
	Rating      int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
