package services

import (
	"totesbackend/models"
	"totesbackend/repositories"
)

type HistoricalItemPriceService struct {
	Repo *repositories.HistoricalItemPriceRepository
}

func NewHistoricalItemPriceService(repo *repositories.HistoricalItemPriceRepository) *HistoricalItemPriceService {
	return &HistoricalItemPriceService{Repo: repo}
}

// Obtener historial de precios de un ítem específico
func (s *HistoricalItemPriceService) GetHistoricalItemPrice(itemID string) ([]models.HistoricalItemPrice, error) {
	return s.Repo.GetHistoricalItemPrice(itemID)
}
