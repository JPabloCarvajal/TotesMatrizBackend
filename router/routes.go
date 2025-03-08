package routes

import (
	"totesbackend/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterItemTypeRoutes(router *gin.Engine, controller *controllers.ItemTypeController) {

	router.GET("/item-type", controller.GetItemTypes)
	router.GET("/item-type/:id", controller.GetItemTypeByID)

}

func RegisterItemRoutes(router *gin.Engine, controller *controllers.ItemController) {

	router.GET("/item/:id", controller.GetItemByID)
	router.GET("/item", controller.GetAllItems)
	router.GET("/item/searchById", controller.SearchItemsByID)
	router.GET("/item/searchByName", controller.SearchItemsByName)
	router.PATCH("/item/:id/state", controller.UpdateItemState)
	router.PUT("/item/:id", controller.UpdateItem)
}

func RegisterPermissionRoutes(router *gin.Engine, controller *controllers.PermissionController) {
	permissions := router.Group("/permissions")
	{
		permissions.GET("/", controller.GetAllPermissions)    //sirve
		permissions.GET("/:id", controller.GetPermissionByID) //sirve
	}
}

func RegisterRoleRoutes(router *gin.Engine, controller *controllers.RoleController) {
	roles := router.Group("/roles")
	{
		roles.GET("/:id", controller.GetRoleByID)                         // sirve
		roles.GET("/:id/permissions", controller.GetAllPermissionsOfRole) // sirve
		roles.GET("/:id/exists", controller.ExistRole)                    // sirve
	}
}

func RegisterUserTypeRoutes(router *gin.Engine, controller *controllers.UserTypeController) {
	userTypes := router.Group("/user-types")
	{
		userTypes.GET("/", controller.ObtainAllUserTypes)
		userTypes.GET("/:id", controller.ObtainUserTypeByID)
		userTypes.GET("/:id/exists", controller.Exists)

	}
}
