package dtos

type GetCustomerDTO struct {
	ID               int    `json:"id"`
	CustomerName     string `json:"customerName"`
	CustomerId       string `json:"customerId"`
	IsBusiness       bool   `json:"isBusiness"`
	Address          string `json:"address,omitempty"`
	PhoneNumbers     string `json:"phoneNumbers,omitempty"`
	CustomerState    bool   `json:"customerState"`
	Email            string `json:"email"`
	LastName         string `json:"lastName"`
	IdentifierTypeID int    `json:"identifierTypeId"`
}

type CreateCustomerDTO struct {
	CustomerName     string `json:"customerName" binding:"required"`
	CustomerId       string `json:"customerId" binding:"required"`
	IsBusiness       bool   `json:"isBusiness"`
	Address          string `json:"address,omitempty"`
	PhoneNumbers     string `json:"phoneNumbers,omitempty"`
	CustomerState    bool   `json:"customerState"`
	Email            string `json:"email" binding:"required,email"`
	LastName         string `json:"lastName" binding:"required"`
	IdentifierTypeID int    `json:"identifierTypeId" binding:"required"`
}

type UpdateCustomerDTO struct {
	CustomerName     string `json:"customerName"`
	CustomerId       string `json:"customerId"`
	IsBusiness       bool   `json:"isBusiness"`
	Address          string `json:"address,omitempty"`
	PhoneNumbers     string `json:"phoneNumbers,omitempty"`
	CustomerState    bool   `json:"customerState"`
	Email            string `json:"email"`
	LastName         string `json:"lastName"`
	IdentifierTypeID int    `json:"identifierTypeId"`
}
