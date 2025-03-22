package services

import (
	"totesbackend/models"
	"totesbackend/repositories"
)

type DiscountTypeService struct {
	Repo *repositories.DiscountTypeRepository
}

func NewDiscountTypeService(repo *repositories.DiscountTypeRepository) *DiscountTypeService {
	return &DiscountTypeService{Repo: repo}
}

func (s *DiscountTypeService) GetAllDiscountTypes() ([]models.DiscountType, error) {
	return s.Repo.GetAllDiscountTypes()
}

func (s *DiscountTypeService) GetDiscountTypeByID(id string) (*models.DiscountType, error) {
	return s.Repo.GetDiscountTypeByID(id)
}
