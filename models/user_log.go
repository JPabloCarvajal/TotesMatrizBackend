package models

import (
	"time"
)

type UserLog struct {
	ID       int       `gorm:"primaryKey;autoIncrement;size:50" json:"id"`
	UserID   int       `gorm:"size:50;not null;index" json:"user_id"`
	Log      string    `gorm:"size:500;not null" json:"log"`
	DateTime time.Time `gorm:"not null" json:"date_time,omitempty"`
}
