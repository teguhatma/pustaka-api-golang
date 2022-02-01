package auths

import (
	"pustaka-api/config/database"
	"pustaka-api/repository"
	"pustaka-api/utils/auth_jwt"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func Authorize(c *gin.Context) (bool, error) {
	db := database.DatabaseConn()
	var user repository.User

	// Get token from header
	tokenString := c.Request.Header["Token"]
	if len(tokenString) == 0 || tokenString[0] == "" {
		return false, nil
	}

	// Decode Token
	decodeToken, err := auth_jwt.DecodeJWT(tokenString[0])

	// Get email from decode token
	val := decodeToken.Claims.(jwt.MapClaims)
	email := val["email"]

	// Get user by email
	db.Where("email=?", email).First(&user)

	if user.Token != tokenString[0] {
		return false, nil
	}

	return true, err
}
