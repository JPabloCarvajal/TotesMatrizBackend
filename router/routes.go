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

func RegisterUserStateTypeRoutes(router *gin.Engine, controller *controllers.UserStateTypeController) {
	router.GET("/user-state-type", controller.GetUserStateTypes)
	router.GET("/user-state-type/:id", controller.GetUserStateTypeByID)
}

func RegisterIdentifierTypeRoutes(router *gin.Engine, controller *controllers.IdentifierTypeController) {
	router.GET("/identifier-type", controller.GetIdentifierTypes)
	router.GET("/identifier-type/:id", controller.GetIdentifierTypeByID)
}
