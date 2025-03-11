package repositories

import (
	"totesbackend/models"

	"gorm.io/gorm"
)

type AdditionalExpenseRepository struct {
	DB *gorm.DB
}

func NewAdditionalExpenseRepository(db *gorm.DB) *AdditionalExpenseRepository {
	return &AdditionalExpenseRepository{DB: db}
}

func (r *AdditionalExpenseRepository) GetAllAdditionalExpenses() ([]models.AdditionalExpense, error) {
	var additionalExpenses []models.AdditionalExpense
	err := r.DB.Find(&additionalExpenses).Error
	return additionalExpenses, err
}

func (r *AdditionalExpenseRepository) GetAdditionalExpenseByID(id string) (*models.AdditionalExpense, error) {
	var additionalExpense models.AdditionalExpense
	err := r.DB.First(&additionalExpense, id).Error
	if err != nil {
		return nil, err
	}
	return &additionalExpense, nil
}

func (r *AdditionalExpenseRepository) CreateAdditionalExpense(expense *models.AdditionalExpense) (*models.AdditionalExpense, error) {
	err := r.DB.Create(expense).Error
	if err != nil {
		return nil, err
	}
	return expense, nil
}

func (r *AdditionalExpenseRepository) DeleteAdditionalExpense(id string) error {
	result := r.DB.Delete(&models.AdditionalExpense{}, id)
	return result.Error
}

func (r *AdditionalExpenseRepository) UpdateAdditionalExpense(expense *models.AdditionalExpense) (*models.AdditionalExpense, error) {
	err := r.DB.Save(expense).Error
	if err != nil {
		return nil, err
	}

	return expense, nil
}
