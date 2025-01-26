package routes

import (
	"task-tracker/api/controllers"
	"task-tracker/api/usecases"

	"github.com/gin-gonic/gin"
)

// RegisterAuthRoutes registers the auth routes
func RegisterAuthRoutes(router *gin.Engine, authUseCases usecases.AuthUseCase) {
	authController := controllers.NewAuthController(authUseCases)

	userRoutes := router.Group("/auth")
	{
		userRoutes.POST("/signup", authController.SignUp)
		userRoutes.POST("/signin", authController.SignIn)

	}
}
