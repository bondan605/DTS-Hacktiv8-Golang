package routers

import (
	"middleware/controllers"
	"middleware/repository"
	"middleware/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupAuthRoute(router *gin.Engine, db *gorm.DB) {
	userRepository := repository.NewUserRepository(db)
	userService := services.NewUserService(userRepository)
	userController := controllers.NewUserController(*userService)

	userRouter := router.Group("/users")
	{
		userRouter.POST("/register", userController.Registration)
		userRouter.POST("/login", userController.Login)
	}
}
