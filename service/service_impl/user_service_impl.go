package serviceimpl

import (
	"main/data/req"
	"main/helper"
	"main/model"
	"main/repository"
	"main/service"
	"main/utils"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	Validate       *validator.Validate
}

// Create implements service.UserService.
func (u *UserServiceImpl) Create(user req.UserRequest) model.Users {
	id := uuid.New()
	hash := utils.HashPassword([]byte(user.Password))
	userModel := model.Users{
		Id:       id,
		FullName: user.FullName,
		Email:    user.Email,
		Password: hash,
		Phone:    user.Phone,
	}
	u.UserRepository.Create(userModel)

	return userModel
}

// FindById implements service.UserService.
func (u *UserServiceImpl) FindById(userId uuid.UUID) model.Users {
	userData, err := u.UserRepository.FindById(userId)
	helper.ErrorPanic(err)

	response := model.Users{
		Id:       userData.Id,
		FullName: userData.FullName,
		Email:    userData.Email,
		Password: userData.Password,
		Phone:    userData.Phone,
	}
	return response
}

// Update implements service.UserService.
func (u *UserServiceImpl) Update(user req.UserUpdateRequest) {
	userData, err := u.UserRepository.FindById(user.Id)
	helper.ErrorPanic(err)
	userData.FullName = user.FullName
	userData.Phone = user.Phone
	u.UserRepository.Update(userData)
}

func NewUserServiceImpl(userRepository repository.UserRepository,
	validate *validator.Validate) service.UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
		Validate:       validate,
	}
}
