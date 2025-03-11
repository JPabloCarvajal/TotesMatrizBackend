package models

import "time"

type HistoricalItemPrice struct {
	ID         int       `gorm:"primaryKey;autoIncrement" json:"id"`
	ItemID     int       `gorm:"size:50;not null;index" json:"item_id"`
	Price      float64   `gorm:"not null" json:"price"`
	ModifiedAt time.Time `gorm:"not null" json:"modified_at,omitempty"`
}
