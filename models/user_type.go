package models

type UserType struct {
	ID          uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string `gorm:"size:100;not null" json:"name"`
	Description string `gorm:"size:300" json:"description,omitempty"`
	Roles       []Role `gorm:"many2many:user_type_has_role;" json:"permissions"`
}
