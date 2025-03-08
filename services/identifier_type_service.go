package services

import (
	"totesbackend/models"
	"totesbackend/repositories"
)

type IdentifierTypeService struct {
	Repo *repositories.IdentifierTypeRepository
}

func NewIdentifierTypeService(repo *repositories.IdentifierTypeRepository) *IdentifierTypeService {
	return &IdentifierTypeService{Repo: repo}
}

func (s *IdentifierTypeService) GetAllIdentifierTypes() ([]models.IdentifierType, error) {
	return s.Repo.GetAllIdentifierTypes()
}

func (s *IdentifierTypeService) GetIdentifierTypeByID(id string) (*models.IdentifierType, error) {
	return s.Repo.GetIdentifierTypeByID(id)
}
