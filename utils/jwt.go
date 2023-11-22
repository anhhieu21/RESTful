package utils

import (
	"main/model"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const SECRET_KEY = "SCnJhUUVYTSmDx0ye293Rd90qdLdRRRo"

func GenToken(user model.Users) (string, error) {
	claims := &model.JwtCustomClaims{
		UserId: user.Id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	result, err := token.SignedString([]byte(SECRET_KEY))
	if err != nil {
		return "", err
	}

	return result, nil
}
