package models

type DiscountType struct {
	ID           int     `gorm:"primaryKey;autoIncrement;size:50" json:"id"`
	Name         string  `gorm:"size:100;not null" json:"name"`
	Description  string  `gorm:"size:300" json:"description,omitempty"`
	IsPercentage bool    `gorm:"not null" json:"is_percentage"`
	Value        float64 `gorm:"not null" json:"value"`
}
