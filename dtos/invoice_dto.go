package dtos

import (
	"time"
)

type GetInoviceDTO struct {
	ID             int       `json:"id"`
	EnterpriseData string    `json:"enterprise_data"`
	DateTime       time.Time `json:"date_time"`
	CustomerID     int       `json:"-"`
	Items          []string  `json:"items"`
	Subtotal       float64   `json:"subtotal"`
	Discounts      []string  `json:"discounts"`
	Taxes          []string  `json:"taxes"`
	Total          float64   `json:"total"`
}
