package auth_jwt

import (
	"errors"
	"pustaka-api/utils"

	"github.com/golang-jwt/jwt/v4"
)

func DecodeJWT(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(utils.EnvVariable("SECRET_KEY")), nil
	})

	return token, err
}
