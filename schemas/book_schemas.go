package schemas

type APIResponseID struct {
	Code    int
	Message string
	ID      int
}

type APIResponse404 struct {
	Code    int
	Message string
}

type APIResponse400 struct {
	Code    int
	Message string
}

type APIResponseTitle struct {
	Code    int
	Message string
	Title   string
}

type BookRequest struct {
	Title       string `json:"title" binding:"required"`
	Price       int    `json:"price" binding:"required,number"`
	Description string `json:"description"`
	Rating      int    `json:"rating"`
}
