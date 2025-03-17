package models

import "time"

type Appointment struct {
	ID         int       `gorm:"primaryKey;autoIncrement" json:"id"`
	DateTime   time.Time `gorm:"not null" json:"dateTime"`
	State      bool      `gorm:"not null" json:"state"`
	CustomerID int       `gorm:"not null;index" json:"customerId"`
}
