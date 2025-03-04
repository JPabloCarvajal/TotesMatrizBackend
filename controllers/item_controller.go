package controllers

import (
	"errors"
	"fmt"
	"net/http"

	"totesbackend/dtos"
	"totesbackend/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ItemController struct {
	Service *services.ItemService
}

func NewItemController(service *services.ItemService) *ItemController {
	return &ItemController{Service: service}
}

func (ic *ItemController) GetItemByID(c *gin.Context) {
	username := c.GetHeader("Username")
	fmt.Println("Request made by user:", username)

	id := c.Param("id")

	item, err := ic.Service.GetItemByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}

	additionalExpenseIDs := make([]int, len(item.AdditionalExpenses))
	for i, expense := range item.AdditionalExpenses {
		additionalExpenseIDs[i] = expense.ID
	}

	itemDTO := dtos.GetItemDTO{
		ID:                 item.ID,
		Name:               item.Name,
		Description:        item.Description,
		Stock:              item.Stock,
		SellingPrice:       item.SellingPrice,
		PurchasePrice:      item.PurchasePrice,
		ItemState:          item.ItemState,
		ItemTypeID:         item.ItemTypeID,
		AdditionalExpenses: additionalExpenseIDs,
	}

	c.JSON(http.StatusOK, itemDTO)
}

func (ic *ItemController) GetAllItems(c *gin.Context) {
	username := c.GetHeader("Username")
	fmt.Println("Request made by user:", username)

	items, err := ic.Service.GetAllItems()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving items"})
		return
	}

	var itemsDTO []dtos.GetItemDTO
	for _, item := range items {
		additionalExpenseIDs := make([]int, len(item.AdditionalExpenses))
		for i, expense := range item.AdditionalExpenses {
			additionalExpenseIDs[i] = expense.ID
		}

		itemDTO := dtos.GetItemDTO{
			ID:                 item.ID,
			Name:               item.Name,
			Description:        item.Description,
			Stock:              item.Stock,
			SellingPrice:       item.SellingPrice,
			PurchasePrice:      item.PurchasePrice,
			ItemState:          item.ItemState,
			ItemTypeID:         item.ItemTypeID,
			AdditionalExpenses: additionalExpenseIDs,
		}

		itemsDTO = append(itemsDTO, itemDTO)
	}

	c.JSON(http.StatusOK, itemsDTO)
}

func (ic *ItemController) SearchItemsByID(c *gin.Context) {
	username := c.GetHeader("Username")
	fmt.Println("Request made by user:", username)

	query := c.Query("id")

	if query == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Search query is required"})
		return
	}

	items, err := ic.Service.SearchItemsByID(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving items"})
		return
	}

	if len(items) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No items found"})
		return
	}

	var itemsDTO []dtos.GetItemDTO
	for _, item := range items {
		additionalExpenseIDs := make([]int, len(item.AdditionalExpenses))
		for i, expense := range item.AdditionalExpenses {
			additionalExpenseIDs[i] = expense.ID
		}

		itemDTO := dtos.GetItemDTO{
			ID:                 item.ID,
			Name:               item.Name,
			Description:        item.Description,
			Stock:              item.Stock,
			SellingPrice:       item.SellingPrice,
			PurchasePrice:      item.PurchasePrice,
			ItemState:          item.ItemState,
			ItemTypeID:         item.ItemTypeID,
			AdditionalExpenses: additionalExpenseIDs,
		}

		itemsDTO = append(itemsDTO, itemDTO)
	}

	c.JSON(http.StatusOK, itemsDTO)
}

func (ic *ItemController) SearchItemsByName(c *gin.Context) {
	username := c.GetHeader("Username")
	fmt.Println("Request made by user:", username)

	query := c.Query("name")

	if query == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Search query is required"})
		return
	}

	items, err := ic.Service.SearchItemsByName(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving items"})
		return
	}

	if len(items) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No items found"})
		return
	}

	var itemsDTO []dtos.GetItemDTO
	for _, item := range items {
		additionalExpenseIDs := make([]int, len(item.AdditionalExpenses))
		for i, expense := range item.AdditionalExpenses {
			additionalExpenseIDs[i] = expense.ID
		}

		itemDTO := dtos.GetItemDTO{
			ID:                 item.ID,
			Name:               item.Name,
			Description:        item.Description,
			Stock:              item.Stock,
			SellingPrice:       item.SellingPrice,
			PurchasePrice:      item.PurchasePrice,
			ItemState:          item.ItemState,
			ItemTypeID:         item.ItemTypeID,
			AdditionalExpenses: additionalExpenseIDs,
		}

		itemsDTO = append(itemsDTO, itemDTO)
	}

	c.JSON(http.StatusOK, itemsDTO)
}

func (ic *ItemController) UpdateItemState(c *gin.Context) {
	username := c.GetHeader("Username")
	fmt.Println("Request made by user:", username)

	id := c.Param("id")

	var request struct {
		ItemState bool `json:"item_state"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	item, err := ic.Service.UpdateItemState(id, request.ItemState)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}

	additionalExpenseIDs := make([]int, len(item.AdditionalExpenses))
	for i, expense := range item.AdditionalExpenses {
		additionalExpenseIDs[i] = expense.ID
	}

	itemDTO := dtos.GetItemDTO{
		ID:                 item.ID,
		Name:               item.Name,
		Description:        item.Description,
		Stock:              item.Stock,
		SellingPrice:       item.SellingPrice,
		PurchasePrice:      item.PurchasePrice,
		ItemState:          item.ItemState,
		ItemTypeID:         item.ItemTypeID,
		AdditionalExpenses: additionalExpenseIDs,
	}

	c.JSON(http.StatusOK, itemDTO)

}

func (ic *ItemController) UpdateItem(c *gin.Context) {
	id := c.Param("id") // Obtener el ID del item

	var dto dtos.UpdateItemDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
		return
	}

	// Buscar el item en la base de datos
	item, err := ic.Service.GetItemByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving item"})
		return
	}

	// Asignar los valores del DTO al modelo
	item.Name = dto.Name
	item.Description = dto.Description
	item.Stock = dto.Stock
	item.SellingPrice = dto.SellingPrice
	item.PurchasePrice = dto.PurchasePrice
	item.ItemState = dto.ItemState
	item.ItemTypeID = dto.ItemTypeID

	// Llamar al servicio para actualizar el item
	err = ic.Service.UpdateItem(item)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating item"})
		return
	}

	c.JSON(http.StatusOK, dto)
}
