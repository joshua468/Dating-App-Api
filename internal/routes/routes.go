package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/joshua468/Dating-App-Api/controllers"
	"github.com/joshua468/Dating-App-Api/internal/middleware"
)

func SetupRoutes(router *gin.Engine) {
	userRoutes := router.Group("/api/users")
	{
		userRoutes.POST("/register", controllers.RegisterUser)
		userRoutes.POST("/login", controllers.LoginUser)
	}

	matchRoutes := router.Group("/api/matches")
	matchRoutes.Use(middleware.AuthMiddleware())
	{
		matchRoutes.POST("/", controllers.CreateMatch)
		matchRoutes.GET("/", controllers.GetMatches)
	}
}
