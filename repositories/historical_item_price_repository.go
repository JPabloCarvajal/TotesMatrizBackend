package repositories

import (
	"totesbackend/models"

	"gorm.io/gorm"
)

type HistoricalItemPriceRepository struct {
	DB *gorm.DB
}

func NewHistoricalItemPriceRepository(db *gorm.DB) *HistoricalItemPriceRepository {
	return &HistoricalItemPriceRepository{DB: db}
}

func (r *HistoricalItemPriceRepository) CreateHistoricalItemPrice(price *models.HistoricalItemPrice) error {
	return r.DB.Create(price).Error
}

func (r *HistoricalItemPriceRepository) GetHistoricalItemPrice(itemID string) ([]models.HistoricalItemPrice, error) {
	var historicalPrices []models.HistoricalItemPrice
	err := r.DB.Where("item_id = ?", itemID).Order("added_at DESC").Find(&historicalPrices).Error
	return historicalPrices, err
}
