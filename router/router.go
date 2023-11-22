package router

import (
	"main/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter(userController *controller.UserController, tagController *controller.TagController) *gin.Engine {
	service := gin.Default()

	service.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "welcome home")
	})
	service.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	router := service.Group("/api")
	tagRouter := router.Group("/tag")
	userRouter := router.Group("/user")
	// user
	userRouter.POST("", userController.Create)
	userRouter.GET("/:id", userController.FindById)
	userRouter.PATCH("/:id", userController.Update)

	// tag
	tagRouter.GET("", tagController.FindAll)
	tagRouter.GET("/:tagId", tagController.FindById)
	tagRouter.POST("", tagController.Create)
	tagRouter.PATCH("/:tagId", tagController.Update)
	tagRouter.DELETE("/:tagId", tagController.Delete)
	return service
}
