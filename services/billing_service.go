package services

import (
	"errors"
	"strconv"
	"totesbackend/dtos"
	"totesbackend/repositories"
)

type BillingService struct {
	Repo         *repositories.ItemRepository
	DiscountRepo *repositories.DiscountTypeRepository
	TaxRepo      *repositories.TaxTypeRepository
}

func NewBillingService(repo *repositories.ItemRepository, discountRepo *repositories.DiscountTypeRepository, taxRepo *repositories.TaxTypeRepository) *BillingService {
	return &BillingService{Repo: repo, DiscountRepo: discountRepo, TaxRepo: taxRepo}
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

func (s *BillingService) CalculateTotal(discountTypesIds []string, taxTypesIds []string,
	itemsDTO []dtos.BillingItemDTO) (float64, error) {
	subtotal, err := s.CalculateSubtotal(itemsDTO)
	if err != nil {
		return 0, err
	}

	total := subtotal

	for _, discountID := range discountTypesIds {
		discount, err := s.DiscountRepo.GetDiscountTypeByID(discountID)
		if err != nil {
			return 0, errors.New("discount not found with ID: " + discountID)
		}

		if discount.IsPercentage {
			total -= (subtotal * (discount.Value / 100))
		} else {
			total -= discount.Value
		}
	}

	for _, taxID := range taxTypesIds {
		tax, err := s.TaxRepo.GetTaxTypeByID(taxID)
		if err != nil {
			return 0, errors.New("tax not found with ID: " + taxID)
		}

		if tax.IsPercentage {
			total += (subtotal * (tax.Value / 100))
		} else {
			total += tax.Value
		}
	}

	return total, nil
}
