package repoimpl

import (
	"errors"
	"main/data/req"
	"main/helper"
	"main/model"
	"main/repository"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	Db *gorm.DB
}

func NewUserRepositoryImpl(Db *gorm.DB) repository.UserRepository {
	return &UserRepositoryImpl{Db: Db}
}

// Create implements repository.UserRepository.
func (u *UserRepositoryImpl) Create(user model.Users) model.Users{
	var data model.Users
	result := u.Db.Create(&user)
	helper.ErrorPanic(result.Error)

	u.Db.Find(&data, user.Id)

	return data
}

// FindById implements repository.UserRepository.
func (u *UserRepositoryImpl) FindById(userId uuid.UUID) (users model.Users, err error) {
	var user model.Users
	result := u.Db.Find(&user, userId)
	if result != nil {
		return user, nil
	} else {
		return user, errors.New("not found")
	}
}

// Update implements repository.UserRepository.
func (u *UserRepositoryImpl) Update(user model.Users) {
	var updateUser = req.UserUpdateRequest{
		FullName: user.FullName,
		Phone:    user.Phone,
	}
	result := u.Db.Model(&user).Updates(updateUser)
	helper.ErrorPanic(result.Error)
}
