package controllers

import (
	"errors"
	"net/http"
	"strconv"

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
	Log     *utilities.LogUtil
}

func NewItemController(service *services.ItemService, auth *utilities.AuthorizationUtil, log *utilities.LogUtil) *ItemController {
	return &ItemController{Service: service, Auth: auth, Log: log}
}
func (ic *ItemController) CheckItemStock(c *gin.Context) {
	idParam := c.Param("id")
	quantityParam := c.Query("quantity")

	if ic.Log.RegisterLog(c, "Checking stock for item ID: "+idParam+" with quantity: "+quantityParam) != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error registering log"})
		return
	}

	permissionId := config.PERMISSION_CHECK_ITEM_STOCK
	if !ic.Auth.CheckPermission(c, permissionId) {
		_ = ic.Log.RegisterLog(c, "Access denied for CheckItemStock")
		return
	}

	quantity, err := strconv.Atoi(quantityParam)
	if err != nil {
		_ = ic.Log.RegisterLog(c, "Invalid quantity: "+quantityParam)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid quantity"})
		return
	}

	hasStock, err := ic.Service.HasEnoughStock(idParam, quantity)
	if err != nil {
		_ = ic.Log.RegisterLog(c, "Error checking stock for item ID "+idParam+": "+err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error checking stock"})
		return
	}

	_ = ic.Log.RegisterLog(c, "Stock check successful for item ID: "+idParam+" - Has enough stock: "+strconv.FormatBool(hasStock))

	c.JSON(http.StatusOK, gin.H{"hasEnoughStock": hasStock})
}
func (ic *ItemController) GetItemByID(c *gin.Context) {
	id := c.Param("id")

	if ic.Log.RegisterLog(c, "Fetching item by ID: "+id) != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error registering log"})
		return
	}
	item, err := ic.Service.GetItemByID(id)
	if err != nil {
		_ = ic.Log.RegisterLog(c, "Item not found with ID: "+id)
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

	_ = ic.Log.RegisterLog(c, "Successfully fetched item with ID: "+id)

	c.JSON(http.StatusOK, itemDTO)
}
func (ic *ItemController) GetAllItems(c *gin.Context) {
	if ic.Log.RegisterLog(c, "Fetching all items") != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error registering log"})
		return
	}

	items, err := ic.Service.GetAllItems()
	if err != nil {
		_ = ic.Log.RegisterLog(c, "Error retrieving items")
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

	_ = ic.Log.RegisterLog(c, "Successfully retrieved all items")

	c.JSON(http.StatusOK, itemsDTO)
}
func (ic *ItemController) SearchItemsByID(c *gin.Context) {
	if ic.Log.RegisterLog(c, "Searching items by ID") != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error registering log"})
		return
	}

	permissionId := config.PERMISSION_SEARCH_ITEMS_BY_ID
	if !ic.Auth.CheckPermission(c, permissionId) {
		_ = ic.Log.RegisterLog(c, "Access denied for SearchItemsByID")
		return
	}

	query := c.Query("id")
	if query == "" {
		_ = ic.Log.RegisterLog(c, "Search query is missing")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Search query is required"})
		return
	}

	items, err := ic.Service.SearchItemsByID(query)
	if err != nil {
		_ = ic.Log.RegisterLog(c, "Error retrieving items from database")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving items"})
		return
	}

	if len(items) == 0 {
		_ = ic.Log.RegisterLog(c, "No items found for query: "+query)
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

	_ = ic.Log.RegisterLog(c, "Successfully retrieved items for query: "+query)

	c.JSON(http.StatusOK, itemsDTO)
}

func (ic *ItemController) SearchItemsByName(c *gin.Context) {
	if ic.Log.RegisterLog(c, "Searching items by name") != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error registering log"})
		return
	}

	permissionId := config.PERMISSION_SEARCH_ITEMS_BY_NAME
	if !ic.Auth.CheckPermission(c, permissionId) {
		_ = ic.Log.RegisterLog(c, "Access denied for SearchItemsByName")
		return
	}

	query := c.Query("name")
	if query == "" {
		_ = ic.Log.RegisterLog(c, "Search query is missing")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Search query is required"})
		return
	}

	items, err := ic.Service.SearchItemsByName(query)
	if err != nil {
		_ = ic.Log.RegisterLog(c, "Error retrieving items from database")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving items"})
		return
	}

	if len(items) == 0 {
		_ = ic.Log.RegisterLog(c, "No items found for query: "+query)
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

	_ = ic.Log.RegisterLog(c, "Successfully retrieved items for query: "+query)

	c.JSON(http.StatusOK, itemsDTO)
}

