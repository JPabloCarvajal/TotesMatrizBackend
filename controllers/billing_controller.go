package controllers

import (
	"net/http"
	"totesbackend/config"
	"totesbackend/controllers/utilities"
	"totesbackend/dtos"
	"totesbackend/services"

	"github.com/gin-gonic/gin"
)

type BillingController struct {
	Service *services.BillingService
	Auth    *utilities.AuthorizationUtil
}

func NewBillingController(service *services.BillingService, auth *utilities.AuthorizationUtil) *BillingController {
	return &BillingController{Service: service, Auth: auth}
}

func (bc *BillingController) CalculateSubtotal(c *gin.Context) {
	permissionId := config.PERMISSION_CREATE_EMPLOYEE

	if !bc.Auth.CheckPermission(c, permissionId) {
		return
	}

	var itemsDTO []dtos.BillingItemDTO
	if err := c.ShouldBindJSON(&itemsDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	subtotal, err := bc.Service.CalculateSubtotal(itemsDTO)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"subtotal": subtotal})
}
