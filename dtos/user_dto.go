package dtos

type GetUserDTO struct {
	ID          int    `json:"id"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	UserTypeID  int    `json:"user_type"`
	UserStateID int    `json:"user_state"`
}

type UpdateUserDTO struct {
	Email       string `json:"email"`
	Password    string `json:"password"`
	UserTypeID  int    `json:"user_type"`
	UserStateID int    `json:"user_state"`
}

type CreateUserDTO struct {
	Email       string `json:"email"`
	Password    string `json:"password"`
	UserTypeID  int    `json:"user_type"`
	UserStateID int    `json:"user_state"`
}
