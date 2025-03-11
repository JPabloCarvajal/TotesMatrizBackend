package models

type Comment struct {
	ID             int    `gorm:"primaryKey;autoIncrement" json:"id"`
	Name           string `gorm:"size:100;not null" json:"name"`
	LastName       string `gorm:"size:100;not null" json:"lastname"`
	Email          string `gorm:"size:80;not null" json:"email"`
	Phone          string `gorm:"size:50" json:"phone,omitempty"`
	ResidenceState string `gorm:"size:50" json:"residenceState,omitempty"`
	ResidenceCity  string `gorm:"size:50" json:"residenceCity,omitempty"`
	Comment        string `gorm:"size:1000" json:"comment,omitempty"`
}
