package middlewares

import (
	"errors"
	"main/model"
	"main/utils"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenString := context.GetHeader("Authorization")
		if tokenString == "" {
			context.JSON(401, gin.H{"error": "request does not contain an access token"})
			context.Abort()
			return
		}
		err := ValidateToken(tokenString)
		if err != nil {
			context.JSON(401, gin.H{"error": err.Error()})
			context.Abort()
			return
		}
		context.Next()
	}
}
func ValidateToken(signedToken string) (err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&model.JwtCustomClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(utils.SECRET_KEY), nil
		},
	)
	if err != nil {
		return
	}
	claims, ok := token.Claims.(*model.JwtCustomClaims)
	if !ok {
		err = errors.New("couldn't parse claims")
		return
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return
	}
	return
}
