package services

import (
	"totesbackend/models"
	"totesbackend/repositories"
)

type AdditionalExpenseService struct {
	Repo *repositories.AdditionalExpenseRepository
}

func NewAdditionalExpenseService(repo *repositories.AdditionalExpenseRepository) *AdditionalExpenseService {
	return &AdditionalExpenseService{Repo: repo}
}

func (s *AdditionalExpenseService) GetAllAdditionalExpenses() ([]models.AdditionalExpense, error) {
	return s.Repo.GetAllAdditionalExpenses()
}

func (s *AdditionalExpenseService) GetAdditionalExpenseByID(id string) (*models.AdditionalExpense, error) {
	return s.Repo.GetAdditionalExpenseByID(id)
}

func (s *AdditionalExpenseService) CreateAdditionalExpense(expense *models.AdditionalExpense) (*models.AdditionalExpense, error) {
	return s.Repo.CreateAdditionalExpense(expense)
}

func (s *AdditionalExpenseService) DeleteAdditionalExpense(id string) error {
	return s.Repo.DeleteAdditionalExpense(id)
}

func (s *AdditionalExpenseService) UpdateAdditionalExpense(expense *models.AdditionalExpense) (*models.AdditionalExpense, error) {
	return s.Repo.UpdateAdditionalExpense(expense)
}
