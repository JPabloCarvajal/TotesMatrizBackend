package controllers

import (
	"net/http"
	"strconv"

	"totesbackend/controllers/utilities"
	"totesbackend/services"

	"github.com/gin-gonic/gin"
)

type AuthorizationController struct {
	Service *services.AuthorizationService
	Log     *utilities.LogUtil
}

func NewAuthorizationController(service *services.AuthorizationService, log *utilities.LogUtil) *AuthorizationController {
	return &AuthorizationController{Service: service, Log: log}
}

func (ac *AuthorizationController) CheckUserPermission(c *gin.Context) {
	email := c.Query("email")
	permissionID := c.Query("permission_id")
	permissionStr, err := strconv.Atoi(permissionID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Permission ID"})
		return
	}

	if email == "" || permissionID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email and permission_id are required"})
		return
	}

	hasPermission, err := ac.Service.UserHasPermission(email, permissionStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error checking permission"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"has_permission": hasPermission})
}

func (ac *AuthorizationController) Login(c *gin.Context) {
	if ac.Log.RegisterLog(c, "Attempting user login") != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error registering log"})
		return
	}

	var loginData struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&loginData); err != nil {
		_ = ac.Log.RegisterLog(c, "Invalid request body for login")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	err := ac.Service.LoginUser(loginData.Email, loginData.Password)
	if err != nil {
		if err.Error() == "user is not active" {
			_ = ac.Log.RegisterLog(c, "Login attempt for inactive user: "+loginData.Email)
			c.JSON(http.StatusForbidden, gin.H{"error": "User account is not active"})
			return
		}

		_ = ac.Log.RegisterLog(c, "Login failed for user: "+loginData.Email)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	_ = ac.Log.RegisterLog(c, "Login successful for user: "+loginData.Email)

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
	})
}
