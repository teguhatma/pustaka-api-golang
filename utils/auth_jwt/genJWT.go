package auth_jwt

import (
	"fmt"
	"time"

	"pustaka-api/utils"

	"github.com/golang-jwt/jwt/v4"
)

func GenerateJWT(email string) (string, error) {
	var mySigningKey = []byte(utils.EnvVariable("SECRET_KEY"))
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["email"] = email
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		fmt.Errorf("Something Went Wrong: %s", err.Error())
		return "", err
	}
	return tokenString, nil
}
