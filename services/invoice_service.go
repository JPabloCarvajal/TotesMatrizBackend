package services

import (
	"errors"
	"strconv"
	"totesbackend/dtos"
	"totesbackend/models"
	"totesbackend/repositories"
)

type InvoiceService struct {
	InvoiceRepo    *repositories.InvoiceRepository
	ItemRepo       *repositories.ItemRepository
	BillingService *BillingService
}

func NewInvoiceService(invoiceRepo *repositories.InvoiceRepository,
	itemRepo *repositories.ItemRepository, billingService *BillingService) *InvoiceService {
	return &InvoiceService{
		InvoiceRepo:    invoiceRepo,
		ItemRepo:       itemRepo,
		BillingService: billingService,
	}
}
func (s *InvoiceService) CreateInvoice(dto *dtos.CreateInvoiceDTO) (*models.Invoice, error) {
	// Verificar si hay suficiente stock para cada item
	for _, item := range dto.Items {
		itemID := strconv.Itoa(item.ID)
		hasStock, err := s.ItemRepo.HasEnoughStock(itemID, item.Stock)
		if err != nil {
			return nil, err
		}
		if !hasStock {
			return nil, errors.New("stock insuficiente para el item con ID " + itemID)
		}
	}

	// Calcular subtotal
	subtotal, err := s.BillingService.CalculateSubtotal(dto.Items)
	if err != nil {
		return nil, err
	}

	// Convertir los IDs de descuentos e impuestos a strings
	var discountIDs []string
	for _, id := range dto.Discounts {
		discountIDs = append(discountIDs, strconv.Itoa(id))
	}

	var taxIDs []string
	for _, id := range dto.Taxes {
		taxIDs = append(taxIDs, strconv.Itoa(id))
	}

	// Calcular total
	total, err := s.BillingService.CalculateTotal(discountIDs, taxIDs, dto.Items)
	if err != nil {
		return nil, err
	}

	// Crear la factura con los valores calculados
	invoice, err := s.InvoiceRepo.CreateInvoice(dto, subtotal, total)
	if err != nil {
		return nil, err
	}

	return invoice, nil
}

func (s *InvoiceService) GetInvoiceByID(id string) (*models.Invoice, error) {
	return s.InvoiceRepo.GetInvoiceByID(id)
}

func (s *InvoiceService) GetAllInvoices() ([]models.Invoice, error) {
	return s.InvoiceRepo.GetAllInvoices()
}

func (s *InvoiceService) SearchInvoiceByID(query string) ([]models.Invoice, error) {
	return s.InvoiceRepo.SearchInvoiceByID(query)
}

func (s *InvoiceService) SearchInvoiceByCustomerPersonalId(query string) ([]models.Invoice, error) {
	return s.InvoiceRepo.SearchInvoiceByCustomerPersonalId(query)
}
