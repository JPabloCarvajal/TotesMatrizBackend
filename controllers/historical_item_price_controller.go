package controllers

import (
	"net/http"
	"totesbackend/services"

	"github.com/gin-gonic/gin"
)

type HistoricalItemPriceController struct {
	Service *services.HistoricalItemPriceService
}

func NewHistoricalItemPriceController(service *services.HistoricalItemPriceService) *HistoricalItemPriceController {
	return &HistoricalItemPriceController{Service: service}
}

func (c *HistoricalItemPriceController) GetHistoricalItemPrice(ctx *gin.Context) {
	itemID := ctx.Param("id")

	historicalPrices, err := c.Service.GetHistoricalItemPrice(itemID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve historical prices"})
		return
	}

	ctx.JSON(http.StatusOK, historicalPrices)
}
