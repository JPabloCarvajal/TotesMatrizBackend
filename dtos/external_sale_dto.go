package dtos

type GetExternalSaleDTO struct {
	ID            int    `json:"id"`
	ReporterName  string `json:"reporter_name"`
	ReporterID    string `json:"reporter_id"`
	ItemID        int    `json:"item_id"`
	ItemName      string `json:"item_name"`
	CustomerID    int    `json:"customer_id"`
	CustomerEmail string `json:"customer_email"`
}

type CreateExternalSaleDTO struct {
	ReporterName     string `json:"reporter_name" binding:"required"`
	ReporterID       string `json:"reporter_id" binding:"required"`
	ItemID           int    `json:"item_id" binding:"required"`
	CustomerName     string `json:"customerName" binding:"required"`
	CustomerID       string `json:"customerId" binding:"required"`
	IsBusiness       bool   `json:"isBusiness"`
	Address          string `json:"address,omitempty"`
	PhoneNumbers     string `json:"phoneNumbers,omitempty"`
	Email            string `json:"email" binding:"required,email"`
	LastName         string `json:"lastName" binding:"required"`
	IdentifierTypeID int    `json:"identifierTypeId" binding:"required"`
}
