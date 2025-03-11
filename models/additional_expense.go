package models

type AdditionalExpense struct {
	ID           int     `gorm:"primaryKey;autoIncrement" json:"id"`
	Name         string  `gorm:"not null;size:100" json:"name"`
	ItemID       int     `gorm:"size:50;not null;index" json:"item_id"`
	Expense      float64 `gorm:"not null" json:"expense"`
	IsPercentage bool    `gorm:"not null" json:"is_percentage"`
	Description  string  `gorm:"size:200" json:"description,omitempty"`
}
