package controllers

import (
	"net/http"
	"totesbackend/services"

	"github.com/gin-gonic/gin"
)

type UserLogController struct {
	Service *services.UserLogService
}

func NewUserLogController(service *services.UserLogService) *UserLogController {
	return &UserLogController{Service: service}
}

func (c *UserLogController) GetUserLogs(ctx *gin.Context) {
	userID := ctx.Param("id")

	userLogs, err := c.Service.GetUserLogs(userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user logs"})
		return
	}

	ctx.JSON(http.StatusOK, userLogs)
}
