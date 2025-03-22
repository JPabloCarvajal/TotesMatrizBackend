package controllers

import (
	"net/http"

	"totesbackend/config"
	"totesbackend/controllers/utilities"
	"totesbackend/services"

	"github.com/gin-gonic/gin"
)

type DiscountTypeController struct {
	Service *services.DiscountTypeService
	Auth    *utilities.AuthorizationUtil
}

func NewDiscountTypeController(service *services.DiscountTypeService, auth *utilities.AuthorizationUtil) *DiscountTypeController {
	return &DiscountTypeController{Service: service, Auth: auth}
}

func (dtc *DiscountTypeController) GetDiscountTypeByID(c *gin.Context) {
	permissionId := config.PERMISSION_GET_DISCOUNT_TYPE_BY_ID

	if !dtc.Auth.CheckPermission(c, permissionId) {
		return
	}

	id := c.Param("id")
	discountType, err := dtc.Service.GetDiscountTypeByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Discount Type not found"})
		return
	}

	c.JSON(http.StatusOK, discountType)
}

func (dtc *DiscountTypeController) GetAllDiscountTypes(c *gin.Context) {
	permissionId := config.PERMISSION_GET_ALL_DISCOUNT_TYPES

	if !dtc.Auth.CheckPermission(c, permissionId) {
		return
	}

	discountTypes, err := dtc.Service.GetAllDiscountTypes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving Discount Types"})
		return
	}
	c.JSON(http.StatusOK, discountTypes)
}
