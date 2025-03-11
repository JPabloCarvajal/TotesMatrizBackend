package services

import (
	"time"
	"totesbackend/models"
	"totesbackend/repositories"
)

type ItemService struct {
	Repo *repositories.ItemRepository
}

func NewItemService(repo *repositories.ItemRepository) *ItemService {
	return &ItemService{Repo: repo}
}

func (s *ItemService) GetItemByID(id string) (*models.Item, error) {
	return s.Repo.GetItemByID(id)
}

func (s *ItemService) GetAllItems() ([]models.Item, error) {
	return s.Repo.GetAllItems()
}

func (s *ItemService) SearchItemsByID(query string) ([]models.Item, error) {
	return s.Repo.SearchItemsByID(query)
}

func (s *ItemService) SearchItemsByName(query string) ([]models.Item, error) {
	return s.Repo.SearchItemsByName(query)
}

func (s *ItemService) UpdateItemState(id string, state bool) (*models.Item, error) {
	return s.Repo.UpdateItemState(id, state)
}

func (s *ItemService) UpdateItem(item *models.Item) error {
	hisRepo := repositories.NewHistoricalItemPriceRepository(s.Repo.DB)

	SellingPriceChanged, err := s.Repo.UpdateItem(item)
	if err != nil {
		return err
	}

	if !SellingPriceChanged {
		return nil
	}

	historicalPrice := models.HistoricalItemPrice{
		ItemID:     item.ID,
		Price:      item.SellingPrice,
		ModifiedAt: time.Now(),
	}

	if err := hisRepo.CreateHistoricalItemPrice(&historicalPrice); err != nil {
		return err
	}
	return nil
}

func (s *ItemService) CreateItem(item *models.Item) (*models.Item, error) {
	return s.Repo.CreateItem(item)
}
