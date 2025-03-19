package utilities

import (
	"net/http"
	"totesbackend/services"

	"github.com/gin-gonic/gin"
)

type AuthorizationUtil struct {
	Service *services.AuthorizationService
}

func NewAuthorizationUtil(service *services.AuthorizationService) *AuthorizationUtil {
	return &AuthorizationUtil{Service: service}
}

func (u *AuthorizationUtil) CheckPermission(c *gin.Context, permissionID int) bool {
	username := c.GetHeader("Username")
	authResult, err := u.Service.UserHasPermission(username, permissionID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Authorization service error"})
		return false
	}

	if !authResult {
		c.JSON(http.StatusForbidden, gin.H{"error": "User does not have permission"})
		return false
	}

	return true
}
