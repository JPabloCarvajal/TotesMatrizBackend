package controllers

import (
	"fmt"
	"net/http"
	"totesbackend/config"
	"totesbackend/controllers/utilities"
	"totesbackend/dtos"
	"totesbackend/services"

	"github.com/gin-gonic/gin"
)

type UserTypeController struct {
	Service *services.UserTypeService
	Auth    *utilities.AuthorizationUtil
	Log     *utilities.LogUtil
}

func NewUserTypeController(service *services.UserTypeService, auth *utilities.AuthorizationUtil, log *utilities.LogUtil) *UserTypeController {
	return &UserTypeController{Service: service, Auth: auth, Log: log}
}

func (utc *UserTypeController) GetUserTypeByID(c *gin.Context) {
	permissionId := config.PERMISSION_GET_USER_TYPE_BY_ID

	if !utc.Auth.CheckPermission(c, permissionId) {
		return
	}

	idParam := c.Param("id")

	if utc.Log.RegisterLog(c, "Attempting to retrieve user type with ID: "+idParam) != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error registering log"})
		return
	}

	var id uint
	if _, err := fmt.Sscanf(idParam, "%d", &id); err != nil {
		_ = utc.Log.RegisterLog(c, "Invalid user type ID format: "+idParam)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user type ID"})
		return
	}

	userType, err := utc.Service.GetUserTypeByID(id)
	if err != nil {
		_ = utc.Log.RegisterLog(c, "User type not found with ID: "+idParam)
		c.JSON(http.StatusNotFound, gin.H{"error": "User type not found"})
		return
	}

	roleIDs, err := utc.Service.GetRolesForUserType(id)
	if err != nil {
		_ = utc.Log.RegisterLog(c, "Error retrieving roles for user type with ID: "+idParam)
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

	_ = utc.Log.RegisterLog(c, "Successfully retrieved user type with ID: "+idParam)
	c.JSON(http.StatusOK, userTypeDTO)
}

func (utc *UserTypeController) GetAllUserTypes(c *gin.Context) {
	permissionId := config.PERMISSION_GET_ALL_USER_TYPES

	if !utc.Auth.CheckPermission(c, permissionId) {
		return
	}

	if utc.Log.RegisterLog(c, "Attempting to retrieve all user types") != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error registering log"})
		return
	}

	userTypes, err := utc.Service.ObtainAllUserTypes()
	if err != nil {
		_ = utc.Log.RegisterLog(c, "Error retrieving all user types")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving user types"})
		return
	}

	var userTypesDTO []dtos.UserTypeDTO
	for _, userType := range userTypes {
		roleIDs, err := utc.Service.GetRolesForUserType(userType.ID)
		if err != nil {
			_ = utc.Log.RegisterLog(c, fmt.Sprintf("Error retrieving roles for user type ID: %d", userType.ID))
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

	_ = utc.Log.RegisterLog(c, "Successfully retrieved all user types")
	c.JSON(http.StatusOK, userTypesDTO)
}

func (utc *UserTypeController) ExistsUserType(c *gin.Context) {
	permissionId := config.PERMISSION_EXIST_USER_TYPE

	if !utc.Auth.CheckPermission(c, permissionId) {
		return
	}

	idParam := c.Param("id")

	if utc.Log.RegisterLog(c, "Attempting to check existence of user type with ID: "+idParam) != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error registering log"})
		return
	}

	var id uint
	if _, err := fmt.Sscanf(idParam, "%d", &id); err != nil {
		_ = utc.Log.RegisterLog(c, "Invalid user type ID format: "+idParam)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user type ID"})
		return
	}

	exists, err := utc.Service.Exists(id)
	if err != nil {
		_ = utc.Log.RegisterLog(c, "Error checking existence for user type ID: "+idParam)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error checking user type existence"})
		return
	}

	_ = utc.Log.RegisterLog(c, "Checked existence of user type ID: "+idParam+", exists: "+fmt.Sprintf("%v", exists))
	c.JSON(http.StatusOK, gin.H{"exists": exists})
}

func (utc *UserTypeController) SearchUserTypesByID(c *gin.Context) {
	permissionId := config.PERMISSION_SEARCH_USER_TYPES_BY_ID

	if !utc.Auth.CheckPermission(c, permissionId) {
		return
	}

	query := c.Query("id")

	if utc.Log.RegisterLog(c, "Attempting to search user types by ID: "+query) != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error registering log"})
		return
	}

	userTypes, err := utc.Service.SearchUserTypesByID(query)
	if err != nil {
		_ = utc.Log.RegisterLog(c, "Error retrieving user types by ID query: "+query)
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

	_ = utc.Log.RegisterLog(c, "Successfully searched user types by ID query: "+query)
	c.JSON(http.StatusOK, userTypesDTO)
}

func (utc *UserTypeController) SearchUserTypesByName(c *gin.Context) {
	permissionId := config.PERMISSION_SEARCH_USER_TYPES_BY_NAME

	if !utc.Auth.CheckPermission(c, permissionId) {
		return
	}

	query := c.Query("name")

	if utc.Log.RegisterLog(c, "Attempting to search user types by name: "+query) != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error registering log"})
		return
	}

	userTypes, err := utc.Service.SearchUserTypesByName(query)
	if err != nil {
		_ = utc.Log.RegisterLog(c, "Error retrieving user types by name query: "+query)
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

	_ = utc.Log.RegisterLog(c, "Successfully searched user types by name query: "+query)
	c.JSON(http.StatusOK, userTypesDTO)
}
