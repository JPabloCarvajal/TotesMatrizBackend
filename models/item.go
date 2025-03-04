package models

type Item struct {
	ID                 int                 `gorm:"primaryKey;autoIncrement;size:50" json:"id"`
	Name               string              `gorm:"size:255;not null" json:"name"`
	Description        string              `gorm:"size:300" json:"description,omitempty"`
	Stock              int                 `gorm:"not null" json:"stock"`
	SellingPrice       float64             `gorm:"not null" json:"selling_price"`
	PurchasePrice      float64             `gorm:"not null" json:"purchase_price"`
	ItemState          bool                `gorm:"not null" json:"item_state"`
	ItemTypeID         string              `gorm:"size:50;not null" json:"-"`
	ItemType           ItemType            `gorm:"foreignKey:ItemTypeID;references:ID" json:"item_type"`
	AdditionalExpenses []AdditionalExpense `gorm:"foreignKey:ItemID" json:"additional_expenses"`
}
