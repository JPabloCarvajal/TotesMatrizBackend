package controllers

import (
	"net/http"
	"totesbackend/config"
	"totesbackend/controllers/utilities"
	"totesbackend/services"

	"github.com/gin-gonic/gin"
)

type UserStateTypeController struct {
	Service *services.UserStateTypeService
	Auth    *utilities.AuthorizationUtil
	Log     *utilities.LogUtil
}

func NewUserStateTypeController(service *services.UserStateTypeService, auth *utilities.AuthorizationUtil, log *utilities.LogUtil) *UserStateTypeController {
	return &UserStateTypeController{Service: service, Auth: auth, Log: log}
}

func (ustc *UserStateTypeController) GetUserStateTypeByID(c *gin.Context) {
	permissionId := config.PERMISSION_GET_USER_STATE_TYPE_BY_ID

	if !ustc.Auth.CheckPermission(c, permissionId) {
		_ = ustc.Log.RegisterLog(c, "Access denied for GetUserStateTypeByID")
		return
	}

	id := c.Param("id")

	if ustc.Log.RegisterLog(c, "Attempting to retrieve user state type with ID: "+id) != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error registering log"})
		return
	}

	userStateType, err := ustc.Service.GetUserStateTypeByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User State Type not found"})
		return
	}

	_ = ustc.Log.RegisterLog(c, "Successfully retrieved user state type with ID: "+id)

	c.JSON(http.StatusOK, userStateType)
}

func (ustc *UserStateTypeController) GetAllUserStateTypes(c *gin.Context) {
	permissionId := config.PERMISSION_GET_ALL_USER_STATE_TYPES

	if !ustc.Auth.CheckPermission(c, permissionId) {
		_ = ustc.Log.RegisterLog(c, "Access denied for GetAllUserStateTypes")
		return
	}

	if ustc.Log.RegisterLog(c, "Attempting to retrieve all user state types") != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error registering log"})
		return
	}

	userStateTypes, err := ustc.Service.GetAllUserStateTypes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving User State Types"})
		return
	}

	_ = ustc.Log.RegisterLog(c, "Successfully retrieved all user state types")

	c.JSON(http.StatusOK, userStateTypes)
}
