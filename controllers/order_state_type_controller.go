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
}

func NewOrderStateTypeController(service *services.OrderStateTypeService, auth *utilities.AuthorizationUtil) *OrderStateTypeController {
	return &OrderStateTypeController{Service: service, Auth: auth}
}

func (ostc *OrderStateTypeController) GetOrderStateTypeByID(c *gin.Context) {
	permissionId := config.PERMISSION_GET_ORDER_STATE_TYPE_BY_ID

	if !ostc.Auth.CheckPermission(c, permissionId) {
		return
	}

	id := c.Param("id")

	orderStateType, err := ostc.Service.GetOrderStateTypeByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order State Type not found"})
		return
	}

	c.JSON(http.StatusOK, orderStateType)
}

func (ostc *OrderStateTypeController) GetAllOrderStateTypes(c *gin.Context) {
	permissionId := config.PERMISSION_GET_ALL_ORDER_STATE_TYPES

	if !ostc.Auth.CheckPermission(c, permissionId) {
		return
	}

	orderStateTypes, err := ostc.Service.GetAllOrderStateTypes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving Order State Types"})
		return
	}
	c.JSON(http.StatusOK, orderStateTypes)

}
