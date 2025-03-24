package services

import (
	"errors"
	"strconv"
	"totesbackend/dtos"
	"totesbackend/repositories"
)

type BillingService struct {
	Repo *repositories.ItemRepository
}

func (s *BillingService) CalculateSubtotal(itemsDTO []dtos.BillingItemDTO) (float64, error) {
	var subtotal float64 = 0

	for _, dto := range itemsDTO {
		item, err := s.Repo.GetItemByID(strconv.Itoa(dto.ID))
		if err != nil {
			return 0, errors.New("item not found with ID: " + strconv.Itoa(dto.ID))
		}
		subtotal += item.SellingPrice * float64(dto.Stock)
	}

	return subtotal, nil
}
