package models

type Customer struct {
	ID               int    `gorm:"primaryKey;autoIncrement" json:"id"`
	CustomerName     string `gorm:"size:255;not null" json:"customerName"`
	CustomerId       string `gorm:"size:100;not null;unique" json:"customerId"`
	IsBusiness       bool   `gorm:"not null" json:"isBusiness"`
	Address          string `gorm:"size:100" json:"address,omitempty"`
	PhoneNumbers     string `gorm:"size:100" json:"phoneNumbers,omitempty"`
	CustomerState    bool   `gorm:"not null" json:"customerState"`
	Email            string `gorm:"size:255;not null;unique" json:"email"`
	LastName         string `gorm:"size:255;not null" json:"lastName"`
	IdentifierTypeID int    `gorm:"not null" json:"identifierTypeId"`
}
