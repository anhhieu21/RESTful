package model

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

type JwtCustomClaims struct {
	UserId uuid.UUID
	jwt.StandardClaims
}