package api

import (
	"github.com/Skapar/simple-rest/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type HealthHandler struct{}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (hh *HealthHandler) CheckHealthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, models.Response[interface{}]{
		Data:    nil,
		Message: "server is running",
		Success: true,
	})
}
