package services

import (
	"totesbackend/models"
	"totesbackend/repositories"
)

type TaxTypeService struct {
	Repo *repositories.TaxTypeRepository
}

func NewTaxTypeService(repo *repositories.TaxTypeRepository) *TaxTypeService {
	return &TaxTypeService{Repo: repo}
}

func (s *TaxTypeService) GetAllTaxTypes() ([]models.TaxType, error) {
	return s.Repo.GetAllTaxTypes()
}

func (s *TaxTypeService) GetTaxTypeByID(id string) (*models.TaxType, error) {
	return s.Repo.GetTaxTypeByID(id)
}
