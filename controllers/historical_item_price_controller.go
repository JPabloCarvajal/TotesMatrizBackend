package controllers

import (
	"net/http"
	"totesbackend/config"
	"totesbackend/controllers/utilities"
	"totesbackend/services"

	"github.com/gin-gonic/gin"
)

type HistoricalItemPriceController struct {
	Service *services.HistoricalItemPriceService
	Auth    *utilities.AuthorizationUtil
}

func NewHistoricalItemPriceController(service *services.HistoricalItemPriceService, auth *utilities.AuthorizationUtil) *HistoricalItemPriceController {
	return &HistoricalItemPriceController{Service: service, Auth: auth}
}

func (c *HistoricalItemPriceController) GetHistoricalItemPrice(ctx *gin.Context) {
	permissionId := config.PERMISSION_GET_HISTORICAL_ITEM_PRICE

	if !c.Auth.CheckPermission(ctx, permissionId) {
		return
	}

	itemID := ctx.Param("id")

	historicalPrices, err := c.Service.GetHistoricalItemPrice(itemID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve historical prices"})
		return
	}

	ctx.JSON(http.StatusOK, historicalPrices)
}
