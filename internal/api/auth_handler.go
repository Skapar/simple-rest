package api

import (
	"net/http"

	"github.com/Skapar/simple-rest/internal/models"
	"github.com/Skapar/simple-rest/internal/models/entities"
	"github.com/Skapar/simple-rest/internal/service"
	"github.com/Skapar/simple-rest/internal/utils"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService service.AuthService
}

func NewAuthHandler(authService service.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

func (ah *AuthHandler) RegisterUserHandler(c *gin.Context) {
	var user entities.User
	if err := c.ShouldBindJSON(&user); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}

	accessToken, refreshToken, err := ah.authService.RegisterUser(&user)
	if err != nil {
		handleRegisterUserError(c, err)
		return
	}

	c.JSON(http.StatusCreated, models.Response[interface{}]{
		Data:    gin.H{"user": user, "access_token": accessToken, "refresh_token": refreshToken},
		Message: "user registered successfully",
		Success: true,
	})
}

func handleRegisterUserError(c *gin.Context, err error) {
	switch err.Error() {
	case "user already exists":
		utils.RespondWithError(c, http.StatusConflict, err.Error())
	default:
		utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
	}
}
