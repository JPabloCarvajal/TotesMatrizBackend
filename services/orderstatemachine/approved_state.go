package orderstatemachine

import (
	"errors"
	"totesbackend/config"
	"totesbackend/dtos"
	"totesbackend/models"
)

type ApprovedState struct {
	context *OrderStateMachine
	state   *models.OrderStateType
	invoice *models.Invoice
}

func NewApprovedState(context *OrderStateMachine) *ApprovedState {
	po := context.PurchaseOrder
	invoiceRepo := context.InvoiceRepo

	// Construir el DTO con los datos de la orden de compra
	var billingItems []dtos.BillingItemDTO
	for _, item := range po.Items {
		billingItems = append(billingItems, dtos.BillingItemDTO{
			ID:    item.ItemID,
			Stock: item.Amount,
		})
	}

	var discountIDs []int
	for _, d := range po.Discounts {
		discountIDs = append(discountIDs, d.ID)
	}

	var taxIDs []int
	for _, t := range po.Taxes {
		taxIDs = append(taxIDs, t.ID)
	}

	dto := &dtos.CreateInvoiceDTO{
		EnterpriseData: config.ENTERPRISE_INVOICE_DATA, // Se deja vacío por ahora
		CustomerID:     po.CustomerID,
		Items:          billingItems,
		Discounts:      discountIDs,
		Taxes:          taxIDs,
	}

	// Crear la factura usando el repositorio
	invoice, err := invoiceRepo.CreateInvoiceWithoutStockReduction(dto, po.SubTotal, po.Total)
	if err != nil {
		invoice = nil
	}

	return &ApprovedState{
		context: context,
		state: &models.OrderStateType{
			ID:          4,
			Description: "ApprovedState",
		},
		invoice: invoice,
	}
}

func (s *ApprovedState) ChangeState(stateID string) error {
	return errors.New("cannot change state: approved orders cannot transition to another state")
}

func (s *ApprovedState) GetId() int {
	return s.state.ID
}

func (s *ApprovedState) GetDescription() string {
	return s.state.Description
}

func (s *ApprovedState) GetGeneratedInvoice() *models.Invoice {
	return s.invoice
}
