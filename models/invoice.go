package models

import "time"

type Invoice struct {
	ID             int            `gorm:"primaryKey;autoIncrement;size:50" json:"id"`
	EnterpriseData string         `gorm:"size:300;not null" json:"enterprise_data"`
	DateTime       time.Time      `gorm:"not null" json:"date_time"`
	CustomerID     int            `gorm:"not null" json:"-"`
	Customer       Customer       `gorm:"foreignKey:CustomerID;references:ID" json:"customer"`
	Items          []Item         `gorm:"many2many:invoice_items;" json:"items"`
	Subtotal       float64        `gorm:"not null" json:"subtotal"`
	Discounts      []DiscountType `gorm:"many2many:invoice_discounts;" json:"discounts"`
	Taxes          []TaxType      `gorm:"many2many:invoice_taxes;" json:"taxes"`
	Total          float64        `gorm:"not null" json:"total"`
}
