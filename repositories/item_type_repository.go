package repositories

import (
	"totesbackend/models"

	"gorm.io/gorm"
)

type ItemTypeRepository struct {
	DB *gorm.DB
}

func NewItemTypeRepository(db *gorm.DB) *ItemTypeRepository {
	return &ItemTypeRepository{DB: db}
}

func (r *ItemTypeRepository) GetAllItemTypes() ([]models.ItemType, error) {
	var itemTypes []models.ItemType
	err := r.DB.Find(&itemTypes).Error
	return itemTypes, err
}

func (r *ItemTypeRepository) GetItemTypeByID(id string) (*models.ItemType, error) {
	var itemType models.ItemType
	err := r.DB.First(&itemType, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &itemType, nil
}
