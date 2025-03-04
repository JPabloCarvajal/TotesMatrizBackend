package services

import (
	"totesbackend/models"
	"totesbackend/repositories"
)

type ItemTypeService struct {
	Repo *repositories.ItemTypeRepository
}

func NewItemTypeService(repo *repositories.ItemTypeRepository) *ItemTypeService {
	return &ItemTypeService{Repo: repo}
}

func (s *ItemTypeService) GetAllItemTypes() ([]models.ItemType, error) {
	return s.Repo.GetAllItemTypes()
}

func (s *ItemTypeService) GetItemTypeByID(id string) (*models.ItemType, error) {
	return s.Repo.GetItemTypeByID(id)
}
