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

type RoleController struct {
	Service *services.RoleService
	Auth    *utilities.AuthorizationUtil
	Log     *utilities.LogUtil
}

func NewRoleController(service *services.RoleService, auth *utilities.AuthorizationUtil, log *utilities.LogUtil) *RoleController {
	return &RoleController{Service: service, Auth: auth, Log: log}
}

func (rc *RoleController) GetRoleByID(c *gin.Context) {
	permissionId := config.PERMISSION_GET_ROLE_BY_ID

	if !rc.Auth.CheckPermission(c, permissionId) {
		_ = rc.Log.RegisterLog(c, "Access denied for GetRoleByID")
		return
	}

	idParam := c.Param("id")

	if rc.Log.RegisterLog(c, "Attempting to retrieve role with ID: "+idParam) != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error registering log"})
		return
	}

	var id uint
	if _, err := fmt.Sscanf(idParam, "%d", &id); err != nil {
		_ = rc.Log.RegisterLog(c, "Invalid role ID format: "+idParam)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid role ID"})
		return
	}

	role, err := rc.Service.GetRoleByID(id)
	if err != nil {
		_ = rc.Log.RegisterLog(c, "Role not found with ID: "+idParam)
		c.JSON(http.StatusNotFound, gin.H{"error": "Role not found"})
		return
	}

	permissionIDs, err := rc.Service.GetRolePermissions(id)
	if err != nil {
		_ = rc.Log.RegisterLog(c, "Error retrieving role permissions for ID: "+idParam)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving role permissions"})
		return
	}

	roleDTO := dtos.RoleDTO{
		ID:          role.ID,
		Name:        role.Name,
		Description: role.Description,
		Permissions: make([]string, len(permissionIDs)),
	}

	for i, permissionID := range permissionIDs {
		roleDTO.Permissions[i] = fmt.Sprintf("%d", permissionID)
	}

	_ = rc.Log.RegisterLog(c, "Successfully retrieved role with ID: "+idParam)
	c.JSON(http.StatusOK, roleDTO)
}

func (rc *RoleController) GetAllRoles(c *gin.Context) {
	permissionId := config.PERMISSION_GET_ALL_ROLES

	if !rc.Auth.CheckPermission(c, permissionId) {
		_ = rc.Log.RegisterLog(c, "Access denied for GetAllRoles")
		return
	}

	if rc.Log.RegisterLog(c, "Attempting to retrieve all roles") != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error registering log"})
		return
	}

	roles, err := rc.Service.GetAllRoles()
	if err != nil {
		_ = rc.Log.RegisterLog(c, "Error retrieving roles")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving roles"})
		return
	}

	var rolesDTO []dtos.RoleDTO
	for _, role := range roles {
		permissionIDs, err := rc.Service.GetRolePermissions(role.ID)
		if err != nil {
			_ = rc.Log.RegisterLog(c, "Error retrieving permissions for role ID: "+fmt.Sprintf("%d", role.ID))
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving role permissions"})
			return
		}

		roleDTO := dtos.RoleDTO{
			ID:          role.ID,
			Name:        role.Name,
			Description: role.Description,
			Permissions: make([]string, len(permissionIDs)),
		}

		for i, permissionID := range permissionIDs {
			roleDTO.Permissions[i] = fmt.Sprintf("%d", permissionID)
		}

		rolesDTO = append(rolesDTO, roleDTO)
	}

	_ = rc.Log.RegisterLog(c, "Successfully retrieved all roles")
	c.JSON(http.StatusOK, rolesDTO)
}

func (rc *RoleController) GetAllPermissionsOfRole(c *gin.Context) {
	permissionId := config.PERMISSION_GET_ALL_PERMISSIONS_OF_ROLE

	if !rc.Auth.CheckPermission(c, permissionId) {
		_ = rc.Log.RegisterLog(c, "Access denied for GetAllPermissionsOfRole")
		return
	}

	roleIDParam := c.Param("id")

	if rc.Log.RegisterLog(c, "Attempting to retrieve permissions for role ID: "+roleIDParam) != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error registering log"})
		return
	}

	var roleID uint
	if _, err := fmt.Sscanf(roleIDParam, "%d", &roleID); err != nil {
		_ = rc.Log.RegisterLog(c, "Invalid role ID format: "+roleIDParam)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid role ID"})
		return
	}

	permissions, err := rc.Service.GetAllPermissionsOfRole(roleID)
	if err != nil {
		_ = rc.Log.RegisterLog(c, "Error retrieving permissions for role ID: "+roleIDParam)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving permissions for role"})
		return
	}

	_ = rc.Log.RegisterLog(c, "Successfully retrieved permissions for role ID: "+roleIDParam)
	c.JSON(http.StatusOK, permissions)
}

func (rc *RoleController) ExistRole(c *gin.Context) {
	permissionId := config.PERMISSION_EXIST_ROLE

	if !rc.Auth.CheckPermission(c, permissionId) {
		_ = rc.Log.RegisterLog(c, "Access denied for ExistRole")
		return
	}

	idParam := c.Param("id")

	if rc.Log.RegisterLog(c, "Attempting to check existence of role with ID: "+idParam) != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error registering log"})
		return
	}

	var id uint
	if _, err := fmt.Sscanf(idParam, "%d", &id); err != nil {
		_ = rc.Log.RegisterLog(c, "Invalid role ID format: "+idParam)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid role ID"})
		return
	}

	exists, err := rc.Service.ExistRole(id)
	if err != nil {
		_ = rc.Log.RegisterLog(c, "Error checking existence of role ID: "+idParam)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error checking role existence"})
		return
	}

	_ = rc.Log.RegisterLog(c, "Checked existence of role ID: "+idParam+" Result: "+fmt.Sprintf("%v", exists))
	c.JSON(http.StatusOK, gin.H{"exists": exists})
}

func (rc *RoleController) SearchRolesByID(c *gin.Context) {
	permissionId := config.PERMISSION_SEARCH_ROLE_BY_ID

	if !rc.Auth.CheckPermission(c, permissionId) {
		_ = rc.Log.RegisterLog(c, "Access denied for SearchRolesByID")
		return
	}

	query := c.Query("id")

	if rc.Log.RegisterLog(c, "Attempting to search roles by ID: "+query) != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error registering log"})
		return
	}

	roles, err := rc.Service.SearchRolesByID(query)
	if err != nil {
		_ = rc.Log.RegisterLog(c, "Error searching roles by ID: "+query)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error searching roles by ID"})
		return
	}

	_ = rc.Log.RegisterLog(c, "Successfully searched roles by ID: "+query)
	c.JSON(http.StatusOK, roles)
}

func (rc *RoleController) SearchRolesByName(c *gin.Context) {
	permissionId := config.PERMISSION_SEARCH_ROLE_BY_NAME

	if !rc.Auth.CheckPermission(c, permissionId) {
		_ = rc.Log.RegisterLog(c, "Access denied for SearchRolesByName")
		return
	}

	query := c.Query("name")

	if rc.Log.RegisterLog(c, "Attempting to search roles by name: "+query) != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error registering log"})
		return
	}

	roles, err := rc.Service.SearchRolesByName(query)
	if err != nil {
		_ = rc.Log.RegisterLog(c, "Error searching roles by name: "+query)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error searching roles by name"})
		return
	}

	_ = rc.Log.RegisterLog(c, "Successfully searched roles by name: "+query)
	c.JSON(http.StatusOK, roles)
}
