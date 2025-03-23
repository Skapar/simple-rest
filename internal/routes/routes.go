package routes

import (
	"github.com/Skapar/simple-rest/internal/api"
	"github.com/Skapar/simple-rest/internal/service"
	"github.com/gin-gonic/gin"
)

// SetupRoutes - настройка маршрутов
func SetupRoutes(router *gin.Engine, authService service.AuthService, userService service.UserService) {
	authHandler := api.NewAuthHandler(authService)
	userHandler := api.NewUserHandler(userService)

	router.POST("/signup", authHandler.RegisterUserHandler)
	router.GET("/user/:user_id", userHandler.GetUserHandler)
	router.PUT("/user/:user_id", userHandler.UpdateUserHandler)
	router.DELETE("/user/:user_id", userHandler.DeleteUserHandler)
	router.PATCH("/user/:user_id", userHandler.SoftDeleteUserHandler)
}
