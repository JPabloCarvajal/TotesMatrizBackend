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
}

func NewUserCredentialValidationController(service *services.UserCredentialValidationService, auth *utilities.AuthorizationUtil) *UserCredentialValidationController {
	return &UserCredentialValidationController{Service: service, Auth: auth}
}

func (ucvc *UserCredentialValidationController) ValidateUserCredentials(c *gin.Context) {

	var loginRequest struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	isValid, err := ucvc.Service.ValidateUserCredentials(loginRequest.Email, loginRequest.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	if !isValid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
}
