package dtos

type UserTypeDTO struct {
	ID          uint     `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Roles       []string `json:"roles"`
}
