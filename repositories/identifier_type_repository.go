package repositories

import (
	"totesbackend/models"

	"gorm.io/gorm"
)

type IdentifierTypeRepository struct {
	DB *gorm.DB
}

func NewIdentifierTypeRepository(db *gorm.DB) *IdentifierTypeRepository {
	return &IdentifierTypeRepository{DB: db}
}

func (r *IdentifierTypeRepository) GetAllIdentifierTypes() ([]models.IdentifierType, error) {
	var IdentifierTypes []models.IdentifierType
	err := r.DB.Find(&IdentifierTypes).Error
	if err != nil {
		return nil, err
	}
	return IdentifierTypes, nil
}

func (r *IdentifierTypeRepository) GetIdentifierTypeByID(id string) (*models.IdentifierType, error) {
	var IdentifierType models.IdentifierType
	err := r.DB.First(&IdentifierType, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &IdentifierType, nil
}
