package dtos

type UpdateAdditionalExpenseDTO struct {
	Name        string  `json:"name"`
	ItemID      int     `json:"item_id"`
	Expense     float64 `json:"expense"`
	Description string  `json:"description,omitempty"`
}
