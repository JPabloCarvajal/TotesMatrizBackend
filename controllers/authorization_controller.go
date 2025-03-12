package controllers

import (
	"net/http"

	"totesbackend/services"

	"github.com/gin-gonic/gin"
)

type AuthorizationController struct {
	Service *services.AuthorizationService
}

func NewAuthorizationController(service *services.AuthorizationService) *AuthorizationController {
	return &AuthorizationController{Service: service}
}

func (ac *AuthorizationController) CheckUserPermission(c *gin.Context) {
	email := c.Query("email")
	permissionID := c.Query("permission_id")

	if email == "" || permissionID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email and permission_id are required"})
		return
	}

	hasPermission, err := ac.Service.UserHasPermission(email, permissionID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error checking permission"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"has_permission": hasPermission})
}
