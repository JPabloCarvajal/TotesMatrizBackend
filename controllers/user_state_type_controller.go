package controllers

import (
	"fmt"
	"net/http"
	"totesbackend/services"

	"github.com/gin-gonic/gin"
)

type UserStateTypeController struct {
	Service *services.UserStateTypeService
}

func NewUserStateTypeController(service *services.UserStateTypeService) *UserStateTypeController {
	return &UserStateTypeController{Service: service}
}

func (ustc *UserStateTypeController) GetUserStateTypes(c *gin.Context) {
	username := c.GetHeader("Username")
	fmt.Println("Request made by user:", username)

	userStateTypes, err := ustc.Service.GetAllUserStateTypes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving User State Types"})
		return
	}
	c.JSON(http.StatusOK, userStateTypes)

}

func (ustc *UserStateTypeController) GetUserStateTypeByID(c *gin.Context) {
	username := c.GetHeader("Username")
	fmt.Println("Request made by user:", username)

	id := c.Param("id")

	userStateType, err := ustc.Service.GetUserStateTypeByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User State Type not found"})
		return
	}

	c.JSON(http.StatusOK, userStateType)
}
