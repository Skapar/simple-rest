package api

import (
	"net/http"
	"strconv"

	"github.com/Skapar/simple-rest/internal/models"
	"github.com/Skapar/simple-rest/internal/models/entities"
	"github.com/Skapar/simple-rest/internal/service"
	"github.com/Skapar/simple-rest/internal/utils"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (uh *UserHandler) UpdateUserHandler(c *gin.Context) {
	userID, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "invalid user_id")
		return
	}

	var user entities.User
	if err := c.ShouldBindJSON(&user); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := uh.userService.UpdateUser(userID, &user); err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, models.Response[entities.User]{
		Data:    user,
		Message: "user updated successfully",
		Success: true,
	})
}

func (uh *UserHandler) DeleteUserHandler(c *gin.Context) {
	userID, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "invalid user_id")
		return
	}

	if err := uh.userService.DeleteUser(userID); err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, models.Response[interface{}]{
		Data:    nil,
		Message: "user deleted successfully",
		Success: true,
	})
}

func (uh *UserHandler) SoftDeleteUserHandler(c *gin.Context) {
	userID, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "invalid user_id")
		return
	}

	if err := uh.userService.SoftDeleteUser(userID); err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, models.Response[interface{}]{
		Data:    nil,
		Message: "user soft deleted successfully",
		Success: true,
	})
}
