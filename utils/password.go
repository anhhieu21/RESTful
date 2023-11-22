package utils

import (
	"main/helper"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password []byte) string {
	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.MinCost)
	if err != nil {
		helper.ErrorPanic(err)
	}
	return string(hash)
}
