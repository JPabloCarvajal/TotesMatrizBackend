package repositories

import (
	"errors"
	"time"
	"totesbackend/dtos"
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
func (r *InvoiceRepository) CreateInvoice(dto *dtos.CreateInvoiceDTO, subtotal float64, total float64) (*models.Invoice, error) {
	invoice := &models.Invoice{
		EnterpriseData: dto.EnterpriseData,
		DateTime:       time.Now(),
		CustomerID:     dto.CustomerID,
		Subtotal:       subtotal,
		Total:          total,
	}

	tx := r.DB.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	// Restar stock de los Items
	for _, billingItem := range dto.Items {
		if err := tx.Model(&models.Item{}).
			Where("id = ?", billingItem.ID).
			UpdateColumn("stock", gorm.Expr("stock - ?", billingItem.Stock)).Error; err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	// Crear Invoice
	if err := tx.Create(invoice).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// Registrar InvoiceItems
	for _, billingItem := range dto.Items {
		invoiceItem := &models.InvoiceItem{
			InvoiceID: invoice.ID,
			ItemID:    billingItem.ID,
			Amount:    billingItem.Stock,
		}

		if err := tx.Create(invoiceItem).Error; err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	// Registrar descuentos en la relación many-to-many
	var discounts []models.DiscountType
	if len(dto.Discounts) > 0 {
		if err := tx.Where("id IN ?", dto.Discounts).Find(&discounts).Error; err != nil {
			tx.Rollback()
			return nil, err
		}
		if err := tx.Model(invoice).Association("Discounts").Append(discounts); err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	// Registrar impuestos en la relación many-to-many
	var taxes []models.TaxType
	if len(dto.Taxes) > 0 {
		if err := tx.Where("id IN ?", dto.Taxes).Find(&taxes).Error; err != nil {
			tx.Rollback()
			return nil, err
		}
		if err := tx.Model(invoice).Association("Taxes").Append(taxes); err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	// Confirmar transacción
	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	// Cargar Items con Join
	var fullInvoice models.Invoice
	if err := r.DB.
		Preload("Discounts").
		Preload("Taxes").
		Preload("Items.Item"). // Carga los items y sus productos
		First(&fullInvoice, invoice.ID).Error; err != nil {
		return nil, err
	}

	return &fullInvoice, nil
}
