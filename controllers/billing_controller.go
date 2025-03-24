package controllers

import (
	"net/http"
	"strconv"
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
	permissionId := config.PERMISSION_CALCULATE_SUBTOTAL

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

func (bc *BillingController) CalculateTotal(c *gin.Context) {
	permissionId := config.PERMISSION_CALCULATE_TOTAL

	if !bc.Auth.CheckPermission(c, permissionId) {
		return
	}

	// Estructura del request con arrays de enteros
	var request struct {
		DiscountTypesIds []int                 `json:"discountTypesIds"`
		TaxTypesIds      []int                 `json:"taxTypesIds"`
		ItemsDTO         []dtos.BillingItemDTO `json:"itemsDTO"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	discountTypesIdsStr := make([]string, len(request.DiscountTypesIds))
	for i, id := range request.DiscountTypesIds {
		discountTypesIdsStr[i] = strconv.Itoa(id)
	}

	taxTypesIdsStr := make([]string, len(request.TaxTypesIds))
	for i, id := range request.TaxTypesIds {
		taxTypesIdsStr[i] = strconv.Itoa(id)
	}

	total, err := bc.Service.CalculateTotal(discountTypesIdsStr, taxTypesIdsStr, request.ItemsDTO)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"total": total})
}
