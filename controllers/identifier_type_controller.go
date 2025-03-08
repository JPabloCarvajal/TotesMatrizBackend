package controllers

import (
	"fmt"
	"net/http"
	"totesbackend/services"

	"github.com/gin-gonic/gin"
)

type IdentifierTypeController struct {
	Service *services.IdentifierTypeService
}

func NewIdentifierTypeController(service *services.IdentifierTypeService) *IdentifierTypeController {
	return &IdentifierTypeController{Service: service}
}

func (itc *IdentifierTypeController) GetIdentifierTypes(c *gin.Context) {
	username := c.GetHeader("Username")
	fmt.Println("Request made by user:", username)

	identifierTypes, err := itc.Service.GetAllIdentifierTypes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving Identifier Types"})
		return
	}
	c.JSON(http.StatusOK, identifierTypes)
}

func (itc *IdentifierTypeController) GetIdentifierTypeByID(c *gin.Context) {
	username := c.GetHeader("Username")
	fmt.Println("Request made by user:", username)

	id := c.Param("id")

	identifierType, err := itc.Service.GetIdentifierTypeByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Identifier Type not found"})
		return
	}

	c.JSON(http.StatusOK, identifierType)
}
