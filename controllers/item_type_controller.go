package controllers

import (
	"fmt"
	"net/http"

	"totesbackend/services"

	"github.com/gin-gonic/gin"
)

type ItemTypeController struct {
	Service *services.ItemTypeService
}

func NewItemTypeController(service *services.ItemTypeService) *ItemTypeController {
	return &ItemTypeController{Service: service}
}

func (itc *ItemTypeController) GetItemTypes(c *gin.Context) {
	username := c.GetHeader("Username")
	fmt.Println("Request made by user:", username)

	itemTypes, err := itc.Service.GetAllItemTypes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving Item Types"})
		return
	}
	c.JSON(http.StatusOK, itemTypes)
}

func (itc *ItemTypeController) GetItemTypeByID(c *gin.Context) {
	username := c.GetHeader("Username")
	fmt.Println("Request made by user:", username)

	id := c.Param("id")

	itemType, err := itc.Service.GetItemTypeByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item Type not found"})
		return
	}

	c.JSON(http.StatusOK, itemType)
}
