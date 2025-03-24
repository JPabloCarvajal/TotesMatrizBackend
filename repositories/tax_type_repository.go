package repositories

import (
	"totesbackend/models"

	"gorm.io/gorm"
)

type TaxTypeRepository struct {
	DB *gorm.DB
}

func NewTaxTypeRepository(db *gorm.DB) *TaxTypeRepository {
	return &TaxTypeRepository{DB: db}
}

func (r *TaxTypeRepository) GetAllTaxTypes() ([]models.TaxType, error) {
	var taxTypes []models.TaxType
	err := r.DB.Find(&taxTypes).Error
	return taxTypes, err
}

func (r *TaxTypeRepository) GetTaxTypeByID(id string) (*models.TaxType, error) {
	var taxType models.TaxType
	err := r.DB.First(&taxType, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &taxType, nil
}
