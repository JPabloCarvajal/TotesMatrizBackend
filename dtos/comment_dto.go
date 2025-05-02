package dtos

type GetCommentDTO struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	Email    string `json:"email"`
	Phone    string `json:"phone,omitempty"`
	Comment  string `json:"comment,omitempty"`
}

type UpdateCommentDTO struct {
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	Email    string `json:"email"`
	Phone    string `json:"phone,omitempty"`
	Comment  string `json:"comment,omitempty"`
}

type CreateCommentDTO struct {
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	Email    string `json:"email"`
	Phone    string `json:"phone,omitempty"`
	Comment  string `json:"comment,omitempty"`
}
