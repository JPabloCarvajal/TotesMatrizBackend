package controllers

import (
	"fmt"
	"net/http"
	"totesbackend/dtos"
	"totesbackend/models"
	"totesbackend/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AdditionalExpenseController struct {
	Service *services.AdditionalExpenseService
}

func NewAdditionalExpenseController(service *services.AdditionalExpenseService) *AdditionalExpenseController {
	return &AdditionalExpenseController{Service: service}
}

func (aec *AdditionalExpenseController) GetAllAdditionalExpenses(c *gin.Context) {
	username := c.GetHeader("Username")
	fmt.Println("Request made by user:", username)

	additionalExpenses, err := aec.Service.GetAllAdditionalExpenses()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving additional expenses"})
		return
	}
	c.JSON(http.StatusOK, additionalExpenses)
}

func (aec *AdditionalExpenseController) GetAdditionalExpenseByID(c *gin.Context) {
	idParam := c.Param("id")

	additionalExpense, err := aec.Service.GetAdditionalExpenseByID(idParam)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving Additional Expense"})
		return
	}

	if additionalExpense == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Additional Expense not found"})
		return
	}

	c.JSON(http.StatusOK, additionalExpense)
}

func (aec *AdditionalExpenseController) CreateAdditionalExpense(c *gin.Context) {
	var dto dtos.UpdateAdditionalExpenseDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
		return
	}

	newExpense := &models.AdditionalExpense{
		Name:         dto.Name,
		ItemID:       dto.ItemID,
		Expense:      dto.Expense,
		IsPercentage: dto.IsPercentage,
		Description:  dto.Description,
	}

	createdExpense, err := aec.Service.CreateAdditionalExpense(newExpense)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating additional expense"})
		return
	}

	c.JSON(http.StatusCreated, createdExpense)
}

func (aec *AdditionalExpenseController) DeleteAdditionalExpense(c *gin.Context) {
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
	expense.IsPercentage = dto.IsPercentage
	expense.Description = dto.Description

	updatedExpense, err := aec.Service.UpdateAdditionalExpense(expense)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating AdditionalExpense"})
		return
	}

	c.JSON(http.StatusOK, updatedExpense)
}
