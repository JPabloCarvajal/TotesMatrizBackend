package controllers

import (
	"fmt"
	"net/http"
	"totesbackend/config"
	"totesbackend/controllers/utilities"
	"totesbackend/services"

	"github.com/gin-gonic/gin"
)

type PermissionController struct {
	Service *services.PermissionService
	Auth    *utilities.AuthorizationUtil
	Log     *utilities.LogUtil
}

func NewPermissionController(service *services.PermissionService, auth *utilities.AuthorizationUtil, log *utilities.LogUtil) *PermissionController {
	return &PermissionController{Service: service, Auth: auth, Log: log}
}

func (pc *PermissionController) GetPermissionByID(c *gin.Context) {
	permissionId := config.PERMISSION_GET_PERMISSION_BY_ID

	if !pc.Auth.CheckPermission(c, permissionId) {
		if pc.Log.RegisterLog(c, "Access denied for GetPermissionByID") != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error registering log"})
			return
		}
		return
	}

	idParam := c.Param("id")
	var id uint
	if _, err := fmt.Sscanf(idParam, "%d", &id); err != nil {
		if pc.Log.RegisterLog(c, "Invalid permission ID: "+idParam) != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error registering log"})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid permission ID"})
		return
	}

	if pc.Log.RegisterLog(c, "Attempting to retrieve Permission with ID: "+idParam) != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error registering log"})
		return
	}

	permission, err := pc.Service.GetPermissionByID(id)
	if err != nil {
		if pc.Log.RegisterLog(c, "Permission with ID "+idParam+" not found") != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error registering log"})
			return
		}
		c.JSON(http.StatusNotFound, gin.H{"error": "Permission not found"})
		return
	}

	if pc.Log.RegisterLog(c, "Successfully retrieved Permission with ID: "+idParam) != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error registering log"})
		return
	}

	c.JSON(http.StatusOK, permission)
}

func (pc *PermissionController) GetAllPermissions(c *gin.Context) {
	permissionId := config.PERMISSION_GET_ALL_PERMISSIONS

	if !pc.Auth.CheckPermission(c, permissionId) {
		if pc.Log.RegisterLog(c, "Access denied for GetAllPermissions") != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error registering log"})
			return
		}
		return
	}

	if pc.Log.RegisterLog(c, "Attempting to retrieve all permissions") != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error registering log"})
		return
	}

	permissions, err := pc.Service.GetAllPermissions()
	if err != nil {
		if pc.Log.RegisterLog(c, "Error retrieving all permissions: "+err.Error()) != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error registering log"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving permissions"})
		return
	}

	if pc.Log.RegisterLog(c, "Successfully retrieved all permissions") != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error registering log"})
		return
	}

	c.JSON(http.StatusOK, permissions)
}

func (pc *PermissionController) SearchPermissionsByID(c *gin.Context) {
	permissionId := config.PERMISSION_SEARCH_PERMISSION_BY_ID

	if !pc.Auth.CheckPermission(c, permissionId) {
		if pc.Log.RegisterLog(c, "Access denied for SearchPermissionsByID") != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error registering log"})
			return
		}
		return
	}

	query := c.Query("id")
	if query == "" {
		if pc.Log.RegisterLog(c, "SearchPermissionsByID: missing 'id' query parameter") != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error registering log"})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": "Search query is required"})
		return
	}

	if pc.Log.RegisterLog(c, "Attempting to search permissions by ID: "+query) != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error registering log"})
		return
	}

	permissions, err := pc.Service.SearchPermissionsByID(query)
	if err != nil {
		if pc.Log.RegisterLog(c, "Error retrieving permissions by ID: "+err.Error()) != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error registering log"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving permissions"})
		return
	}

	if pc.Log.RegisterLog(c, "Successfully retrieved permissions by ID: "+query) != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error registering log"})
		return
	}

	c.JSON(http.StatusOK, permissions)
}

func (pc *PermissionController) SearchPermissionsByName(c *gin.Context) {
	permissionId := config.PERMISSION_SEARCH_PERMISSION_BY_NAME

	if !pc.Auth.CheckPermission(c, permissionId) {
		if pc.Log.RegisterLog(c, "Access denied for SearchPermissionsByName") != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error registering log"})
			return
		}
		return
	}

	query := c.Query("name")
	if query == "" {
		if pc.Log.RegisterLog(c, "SearchPermissionsByName: missing 'name' query parameter") != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error registering log"})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": "Search query is required"})
		return
	}

	if pc.Log.RegisterLog(c, "Attempting to search permissions by name: "+query) != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error registering log"})
		return
	}

	permissions, err := pc.Service.SearchPermissionsByName(query)
	if err != nil {
		if pc.Log.RegisterLog(c, "Error retrieving permissions by name: "+err.Error()) != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error registering log"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving permissions"})
		return
	}

	if pc.Log.RegisterLog(c, "Successfully retrieved permissions by name: "+query) != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error registering log"})
		return
	}

	c.JSON(http.StatusOK, permissions)
}
