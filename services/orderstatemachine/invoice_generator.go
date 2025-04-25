package orderstatemachine

import "totesbackend/models"

type InvoiceGenerator interface {
	GetGeneratedInvoice() *models.Invoice
}
