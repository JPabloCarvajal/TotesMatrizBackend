package controllers

import (
	"fmt"
	"net/http"
	"totesbackend/dtos"
	"totesbackend/services"

	"github.com/gin-gonic/gin"
)

type UserTypeController struct {
	Service *services.UserTypeService
}

func NewUserTypeController(service *services.UserTypeService) *UserTypeController {
	return &UserTypeController{Service: service}
}

func (utc *UserTypeController) GetUserTypeByID(c *gin.Context) {
	username := c.GetHeader("Username")
	fmt.Println("Request made by user:", username)

	idParam := c.Param("id")
	var id uint
	if _, err := fmt.Sscanf(idParam, "%d", &id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user type ID"})
		return
	}

	userType, err := utc.Service.GetUserTypeByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User type not found"})
		return
	}

	roleIDs, err := utc.Service.GetRolesForUserType(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving roles for user type"})
		return
	}

	userTypeDTO := dtos.UserTypeDTO{
		ID:          userType.ID,
		Name:        userType.Name,
		Description: userType.Description,
		Roles:       make([]string, len(roleIDs)),
	}

	for i, roleID := range roleIDs {
		userTypeDTO.Roles[i] = fmt.Sprintf("%d", roleID)
	}

	c.JSON(http.StatusOK, userTypeDTO)
}

func (utc *UserTypeController) GetAllUserTypes(c *gin.Context) {
	username := c.GetHeader("Username")
	fmt.Println("Request made by user:", username)

	userTypes, err := utc.Service.ObtainAllUserTypes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving user types"})
		return
	}

	var userTypesDTO []dtos.UserTypeDTO
	for _, userType := range userTypes {
		roleIDs, err := utc.Service.GetRolesForUserType(userType.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving roles for user type"})
			return
		}

		userTypeDTO := dtos.UserTypeDTO{
			ID:          userType.ID,
			Name:        userType.Name,
			Description: userType.Description,
			Roles:       make([]string, len(roleIDs)),
		}

		for i, roleID := range roleIDs {
			userTypeDTO.Roles[i] = fmt.Sprintf("%d", roleID)
		}

		userTypesDTO = append(userTypesDTO, userTypeDTO)
	}

	c.JSON(http.StatusOK, userTypesDTO)
}

func (utc *UserTypeController) ExistsUserType(c *gin.Context) {
	username := c.GetHeader("Username")
	fmt.Println("Request made by user:", username)

	idParam := c.Param("id")
	var id uint
	if _, err := fmt.Sscanf(idParam, "%d", &id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user type ID"})
		return
	}

	exists, err := utc.Service.Exists(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error checking user type existence"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"exists": exists})
}

func (utc *UserTypeController) SearchUserTypesByID(c *gin.Context) {
	username := c.GetHeader("Username")
	fmt.Println("Request made by user:", username)

	query := c.Query("id")
	userTypes, err := utc.Service.SearchUserTypesByID(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving user types"})
		return
	}

	var userTypesDTO []dtos.UserTypeDTO
	for _, userType := range userTypes {
		roleIDs, _ := utc.Service.GetRolesForUserType(userType.ID)

		userTypeDTO := dtos.UserTypeDTO{
			ID:          userType.ID,
			Name:        userType.Name,
			Description: userType.Description,
			Roles:       make([]string, len(roleIDs)),
		}

		for i, roleID := range roleIDs {
			userTypeDTO.Roles[i] = fmt.Sprintf("%d", roleID)
		}

		userTypesDTO = append(userTypesDTO, userTypeDTO)
	}

	c.JSON(http.StatusOK, userTypesDTO)
}

func (utc *UserTypeController) SearchUserTypesByName(c *gin.Context) {
	username := c.GetHeader("Username")
	fmt.Println("Request made by user:", username)

	query := c.Query("name")
	userTypes, err := utc.Service.SearchUserTypesByName(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving user types"})
		return
	}

	var userTypesDTO []dtos.UserTypeDTO
	for _, userType := range userTypes {
		roleIDs, _ := utc.Service.GetRolesForUserType(userType.ID)

		userTypeDTO := dtos.UserTypeDTO{
			ID:          userType.ID,
			Name:        userType.Name,
			Description: userType.Description,
			Roles:       make([]string, len(roleIDs)),
		}

		for i, roleID := range roleIDs {
			userTypeDTO.Roles[i] = fmt.Sprintf("%d", roleID)
		}

		userTypesDTO = append(userTypesDTO, userTypeDTO)
	}

	c.JSON(http.StatusOK, userTypesDTO)
}
