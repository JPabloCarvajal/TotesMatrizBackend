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
	router.POST("/item", controller.CreateItem)
}

func RegisterPermissionRoutes(router *gin.Engine, controller *controllers.PermissionController) {

	router.GET("/permission", controller.GetAllPermissions)
	router.GET("/permission/:id", controller.GetPermissionByID)

}

func RegisterRoleRoutes(router *gin.Engine, controller *controllers.RoleController) {
	router.GET("/roles/:id", controller.GetRoleByID)

	router.GET("/roles/:id/permission", controller.GetAllPermissionsOfRole)
	router.GET("/roles/:id/exist", controller.ExistRole)
	router.GET("/roles/", controller.GetAllRoles)
}

func RegisterUserTypeRoutes(router *gin.Engine, controller *controllers.UserTypeController) {
	router.GET("/user-types", controller.GetAllUserTypes)
	router.GET("/user-types/:id", controller.GetUserTypeByID)
	router.GET("/user-types/:id/exists", controller.ExistsUserType)
}

func RegisterUserStateTypeRoutes(router *gin.Engine, controller *controllers.UserStateTypeController) {
	router.GET("/user-state-type", controller.GetUserStateTypes)
	router.GET("/user-state-type/:id", controller.GetUserStateTypeByID)
}

func RegisterIdentifierTypeRoutes(router *gin.Engine, controller *controllers.IdentifierTypeController) {
	router.GET("/identifier-type", controller.GetIdentifierTypes)
	router.GET("/identifier-type/:id", controller.GetIdentifierTypeByID)
}

func RegisterUserRoutes(router *gin.Engine, controller *controllers.UserController) {
	router.GET("/user", controller.GetAllUsers)
	router.GET("/user/:id", controller.GetUserByID)
	router.GET("/user/searchByID", controller.SearchUsersByID)
	router.GET("/user/searchByEmail", controller.SearchUsersByEmail)
	router.PATCH("/user/:id/state", controller.UpdateUserState)
	router.PUT("/user/:id", controller.UpdateUser)
	router.POST("/user", controller.CreateUser)
}

func RegisterEmployeeRoutes(router *gin.Engine, controller *controllers.EmployeeController) {

	router.GET("/employee/", controller.GetAllEmployees)
	router.GET("/employee/:id", controller.GetEmployeeByID)

	router.GET("/employee/searchEmployeesByName", controller.SearchEmployeesByName)
	router.POST("/employee/", controller.CreateEmployee)
	router.PUT("/employee/:id", controller.UpdateEmployee)

}

func RegisterAdditionalExpenseRoutes(router *gin.Engine, controller *controllers.AdditionalExpenseController) {
	router.GET("/additional-expense", controller.GetAllAdditionalExpenses)
	router.GET("/additional-expense/:id", controller.GetAdditionalExpenseByID)
	router.POST("/additional-expense", controller.CreateAdditionalExpense)
	router.PUT("/additional-expense/:id", controller.UpdateAdditionalExpense)
	router.DELETE("/additional-expense/:id", controller.DeleteAdditionalExpense)
}

func RegisterHistoricalItemPriceRoutes(router *gin.Engine, controller *controllers.HistoricalItemPriceController) {
	router.GET("/historical-item-price/:id", controller.GetHistoricalItemPrice)
}

func RegisterCommentRoutes(router *gin.Engine, controller *controllers.CommentController) {
	router.GET("/comment/:id", controller.GetCommentByID)
	router.GET("/comments", controller.GetAllComments)
	router.GET("/comments/searchByEmail", controller.SearchCommentsByEmail)
	router.POST("/comment", controller.CreateComment)
	router.PUT("/comment/:id", controller.UpdateComment)
}

func RegisterAuthorizationRoutes(router *gin.Engine, controller *controllers.AuthorizationController) {
	router.GET("/auth/check-permission", controller.CheckUserPermission)
}
