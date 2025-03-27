package repositories

import (
	"errors"
	"totesbackend/models"

	"gorm.io/gorm"
)

type InvoiceRepository struct {
	DB *gorm.DB
}

func NewInvoiceRepository(db *gorm.DB) *InvoiceRepository {
	return &InvoiceRepository{DB: db}
}

func (r *InvoiceRepository) GetInvoiceByID(id string) (*models.Invoice, error) {
	var invoice models.Invoice
	err := r.DB.Preload("Customer").
		Preload("Items.Item").
		Preload("Discounts").
		Preload("Taxes").
		Find(&invoice).Error
	if err != nil {
		return nil, errors.New("invoice not found")
	}
	return &invoice, nil
}

func (r *InvoiceRepository) GetAllInvoices() ([]models.Invoice, error) {
	var invoices []models.Invoice
	err := r.DB.Preload("Customer").
		Preload("Items.Item").
		Preload("Discounts").
		Preload("Taxes").
		Find(&invoices).Error
	if err != nil {
		return nil, errors.New("error retrieving invoices")
	}
	return invoices, nil
}

func (r *InvoiceRepository) SearchInvoiceByID(query string) ([]models.Invoice, error) {
	var invoices []models.Invoice
	err := r.DB.Preload("Customer").Preload("Items.Item").Preload("Discounts").Preload("Taxes").
		Where("CAST(id AS TEXT) LIKE ?", query+"%").Find(&invoices).Error

	if err != nil {
		return nil, err
	}
	return invoices, nil
}

func (r *InvoiceRepository) SearchInvoiceByCustomerPersonalId(query string) ([]models.Invoice, error) {
	var invoices []models.Invoice
	err := r.DB.Preload("Customer").
		Preload("Items.Item").
		Preload("Discounts").
		Preload("Taxes").
		Joins("JOIN customers ON customers.id = invoices.customer_id").
		Where("customers.customer_id ILIKE ?", query+"%"). // ILIKE para búsqueda sin distinción de mayúsculas
		Find(&invoices).Error

	if err != nil {
		return nil, err
	}
	return invoices, nil
}
