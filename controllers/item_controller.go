package controllers

import (
	"fmt"
	"net/http"

	"totesbackend/dtos"
	"totesbackend/services"

	"github.com/gin-gonic/gin"
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

	additionalExpenseIDs := make([]string, len(item.AdditionalExpenses))
	for i, expense := range item.AdditionalExpenses {
		additionalExpenseIDs[i] = fmt.Sprintf("%d", expense.ID)
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
		additionalExpenseIDs := make([]string, len(item.AdditionalExpenses))
		for i, expense := range item.AdditionalExpenses {
			additionalExpenseIDs[i] = fmt.Sprintf("%d", expense.ID)
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
		additionalExpenseIDs := make([]string, len(item.AdditionalExpenses))
		for i, expense := range item.AdditionalExpenses {
			additionalExpenseIDs[i] = fmt.Sprintf("%d", expense.ID)
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
		additionalExpenseIDs := make([]string, len(item.AdditionalExpenses))
		for i, expense := range item.AdditionalExpenses {
			additionalExpenseIDs[i] = fmt.Sprintf("%d", expense.ID)
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

	additionalExpenseIDs := make([]string, len(item.AdditionalExpenses))
	for i, expense := range item.AdditionalExpenses {
		additionalExpenseIDs[i] = fmt.Sprintf("%d", expense.ID)
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
