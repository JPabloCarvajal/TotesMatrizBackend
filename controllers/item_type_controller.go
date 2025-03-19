package controllers

import (
	"net/http"

	"totesbackend/config"
	"totesbackend/controllers/utilities"
	"totesbackend/services"

	"github.com/gin-gonic/gin"
)

type ItemTypeController struct {
	Service *services.ItemTypeService
	Auth    *utilities.AuthorizationUtil
}

func NewItemTypeController(service *services.ItemTypeService, auth *utilities.AuthorizationUtil) *ItemTypeController {
	return &ItemTypeController{Service: service, Auth: auth}
}

func (itc *ItemTypeController) GetItemTypeByID(c *gin.Context) {

	permissionId := config.PERMISSION_GET_ITEM_BY_ID

	if !itc.Auth.CheckPermission(c, permissionId) {
		return
	}

	id := c.Param("id")
	itemType, err := itc.Service.GetItemTypeByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item Type not found"})
		return
	}

	c.JSON(http.StatusOK, itemType)
}

func (itc *ItemTypeController) GetItemTypes(c *gin.Context) {
	permissionId := config.PERMISSION_GET_ITEM_TYPES

	if !itc.Auth.CheckPermission(c, permissionId) {
		return
	}

	itemTypes, err := itc.Service.GetAllItemTypes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving Item Types"})
		return
	}
	c.JSON(http.StatusOK, itemTypes)
}
