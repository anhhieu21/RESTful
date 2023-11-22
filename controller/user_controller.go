package controller

import (
	"main/data/req"
	"main/data/res"
	"main/helper"
	"main/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (u *UserController) Create(ctx *gin.Context) {
	createUserRequets := req.UserRequest{}
	err := ctx.ShouldBindJSON(&createUserRequets)
	helper.ErrorPanic(err)
	u.userService.Create(createUserRequets)
	webResponse := res.Response{
		Code:   200,
		Status: "Ok",
		Data:   nil,
	}
	ctx.JSON(http.StatusOK, webResponse)
}

func (u *UserController) Update(ctx *gin.Context) {
	updateRequest := req.UserUpdateRequest{}
	err := ctx.ShouldBindJSON(&updateRequest)
	helper.ErrorPanic(err)

	userId := ctx.Param("id")
	id := uuid.MustParse(userId)
	updateRequest.Id = id
	u.userService.Update(updateRequest)
	webResponse := res.Response{
		Code:   200,
		Status: "Ok",
		Data:   nil,
	}
	ctx.JSON(http.StatusOK, webResponse)

}
func (u *UserController) FindById(ctx *gin.Context) {
	userId := ctx.Param("id")
	id := uuid.MustParse(userId)
	userReponse := u.userService.FindById(id)
	webResponse := res.Response{
		Code:   200,
		Status: "Ok",
		Data:   userReponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}
