package controllers

import (
	"net/http"
	"strconv"
	"totesbackend/config"
	"totesbackend/controllers/utilities"
	"totesbackend/dtos"
	"totesbackend/models"
	"totesbackend/services"

	"github.com/gin-gonic/gin"
)

type ExternalSaleController struct {
	Service *services.ExternalSaleService
	Auth    *utilities.AuthorizationUtil
	Log     *utilities.LogUtil
}

func NewExternalSaleController(service *services.ExternalSaleService, auth *utilities.AuthorizationUtil, log *utilities.LogUtil) *ExternalSaleController {
	return &ExternalSaleController{Service: service, Auth: auth, Log: log}
}

func (esc *ExternalSaleController) GetExternalSaleByID(c *gin.Context) {
	id := c.Param("id")

	if esc.Log.RegisterLog(c, "Fetching external sale by ID: "+id) != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error registering log"})
		return
	}

	permissionId := config.PERMISSION_GET_EXTERNAL_SALE_BY_ID
	if !esc.Auth.CheckPermission(c, permissionId) {
		_ = esc.Log.RegisterLog(c, "Access denied for GetExternalSaleByID")
		return
	}

	externalSale, err := esc.Service.GetExternalSaleByID(id)
	if err != nil {
		_ = esc.Log.RegisterLog(c, "External Sale not found with ID: "+id)
		c.JSON(http.StatusNotFound, gin.H{"error": "External Sale not found"})
		return
	}

	dto := dtos.GetExternalSaleDTO{
		ID:            externalSale.ID,
		ReporterName:  externalSale.ReporterName,
		ReporterID:    externalSale.ReporterID,
		ItemID:        externalSale.Item.ID,
		ItemName:      externalSale.Item.Name,
		CustomerID:    externalSale.Customer.ID,
		CustomerEmail: externalSale.Customer.Email,
		Stock:         externalSale.Stock,
	}

	_ = esc.Log.RegisterLog(c, "Successfully fetched external sale with ID: "+id)

	c.JSON(http.StatusOK, dto)
}

func (esc *ExternalSaleController) GetAllExternalSales(c *gin.Context) {
	if esc.Log.RegisterLog(c, "Fetching all external sales") != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error registering log"})
		return
	}

	permissionId := config.PERMISSION_GET_ALL_EXTERNAL_SALES
	if !esc.Auth.CheckPermission(c, permissionId) {
		_ = esc.Log.RegisterLog(c, "Access denied for GetAllExternalSales")
		return
	}

	externalSales, err := esc.Service.GetAllExternalSales()
	if err != nil {
		_ = esc.Log.RegisterLog(c, "Error retrieving external sales")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving external sales"})
		return
	}

	var externalSalesDTO []dtos.GetExternalSaleDTO
	for _, sale := range externalSales {
		externalSaleDTO := dtos.GetExternalSaleDTO{
			ID:            sale.ID,
			ReporterName:  sale.ReporterName,
			ReporterID:    sale.ReporterID,
			ItemID:        sale.Item.ID,
			ItemName:      sale.Item.Name,
			CustomerID:    sale.Customer.ID,
			CustomerEmail: sale.Customer.Email,
			Stock:         sale.Stock,
		}

		externalSalesDTO = append(externalSalesDTO, externalSaleDTO)
	}

	_ = esc.Log.RegisterLog(c, "Successfully retrieved all external sales")

	c.JSON(http.StatusOK, externalSalesDTO)
}

func (esc *ExternalSaleController) CreateExternalSale(c *gin.Context) {

	if esc.Log.RegisterLog(c, "Creating new external sale") != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error registering log"})
		return
	}

	var dto dtos.CreateExternalSaleDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		_ = esc.Log.RegisterLog(c, "Invalid JSON format for external sale")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
		return
	}

	_ = esc.Log.RegisterLog(c, "Received request to create external sale: "+dto.ReporterName)

	externalSale := models.ExternalSale{
		ReporterName: dto.ReporterName,
		ReporterID:   dto.ReporterID,
		ItemID:       dto.ItemID,
		Stock:        dto.Stock,
		Customer: models.Customer{
			CustomerName:     dto.CustomerName,
			LastName:         dto.LastName,
			Email:            dto.Email,
			IdentifierTypeID: dto.IdentifierTypeID,
			IsBusiness:       dto.IsBusiness,
			Address:          dto.Address,
			PhoneNumbers:     dto.PhoneNumbers,
			CustomerState:    true,
		},
	}

	externalSaleWithID, err := esc.Service.CreateExternalSale(&externalSale)
	if err != nil {
		_ = esc.Log.RegisterLog(c, "Error creating external sale: "+dto.ReporterName)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating external sale"})
		return
	}

	dtoResponse := dtos.GetExternalSaleDTO{
		ID:            externalSaleWithID.ID,
		ItemName:      externalSaleWithID.Item.Name,
		ReporterName:  externalSaleWithID.ReporterName,
		ReporterID:    externalSaleWithID.ReporterID,
		ItemID:        externalSaleWithID.ItemID,
		CustomerID:    externalSaleWithID.CustomerID,
		CustomerEmail: externalSaleWithID.Customer.Email,
		Stock:         externalSaleWithID.Stock,
	}

	_ = esc.Log.RegisterLog(c, "Successfully created external sale with ID: "+strconv.Itoa(dtoResponse.ID))

	c.JSON(http.StatusCreated, dtoResponse)
}
