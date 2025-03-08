package controllers

import (
	"fmt"
	"net/http"
	"totesbackend/services"

	"github.com/gin-gonic/gin"
)

type PermissionController struct {
	Service *services.PermissionService
}

func NewPermissionController(service *services.PermissionService) *PermissionController {
	return &PermissionController{Service: service}
}

func (pc *PermissionController) GetPermissionByID(c *gin.Context) {
	username := c.GetHeader("Username")
	fmt.Println("Request made by user:", username)

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
	username := c.GetHeader("Username")
	fmt.Println("Request made by user:", username)

	permissions, err := pc.Service.GetAllPermissions()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving permissions"})
		return
	}

	c.JSON(http.StatusOK, permissions)
}
