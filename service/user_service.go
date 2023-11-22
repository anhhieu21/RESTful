package service

import (
	"main/data/req"
	"main/model"

	"github.com/google/uuid"
)

type UserService interface {
	Create(user req.UserRequest) model.Users
	Update(user req.UserUpdateRequest)
	FindById(userId uuid.UUID) model.Users
}
