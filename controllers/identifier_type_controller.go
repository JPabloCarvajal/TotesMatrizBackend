package controllers

import (
	"net/http"
	"totesbackend/config"
	"totesbackend/controllers/utilities"
	"totesbackend/services"

	"github.com/gin-gonic/gin"
)

type IdentifierTypeController struct {
	Service *services.IdentifierTypeService
	Auth    *utilities.AuthorizationUtil
}

func NewIdentifierTypeController(service *services.IdentifierTypeService, auth *utilities.AuthorizationUtil) *IdentifierTypeController {
	return &IdentifierTypeController{Service: service, Auth: auth}
}

func (itc *IdentifierTypeController) GetAllIdentifierTypes(c *gin.Context) {
	permissionId := config.PERMISSION_GET_ALL_IDENTIFIER_TYPES

	if !itc.Auth.CheckPermission(c, permissionId) {
		return
	}

	identifierTypes, err := itc.Service.GetAllIdentifierTypes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving Identifier Types"})
		return
	}
	c.JSON(http.StatusOK, identifierTypes)
}

func (itc *IdentifierTypeController) GetIdentifierTypeByID(c *gin.Context) {
	permissionId := config.PERMISSION_GET_IDENTIFIER_TYPE_BY_ID

	if !itc.Auth.CheckPermission(c, permissionId) {
		return
	}

	id := c.Param("id")

	identifierType, err := itc.Service.GetIdentifierTypeByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Identifier Type not found"})
		return
	}

	c.JSON(http.StatusOK, identifierType)
}
