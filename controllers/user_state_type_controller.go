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
}

func NewUserStateTypeController(service *services.UserStateTypeService, auth *utilities.AuthorizationUtil) *UserStateTypeController {
	return &UserStateTypeController{Service: service, Auth: auth}
}

func (ustc *UserStateTypeController) GetUserStateTypeByID(c *gin.Context) {
	permissionId := config.PERMISSION_GET_USER_STATE_TYPE_BY_ID

	if !ustc.Auth.CheckPermission(c, permissionId) {
		return
	}

	id := c.Param("id")

	userStateType, err := ustc.Service.GetUserStateTypeByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User State Type not found"})
		return
	}

	c.JSON(http.StatusOK, userStateType)
}

func (ustc *UserStateTypeController) GetAllUserStateTypes(c *gin.Context) {
	permissionId := config.PERMISSION_GET_ALL_USER_STATE_TYPES

	if !ustc.Auth.CheckPermission(c, permissionId) {
		return
	}

	userStateTypes, err := ustc.Service.GetAllUserStateTypes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving User State Types"})
		return
	}
	c.JSON(http.StatusOK, userStateTypes)

}
