package repositories

import (
	"totesbackend/models"

	"gorm.io/gorm"
)

type ExternalSaleRepository struct {
	DB *gorm.DB
}

func NewExternalSaleRepository(db *gorm.DB) *ExternalSaleRepository {
	return &ExternalSaleRepository{DB: db}
}

func (r *ExternalSaleRepository) GetExternalSaleByID(id string) (*models.ExternalSale, error) {
	var externalSale models.ExternalSale
	err := r.DB.
		Preload("Item").
		Preload("Item.ItemType").
		Preload("Item.AdditionalExpenses").
		Preload("Customer").
		First(&externalSale, "id = ?", id).Error

	if err != nil {
		return nil, err
	}

	return &externalSale, nil
}

func (r *ExternalSaleRepository) GetAllExternalSales() ([]models.ExternalSale, error) {
	var externalSales []models.ExternalSale
	err := r.DB.
		Preload("Item").
		Preload("Item.ItemType").
		Preload("Item.AdditionalExpenses").
		Preload("Customer").
		Find(&externalSales).Error

	if err != nil {
		return nil, err
	}

	return externalSales, nil
}

func (r *ExternalSaleRepository) CreateExternalSale(externalSale *models.ExternalSale) error {

	if err := r.DB.Create(externalSale).Error; err != nil {
		return err
	}
	return nil
}
