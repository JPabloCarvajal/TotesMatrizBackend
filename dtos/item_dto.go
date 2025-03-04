package dtos

type GetItemDTO struct {
	ID                 int      `json:"id"`
	Name               string   `json:"name"`
	Description        string   `json:"description,omitempty"`
	Stock              int      `json:"stock"`
	SellingPrice       float64  `json:"selling_price"`
	PurchasePrice      float64  `json:"purchase_price"`
	ItemState          bool     `json:"item_state"`
	ItemTypeID         string   `json:"item_type_id"`
	AdditionalExpenses []string `json:"additional_expenses"`
}

type UpdateItemDTO struct {
	Name          int     `json:"name"`
	Description   string  `json:"description,omitempty"`
	Stock         int     `json:"stock"`
	SellingPrice  float64 `json:"selling_price"`
	PurchasePrice float64 `json:"purchase_price"`
	ItemState     bool    `json:"item_state"`
	ItemTypeID    string  `json:"item_type_id"`
}
