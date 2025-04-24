package controllers

import (
	"net/http"

	"totesbackend/config"
	"totesbackend/controllers/utilities"
	"totesbackend/models"
	"totesbackend/services"

	"github.com/gin-gonic/gin"
)

type TaxTypeController struct {
	Service *services.TaxTypeService
	Auth    *utilities.AuthorizationUtil
	Log     *utilities.LogUtil
}

func NewTaxTypeController(service *services.TaxTypeService,
	auth *utilities.AuthorizationUtil, log *utilities.LogUtil) *TaxTypeController {
	return &TaxTypeController{Service: service, Auth: auth, Log: log}
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

func (ttc *TaxTypeController) CreateTaxType(c *gin.Context) {
	if ttc.Log.RegisterLog(c, "Attempting to create a new tax type") != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error registering log"})
		return
	}

	permissionId := config.PERMISSION_CREATE_TAX_TYPE
	if !ttc.Auth.CheckPermission(c, permissionId) {
		_ = ttc.Log.RegisterLog(c, "Access denied for CreateTaxType")
		return
	}

	var tax models.TaxType
	if err := c.ShouldBindJSON(&tax); err != nil {
		_ = ttc.Log.RegisterLog(c, "Invalid input for tax type creation: "+err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos del impuesto"})
		return
	}

	err := ttc.Service.CreateTaxType(&tax)
	if err != nil {
		_ = ttc.Log.RegisterLog(c, "Failed to create tax type: "+err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo crear el impuesto"})
		return
	}

	_ = ttc.Log.RegisterLog(c, "Successfully created new tax type")
	c.JSON(http.StatusCreated, tax)
}
