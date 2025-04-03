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
	Log     *utilities.LogUtil
}

func NewItemTypeController(service *services.ItemTypeService, auth *utilities.AuthorizationUtil, log *utilities.LogUtil) *ItemTypeController {
	return &ItemTypeController{Service: service, Auth: auth, Log: log}
}
func (itc *ItemTypeController) GetItemTypeByID(c *gin.Context) {
	id := c.Param("id")

	if itc.Log.RegisterLog(c, "Attempting to retrieve ItemType with ID: "+id) != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error registering log"})
		return
	}

	permissionId := config.PERMISSION_GET_ITEM_BY_ID
	if !itc.Auth.CheckPermission(c, permissionId) {
		_ = itc.Log.RegisterLog(c, "Access denied for GetItemTypeByID")
		return
	}

	itemType, err := itc.Service.GetItemTypeByID(id)
	if err != nil {
		_ = itc.Log.RegisterLog(c, "Error retrieving ItemType with ID "+id+": "+err.Error())
		c.JSON(http.StatusNotFound, gin.H{"error": "Item Type not found"})
		return
	}

	_ = itc.Log.RegisterLog(c, "Successfully retrieved ItemType with ID: "+id)
	c.JSON(http.StatusOK, itemType)
}

func (itc *ItemTypeController) GetItemTypes(c *gin.Context) {
	if itc.Log.RegisterLog(c, "Attempting to retrieve all ItemTypes") != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error registering log"})
		return
	}

	permissionId := config.PERMISSION_GET_ITEM_TYPES
	if !itc.Auth.CheckPermission(c, permissionId) {
		_ = itc.Log.RegisterLog(c, "Access denied for GetItemTypes")
		return
	}

	itemTypes, err := itc.Service.GetAllItemTypes()
	if err != nil {
		_ = itc.Log.RegisterLog(c, "Error retrieving ItemTypes: "+err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving Item Types"})
		return
	}

	_ = itc.Log.RegisterLog(c, "Successfully retrieved all ItemTypes")
	c.JSON(http.StatusOK, itemTypes)
}
