package controllers

import (
	"errors"
	"net/http"

	"totesbackend/config"
	"totesbackend/controllers/utilities"
	"totesbackend/dtos"
	"totesbackend/models"
	"totesbackend/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ItemController struct {
	Service *services.ItemService
	Auth    *utilities.AuthorizationUtil
}

func NewItemController(service *services.ItemService, auth *utilities.AuthorizationUtil) *ItemController {
	return &ItemController{Service: service, Auth: auth}
}

func (ic *ItemController) GetItemByID(c *gin.Context) {
	permissionId := config.PERMISSION_GET_ITEM_BY_ID

	if !ic.Auth.CheckPermission(c, permissionId) {
		return
	}

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
	permissionId := config.PERMISSION_GET_ALL_ITEMS

	if !ic.Auth.CheckPermission(c, permissionId) {
		return
	}
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
	permissionId := config.PERMISSION_SEARCH_ITEMS_BY_ID

	if !ic.Auth.CheckPermission(c, permissionId) {
		return
	}

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
	permissionId := config.PERMISSION_SEARCH_ITEMS_BY_NAME

	if !ic.Auth.CheckPermission(c, permissionId) {
		return
	}

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
	permissionId := config.PERMISSION_UPDATE_ITEM_STATE

	if !ic.Auth.CheckPermission(c, permissionId) {
		return
	}

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
	permissionId := config.PERMISSION_UPDATE_ITEM

	if !ic.Auth.CheckPermission(c, permissionId) {
		return
	}

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

	var dtoGet dtos.GetItemDTO
	dtoGet.ID = item.ID
	dtoGet.Name = item.Name
	dtoGet.Description = item.Description
	dtoGet.Stock = item.Stock
	dtoGet.SellingPrice = item.SellingPrice
	dtoGet.PurchasePrice = item.PurchasePrice
	dtoGet.ItemState = item.ItemState
	dtoGet.ItemTypeID = item.ItemTypeID

	c.JSON(http.StatusOK, dtoGet)
}

func (ic *ItemController) CreateItem(c *gin.Context) {
	permissionId := config.PERMISSION_CREATE_ITEM

	if !ic.Auth.CheckPermission(c, permissionId) {
		return
	}

	var dto dtos.UpdateItemDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
		return
	}

	// Crear una instancia del modelo Item con los datos del DTO
	item := models.Item{
		Name:          dto.Name,
		Description:   dto.Description,
		Stock:         dto.Stock,
		SellingPrice:  dto.SellingPrice,
		PurchasePrice: dto.PurchasePrice,
		ItemState:     dto.ItemState,
		ItemTypeID:    dto.ItemTypeID,
	}

	// Llamar al servicio para crear el item
	itemWithId, err := ic.Service.CreateItem(&item)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating item"})
		return
	}

	var dtoGet dtos.GetItemDTO
	dtoGet.ID = itemWithId.ID
	dtoGet.Name = itemWithId.Name
	dtoGet.Description = itemWithId.Description
	dtoGet.Stock = itemWithId.Stock
	dtoGet.SellingPrice = itemWithId.SellingPrice
	dtoGet.PurchasePrice = itemWithId.PurchasePrice
	dtoGet.ItemState = itemWithId.ItemState
	dtoGet.ItemTypeID = itemWithId.ItemTypeID

	c.JSON(http.StatusCreated, dtoGet)
}
