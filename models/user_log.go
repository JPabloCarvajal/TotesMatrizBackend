package models

import (
	"time"
)

type UserLog struct {
	ID        int       `gorm:"primaryKey;autoIncrement;size:50" json:"id"`
	UserEmail string    `gorm:"size:80;not null" json:"email"`
	Log       string    `gorm:"size:500;not null" json:"log"`
	DateTime  time.Time `gorm:"not null" json:"date_time,omitempty"`
}
