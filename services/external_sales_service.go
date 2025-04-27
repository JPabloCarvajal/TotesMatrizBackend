package services

import (
	"errors"
	"totesbackend/models"
	"totesbackend/repositories"

	"gorm.io/gorm"
)

type ExternalSaleService struct {
	Repo         *repositories.ExternalSaleRepository
	CustomerRepo *repositories.CustomerRepository
}

func NewExternalSaleService(repo *repositories.ExternalSaleRepository, customerRepo *repositories.CustomerRepository) *ExternalSaleService {
	return &ExternalSaleService{Repo: repo, CustomerRepo: customerRepo}
}

func (s *ExternalSaleService) GetExternalSaleByID(id string) (*models.ExternalSale, error) {
	return s.Repo.GetExternalSaleByID(id)
}

func (s *ExternalSaleService) GetAllExternalSales() ([]models.ExternalSale, error) {
	return s.Repo.GetAllExternalSales()
}

func (s *ExternalSaleService) CreateExternalSale(externalSale *models.ExternalSale) (*models.ExternalSale, error) {

	customer, err := s.CustomerRepo.GetCustomerByEmail(externalSale.Customer.Email)
	if err != nil {

		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}

		customer = &externalSale.Customer
		createdCustomer, err := s.CustomerRepo.CreateCustomer(customer)
		if err != nil {
			return nil, err
		}
		customer = createdCustomer
	}

	externalSale.Customer = *customer

	if err := s.Repo.CreateExternalSale(externalSale); err != nil {
		return nil, err
	}

	return externalSale, nil
}
