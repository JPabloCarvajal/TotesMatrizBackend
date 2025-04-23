package models

import "time"

type Appointment struct {
	ID               int       `gorm:"primaryKey;autoIncrement" json:"id"`
	DateTime         time.Time `gorm:"type:timestamp;not null" json:"dateTime"`
	State            bool      `gorm:"not null" json:"state"`
	CustomerID       int       `gorm:"not null;index" json:"customerId"`
	CustomerName     string    `gorm:"size:255;not null" json:"customerName"`
	IsBusiness       bool      `gorm:"not null" json:"isBusiness"`
	Address          string    `gorm:"size:100" json:"address,omitempty"`
	PhoneNumbers     string    `gorm:"size:100" json:"phoneNumbers,omitempty"`
	CustomerState    bool      `gorm:"not null" json:"customerState"`
	Email            string    `gorm:"size:255;not null" json:"email"`
	LastName         string    `gorm:"size:255;not null" json:"lastName"`
	IdentifierTypeID int       `gorm:"not null" json:"identifierTypeId"`
}
