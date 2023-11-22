package repository

import (
	"main/model"

	"github.com/google/uuid"
)

type UserRepository interface {
	Create(user model.Users)
	Update(user model.Users)
	FindById(userId uuid.UUID) (model.Users, error)
}