func (ic *ItemController) UpdateItemState(c *gin.Context) {
	if ic.Log.RegisterLog(c, "Updating item state") != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error registering log"})
		return
	}

	permissionId := config.PERMISSION_UPDATE_ITEM_STATE
	if !ic.Auth.CheckPermission(c, permissionId) {
		_ = ic.Log.RegisterLog(c, "Access denied for UpdateItemState")
		return
	}

	id := c.Param("id")
	_ = ic.Log.RegisterLog(c, "Received request to update state for item ID: "+id)

	var request struct {
		ItemState bool `json:"item_state"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		_ = ic.Log.RegisterLog(c, "Invalid request body")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	item, err := ic.Service.UpdateItemState(id, request.ItemState)
	if err != nil {
		_ = ic.Log.RegisterLog(c, "Item not found with ID: "+id)
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

	_ = ic.Log.RegisterLog(c, "Successfully updated state for item ID: "+id)

	c.JSON(http.StatusOK, itemDTO)
}

func (ic *ItemController) UpdateItem(c *gin.Context) {
	if ic.Log.RegisterLog(c, "Updating item") != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error registering log"})
		return
	}

	permissionId := config.PERMISSION_UPDATE_ITEM
	if !ic.Auth.CheckPermission(c, permissionId) {
		_ = ic.Log.RegisterLog(c, "Access denied for UpdateItem")
		return
	}

	id := c.Param("id") // Obtener el ID del item
	_ = ic.Log.RegisterLog(c, "Received request to update item with ID: "+id)

	var dto dtos.UpdateItemDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		_ = ic.Log.RegisterLog(c, "Invalid JSON format")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
		return
	}

	// Buscar el item en la base de datos
	item, err := ic.Service.GetItemByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			_ = ic.Log.RegisterLog(c, "Item not found with ID: "+id)
			c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
			return
		}
		_ = ic.Log.RegisterLog(c, "Error retrieving item with ID: "+id)
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
		_ = ic.Log.RegisterLog(c, "Error updating item with ID: "+id)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating item"})
		return
	}

	additionalExpenseIDs := make([]int, len(item.AdditionalExpenses))
	for i, expense := range item.AdditionalExpenses {
		additionalExpenseIDs[i] = expense.ID
	}

	dtoGet := dtos.GetItemDTO{
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

	_ = ic.Log.RegisterLog(c, "Successfully updated item with ID: "+id)

	c.JSON(http.StatusOK, dtoGet)
}
func (ic *ItemController) CreateItem(c *gin.Context) {
	if ic.Log.RegisterLog(c, "Creating new item") != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error registering log"})
		return
	}

	permissionId := config.PERMISSION_CREATE_ITEM
	if !ic.Auth.CheckPermission(c, permissionId) {
		_ = ic.Log.RegisterLog(c, "Access denied for CreateItem")
		return
	}

	var dto dtos.UpdateItemDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		_ = ic.Log.RegisterLog(c, "Invalid JSON format")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
		return
	}

	_ = ic.Log.RegisterLog(c, "Received request to create item: "+dto.Name)

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
		_ = ic.Log.RegisterLog(c, "Error creating item: "+dto.Name)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating item"})
		return
	}

	additionalExpenseIDs := make([]int, len(itemWithId.AdditionalExpenses))
	for i, expense := range itemWithId.AdditionalExpenses {
		additionalExpenseIDs[i] = expense.ID
	}

	dtoGet := dtos.GetItemDTO{
		ID:                 itemWithId.ID,
		Name:               itemWithId.Name,
		Description:        itemWithId.Description,
		Stock:              itemWithId.Stock,
		SellingPrice:       itemWithId.SellingPrice,
		PurchasePrice:      itemWithId.PurchasePrice,
		ItemState:          itemWithId.ItemState,
		ItemTypeID:         itemWithId.ItemTypeID,
		AdditionalExpenses: additionalExpenseIDs,
	}

	_ = ic.Log.RegisterLog(c, "Successfully created item with ID: "+strconv.Itoa(dtoGet.ID))

	c.JSON(http.StatusCreated, dtoGet)
}
