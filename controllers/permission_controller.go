package controllers

import (
	"fmt"
	"net/http"
	"totesbackend/config"
	"totesbackend/controllers/utilities"
	"totesbackend/services"

	"github.com/gin-gonic/gin"
)

type PermissionController struct {
	Service *services.PermissionService
	Auth    *utilities.AuthorizationUtil
}

func NewPermissionController(service *services.PermissionService, auth *utilities.AuthorizationUtil) *PermissionController {
	return &PermissionController{Service: service, Auth: auth}
}

func (pc *PermissionController) GetPermissionByID(c *gin.Context) {
	permissionId := config.PERMISSION_GET_PERMISSION_BY_ID

	if !pc.Auth.CheckPermission(c, permissionId) {
		return
	}
	idParam := c.Param("id")
	var id uint
	if _, err := fmt.Sscanf(idParam, "%d", &id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid permission ID"})
		return
	}

	permission, err := pc.Service.GetPermissionByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Permission not found"})
		return
	}

	c.JSON(http.StatusOK, permission)
}

func (pc *PermissionController) GetAllPermissions(c *gin.Context) {
	permissionId := config.PERMISSION_GET_ALL_PERMISSIONS

	if !pc.Auth.CheckPermission(c, permissionId) {
		return
	}

	permissions, err := pc.Service.GetAllPermissions()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving permissions"})
		return
	}

	c.JSON(http.StatusOK, permissions)
}

func (pc *PermissionController) SearchPermissionsByID(c *gin.Context) {
	permissionId := config.PERMISSION_SEARCH_PERMISSION_BY_ID

	if !pc.Auth.CheckPermission(c, permissionId) {
		return
	}
	query := c.Query("id")
	if query == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Search query is required"})
		return
	}

	permissions, err := pc.Service.SearchPermissionsByID(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving permissions"})
		return
	}

	c.JSON(http.StatusOK, permissions)
}

func (pc *PermissionController) SearchPermissionsByName(c *gin.Context) {

	permissionId := config.PERMISSION_SEARCH_PERMISSION_BY_NAME

	if !pc.Auth.CheckPermission(c, permissionId) {
		return
	}

	query := c.Query("name")
	if query == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Search query is required"})
		return
	}

	permissions, err := pc.Service.SearchPermissionsByName(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving permissions"})
		return
	}

	c.JSON(http.StatusOK, permissions)
}
