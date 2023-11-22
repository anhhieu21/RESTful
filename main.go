package main

import (
	"main/config"
	"main/controller"
	"main/helper"
	"main/model"
	repoimpl "main/repository/repo_impl"
	"main/router"
	serviceimpl "main/service/service_impl"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
)

func main() {

	//database
	db := config.DatabaseConnection()
	validate := validator.New()

	db.Table("tags").AutoMigrate(&model.Tags{})
	db.Table("users").AutoMigrate(&model.Users{})
	//Init Repository
	userRepository := repoimpl.NewUserRepositoryImpl(db)
	tagRepository := repoimpl.NewTagsRepositoryImpl(db)

	//Init Service
	userService := serviceimpl.NewUserServiceImpl(userRepository, validate)
	tagService := serviceimpl.NewTagServiceImpl(tagRepository, validate)

	//Init controller
	userController := controller.NewUserController(userService)
	tagController := controller.NewTagController(tagService)

	//Router
	routes := router.NewRouter(userController, tagController)

	server := &http.Server{
		Addr:           ":8000",
		Handler:        routes,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	err := server.ListenAndServe()
	helper.ErrorPanic(err)
}
