package models

type ExternalSale struct {
	ReporterName string   `gorm:"type:varchar(255);not null" json:"reporter_name"`
	ReporterID   string   `gorm:"type:varchar(100);not null" json:"reporter_tax_id"`
	Stock        int      `gorm:"not null" json:"stock"`
	ID           int      `gorm:"primaryKey;autoIncrement" json:"id"`
	ItemID       int      `gorm:"size:50;not null" json:"-"`
	Item         Item     `gorm:"foreignKey:ItemID;references:ID" json:"item"`
	CustomerID   int      `gorm:"size:50;not null" json:"-"`
	Customer     Customer `gorm:"foreignKey:CustomerID;references:ID" json:"customer"`
}
