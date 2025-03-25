package services

import (
	"totesbackend/models"
	"totesbackend/repositories"
)

type InvoiceService struct {
	Repo *repositories.InvoiceRepository
}

func NewInvoiceService(repo *repositories.InvoiceRepository) *InvoiceService {
	return &InvoiceService{Repo: repo}
}

func (s *InvoiceService) GetInvoiceByID(id string) (*models.Invoice, error) {
	return s.Repo.GetInvoiceByID(id)
}

func (s *InvoiceService) GetAllInvoices() ([]models.Invoice, error) {
	return s.Repo.GetAllInvoices()
}

func (s *InvoiceService) SearchInvoiceByID(query string) ([]models.Invoice, error) {
	return s.Repo.SearchInvoiceByID(query)
}

func (s *InvoiceService) SearchInvoiceByCustomerPersonalId(query string) ([]models.Invoice, error) {
	return s.Repo.SearchInvoiceByCustomerPersonalId(query)
}
