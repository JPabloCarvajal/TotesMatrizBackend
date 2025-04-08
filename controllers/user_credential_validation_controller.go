package controllers

import (
	"net/http"

	"totesbackend/controllers/utilities"
	"totesbackend/services"

	"github.com/gin-gonic/gin"
)

type UserCredentialValidationController struct {
	Service *services.UserCredentialValidationService
	Auth    *utilities.AuthorizationUtil
	Log     *utilities.LogUtil
}

func NewUserCredentialValidationController(service *services.UserCredentialValidationService, auth *utilities.AuthorizationUtil, log *utilities.LogUtil) *UserCredentialValidationController {
	return &UserCredentialValidationController{Service: service, Auth: auth, Log: log}
}

func (ucvc *UserCredentialValidationController) ValidateUserCredentials(c *gin.Context) {
	if ucvc.Log.RegisterLog(c, "Attempting user login") != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error registering log"})
		return
	}

	var loginData struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&loginData); err != nil {
		_ = ucvc.Log.RegisterLog(c, "Invalid request body for login")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	err := ucvc.Service.ValidateUserCredentials(loginData.Email, loginData.Password)
	if err != nil {
		if err.Error() == "user is not active" {
			_ = ucvc.Log.RegisterLog(c, "Login attempt for inactive user: "+loginData.Email)
			c.JSON(http.StatusForbidden, gin.H{"error": "User account is not active"})
			return
		}

		_ = ucvc.Log.RegisterLog(c, "Login failed for user: "+loginData.Email)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	_ = ucvc.Log.RegisterLog(c, "Login successful for user: "+loginData.Email)

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
	})

}
