package schemas

type Authentication struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Token struct {
	Email       string `json:"email"`
	TokenString string `json:"token"`
}

type APIResponseToken struct {
	Code    int    `example:"200"`
	Message string `example:"success"`
	Token   string
}
