package controllers

import (
	"net/http"

	"totesbackend/config"
	"totesbackend/controllers/utilities"
	"totesbackend/services"

	"github.com/gin-gonic/gin"
)

type TaxTypeController struct {
	Service *services.TaxTypeService
	Auth    *utilities.AuthorizationUtil
}

func NewTaxTypeController(service *services.TaxTypeService, auth *utilities.AuthorizationUtil) *TaxTypeController {
	return &TaxTypeController{Service: service, Auth: auth}
}

func (ttc *TaxTypeController) GetTaxTypeByID(c *gin.Context) {
	permissionId := config.PERMISSION_GET_TAX_TYPE_BY_ID

	if !ttc.Auth.CheckPermission(c, permissionId) {
		return
	}

	id := c.Param("id")
	taxType, err := ttc.Service.GetTaxTypeByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Tax Type not found"})
		return
	}

	c.JSON(http.StatusOK, taxType)
}

func (ttc *TaxTypeController) GetAllTaxTypes(c *gin.Context) {
	permissionId := config.PERMISSION_GET_ALL_TAX_TYPES

	if !ttc.Auth.CheckPermission(c, permissionId) {
		return
	}

	taxTypes, err := ttc.Service.GetAllTaxTypes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving Tax Types"})
		return
	}
	c.JSON(http.StatusOK, taxTypes)
}
