package routes

import (
	"github.com/Skapar/simple-rest/internal/api"
	"github.com/Skapar/simple-rest/internal/service"
	"github.com/gin-gonic/gin"
)

// SetupRoutes - настройка маршрутов
func SetupRoutes(router *gin.Engine, authService service.AuthService) {
	authHandler := api.NewAuthHandler(authService)
	router.POST("/signup", authHandler.RegisterUserHandler)
}
