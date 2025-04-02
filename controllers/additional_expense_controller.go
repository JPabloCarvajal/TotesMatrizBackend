package controllers

import (
	"net/http"
	"totesbackend/config"
	"totesbackend/controllers/utilities"
	"totesbackend/dtos"
	"totesbackend/models"
	"totesbackend/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AdditionalExpenseController struct {
	Service *services.AdditionalExpenseService
	Auth    *utilities.AuthorizationUtil
	Log     *utilities.LogUtil
}

func NewAdditionalExpenseController(service *services.AdditionalExpenseService,
	auth *utilities.AuthorizationUtil, log *utilities.LogUtil) *AdditionalExpenseController {
	return &AdditionalExpenseController{Service: service, Auth: auth, Log: log}
}

func (aec *AdditionalExpenseController) GetAdditionalExpenseByID(c *gin.Context) {
	idParam := c.Param("id")

	// Attempt to retrieve AdditionalExpense, with log validation
	if aec.Log.RegisterLog(c, "Attempting to retrieve AdditionalExpense with ID: "+idParam) != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error registering log"})
		return
	}

	permissionId := config.PERMISSION_GET_ADDITIONAL_EXPENSE_BY_ID
	if !aec.Auth.CheckPermission(c, permissionId) {
		_ = aec.Log.RegisterLog(c, "Access denied for GetAdditionalExpenseByID")
		return
	}

	additionalExpense, err := aec.Service.GetAdditionalExpenseByID(idParam)
	if err != nil {
		_ = aec.Log.RegisterLog(c, "Error retrieving AdditionalExpense with ID "+idParam+": "+err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving Additional Expense"})
		return
	}

	if additionalExpense == nil {
		_ = aec.Log.RegisterLog(c, "AdditionalExpense with ID "+idParam+" not found")
		c.JSON(http.StatusNotFound, gin.H{"error": "Additional Expense not found"})
		return
	}

	_ = aec.Log.RegisterLog(c, "Successfully retrieved AdditionalExpense with ID: "+idParam)

	c.JSON(http.StatusOK, additionalExpense)
}

func (aec *AdditionalExpenseController) GetAllAdditionalExpenses(c *gin.Context) {

	permissionId := config.PERMISSION_GET_ALL_ADDITIONAL_EXPENSE

	if !aec.Auth.CheckPermission(c, permissionId) {
		return
	}

	additionalExpenses, err := aec.Service.GetAllAdditionalExpenses()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving additional expenses"})
		return
	}
	c.JSON(http.StatusOK, additionalExpenses)
}

func (aec *AdditionalExpenseController) CreateAdditionalExpense(c *gin.Context) {
	permissionId := config.PERMISSION_CREATE_ADDITIONAL_EXPENSE

	if !aec.Auth.CheckPermission(c, permissionId) {
		return
	}

	var dto dtos.UpdateAdditionalExpenseDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
		return
	}

	newExpense := &models.AdditionalExpense{
		Name:        dto.Name,
		ItemID:      dto.ItemID,
		Expense:     dto.Expense,
		Description: dto.Description,
	}

	createdExpense, err := aec.Service.CreateAdditionalExpense(newExpense)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating additional expense"})
		return
	}

	c.JSON(http.StatusCreated, createdExpense)
}

func (aec *AdditionalExpenseController) DeleteAdditionalExpense(c *gin.Context) {
	permissionId := config.PERMISSION_DELETE_ADDITIONAL_EXPENSE

	if !aec.Auth.CheckPermission(c, permissionId) {
		return
	}

	id := c.Param("id")

	err := aec.Service.DeleteAdditionalExpense(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Additional Expense not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting Additional Expense"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Additional Expense deleted successfully"})
}

func (aec *AdditionalExpenseController) UpdateAdditionalExpense(c *gin.Context) {
	permissionId := config.PERMISSION_UPDATE_ADDITIONAL_EXPENSE

	if !aec.Auth.CheckPermission(c, permissionId) {
		return
	}

	id := c.Param("id")

	var dto dtos.UpdateAdditionalExpenseDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
		return
	}

	expense, err := aec.Service.GetAdditionalExpenseByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "AdditionalExpense not found"})
		return
	}

	expense.Name = dto.Name
	expense.ItemID = dto.ItemID
	expense.Expense = dto.Expense
	expense.Description = dto.Description

	updatedExpense, err := aec.Service.UpdateAdditionalExpense(expense)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating AdditionalExpense"})
		return
	}

	c.JSON(http.StatusOK, updatedExpense)
}
