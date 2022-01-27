package schemas

type APIResponse404 struct {
	Code    int    `example:"404"`
	Message string `example:"not found"`
	Error   string
}

type APIResponse400 struct {
	Code    int    `example:"400"`
	Message string `example:"bad request"`
	Error   string
}

type APIResponse200 struct {
	Code    int    `example:"200"`
	Message string `example:"success"`
}

type APIResponseDelete200 struct {
	Code    int    `example:"200"`
	Message string `example:"success"`
	Id      int    `example:"1"`
}
