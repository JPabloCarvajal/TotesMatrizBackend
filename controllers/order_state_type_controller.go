package controllers

import (
	"net/http"
	"totesbackend/config"
	"totesbackend/controllers/utilities"
	"totesbackend/services"

	"github.com/gin-gonic/gin"
)

type OrderStateTypeController struct {
	Service *services.OrderStateTypeService
	Auth    *utilities.AuthorizationUtil
	Log     *utilities.LogUtil
}

func NewOrderStateTypeController(service *services.OrderStateTypeService, auth *utilities.AuthorizationUtil, log *utilities.LogUtil) *OrderStateTypeController {
	return &OrderStateTypeController{Service: service, Auth: auth, Log: log}
}

func (ostc *OrderStateTypeController) GetOrderStateTypeByID(c *gin.Context) {
	permissionId := config.PERMISSION_GET_ORDER_STATE_TYPE_BY_ID

	if ostc.Log.RegisterLog(c, "Attempting to get order state type by ID") != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error registering log"})
		return
	}

	if !ostc.Auth.CheckPermission(c, permissionId) {
		_ = ostc.Log.RegisterLog(c, "Access denied for GetOrderStateTypeByID")
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		return
	}

	id := c.Param("id")

	orderStateType, err := ostc.Service.GetOrderStateTypeByID(id)
	if err != nil {
		_ = ostc.Log.RegisterLog(c, "Order state type not found with ID: "+id)
		c.JSON(http.StatusNotFound, gin.H{"error": "Order State Type not found"})
		return
	}

	_ = ostc.Log.RegisterLog(c, "Successfully retrieved order state type with ID: "+id)
	c.JSON(http.StatusOK, orderStateType)
}

func (ostc *OrderStateTypeController) GetAllOrderStateTypes(c *gin.Context) {
	permissionId := config.PERMISSION_GET_ALL_ORDER_STATE_TYPES

	if ostc.Log.RegisterLog(c, "Attempting to get all order state types") != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error registering log"})
		return
	}

	if !ostc.Auth.CheckPermission(c, permissionId) {
		_ = ostc.Log.RegisterLog(c, "Access denied for GetAllOrderStateTypes")
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		return
	}

	orderStateTypes, err := ostc.Service.GetAllOrderStateTypes()
	if err != nil {
		_ = ostc.Log.RegisterLog(c, "Error retrieving order state types: "+err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving Order State Types"})
		return
	}

	_ = ostc.Log.RegisterLog(c, "Successfully retrieved all order state types")
	c.JSON(http.StatusOK, orderStateTypes)
}
