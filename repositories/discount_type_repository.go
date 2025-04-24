package repositories

import (
	"totesbackend/models"

	"gorm.io/gorm"
)

type DiscountTypeRepository struct {
	DB *gorm.DB
}

func NewDiscountTypeRepository(db *gorm.DB) *DiscountTypeRepository {
	return &DiscountTypeRepository{DB: db}
}

func (r *DiscountTypeRepository) GetAllDiscountTypes() ([]models.DiscountType, error) {
	var discountTypes []models.DiscountType
	err := r.DB.Find(&discountTypes).Error
	return discountTypes, err
}

func (r *DiscountTypeRepository) GetDiscountTypeByID(id string) (*models.DiscountType, error) {
	var discountType models.DiscountType
	err := r.DB.First(&discountType, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &discountType, nil
}

func (r *DiscountTypeRepository) CreateDiscountType(discount *models.DiscountType) error {
	return r.DB.Create(discount).Error
}
