package app

import (
	"totesbackend/config"
	"totesbackend/controllers"
	"totesbackend/database"
	"totesbackend/repositories"
	routes "totesbackend/router"
	"totesbackend/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db *gorm.DB
var router *gin.Engine

func SetupAndRunApp() error {
	// load env
	err := config.LoadENV()
	if err != nil {
		return err
	}

	// start database
	err = database.StartPostgres()
	if err != nil {
		return err
	}

	// defer closing database
	defer database.ClosePostgres()

	db = database.GetDB()
	router = gin.Default()
	database.MigrateDB()

	setUpItemTypeRouter()
	setUpItemRouter()

	router.Run("localhost:8080")

	return nil
}

func setUpItemTypeRouter() {
	itemTypeRepo := repositories.NewItemTypeRepository(db)
	itemTypeService := services.NewItemTypeService(itemTypeRepo)
	itemTypeController := controllers.NewItemTypeController(itemTypeService)
	routes.RegisterItemTypeRoutes(router, itemTypeController)
}

func setUpItemRouter() {
	itemRepo := repositories.NewItemRepository(db)
	itemService := services.NewItemService(itemRepo)
	itemController := controllers.NewItemController(itemService)
	routes.RegisterItemRoutes(router, itemController)
}
