package services

import (
	"time"
	"totesbackend/models"
	"totesbackend/repositories"
)

type SalesReportService struct {
	InvoiceRepo *repositories.InvoiceRepository
}

func NewSalesReportService(invoiceRepo *repositories.InvoiceRepository) *SalesReportService {
	return &SalesReportService{
		InvoiceRepo: invoiceRepo,
	}
}

// GetInvoicesBetweenDates devuelve las facturas registradas en un rango de fechas.
func (s *SalesReportService) GetInvoicesBetweenDates(startDate, endDate time.Time) ([]models.Invoice, error) {
	return s.InvoiceRepo.GetInvoicesByDateRange(startDate, endDate)
}
