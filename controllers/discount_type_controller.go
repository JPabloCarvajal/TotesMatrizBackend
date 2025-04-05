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
	Log     *utilities.LogUtil //
}

func NewDiscountTypeController(service *services.DiscountTypeService, auth *utilities.AuthorizationUtil, log *utilities.LogUtil) *DiscountTypeController {
	return &DiscountTypeController{Service: service, Auth: auth, Log: log}
}

func (dtc *DiscountTypeController) GetDiscountTypeByID(c *gin.Context) {
	id := c.Param("id")

	if dtc.Log.RegisterLog(c, "Attempting to retrieve discount type with ID: "+id) != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error registering log"})
		return
	}

	permissionId := config.PERMISSION_GET_DISCOUNT_TYPE_BY_ID
	if !dtc.Auth.CheckPermission(c, permissionId) {
		_ = dtc.Log.RegisterLog(c, "Access denied for GetDiscountTypeByID")
		return
	}

	discountType, err := dtc.Service.GetDiscountTypeByID(id)
	if err != nil {
		_ = dtc.Log.RegisterLog(c, "Discount Type with ID "+id+" not found: "+err.Error())
		c.JSON(http.StatusNotFound, gin.H{"error": "Discount Type not found"})
		return
	}

	_ = dtc.Log.RegisterLog(c, "Successfully retrieved discount type with ID: "+id)
	c.JSON(http.StatusOK, discountType)
}

func (dtc *DiscountTypeController) GetAllDiscountTypes(c *gin.Context) {
	if dtc.Log.RegisterLog(c, "Attempting to retrieve all discount types") != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error registering log"})
		return
	}

	permissionId := config.PERMISSION_GET_ALL_DISCOUNT_TYPES
	if !dtc.Auth.CheckPermission(c, permissionId) {
		_ = dtc.Log.RegisterLog(c, "Access denied for GetAllDiscountTypes")
		return
	}

	discountTypes, err := dtc.Service.GetAllDiscountTypes()
	if err != nil {
		_ = dtc.Log.RegisterLog(c, "Error retrieving discount types: "+err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving Discount Types"})
		return
	}

	_ = dtc.Log.RegisterLog(c, "Successfully retrieved all discount types")
	c.JSON(http.StatusOK, discountTypes)
}
