package models

type ItemType struct {
	ID   int    `gorm:"primaryKey;autoIncrement;size:50" json:"id"`
	Name string `gorm:"size:100;not null" json:"name"`
}
