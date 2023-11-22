package controller

import (
	"main/data/req"
	"main/data/res"
	"main/helper"
	"main/model"
	"main/service"
	"main/utils"
	"net/http"
	"strings"
	"github.com/dgrijalva/jwt-go"

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

func (u *UserController) FindMe(ctx *gin.Context) {
	authorizationHeader := ctx.GetHeader("Authorization")
	if authorizationHeader == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		ctx.Abort()
		return
	}
	tokenString := strings.Split(authorizationHeader, " ")[1]
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(utils.SECRET_KEY), nil
	}).(*jwt.Token)
	if err != nil || !token.Valid {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		ctx.Abort()
		return
	}
	claims := token.Claims.(*model.JwtCustomClaims)
}

func (u *UserController) SignUp(ctx *gin.Context) {
	createUserRequets := req.UserRequest{}
	err := ctx.ShouldBindJSON(&createUserRequets)
	helper.ErrorPanic(err)
	user := u.userService.Create(createUserRequets)
	//GenToken
	token, err := utils.GenToken(user)
	helper.ErrorPanic(err)

	ctx.JSON(http.StatusOK, gin.H{"token": token})
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
