package controllers

import (
	"net/http"
	"totesbackend/config"
	"totesbackend/controllers/utilities"
	"totesbackend/services"

	"github.com/gin-gonic/gin"
)

type UserLogController struct {
	Service *services.UserLogService
	Auth    *utilities.AuthorizationUtil
}

func NewUserLogController(service *services.UserLogService, auth *utilities.AuthorizationUtil) *UserLogController {
	return &UserLogController{Service: service, Auth: auth}
}

func (c *UserLogController) GetAllLogsFromUser(ctx *gin.Context) {
	permissionId := config.PERMISSION_GET_ALL_LOGS_FROM_USER

	if !c.Auth.CheckPermission(ctx, permissionId) {
		return
	}

	userID := ctx.Param("id")

	userLogs, err := c.Service.GetUserLogs(userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user logs"})
		return
	}

	ctx.JSON(http.StatusOK, userLogs)
}
