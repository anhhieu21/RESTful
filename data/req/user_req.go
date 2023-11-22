package req

import "github.com/google/uuid"

type UserRequest struct {
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
}

type UserUpdateRequest struct {
	Id       uuid.UUID
	FullName string `json:"full_name"`
	Phone    string `json:"phone"`
}

type UpdatePasswordRequest struct {
	Password string `json:"password"`
}
