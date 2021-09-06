package models

import (
	"time"
)

type Users struct {
	ID        string    `gorm:"primary_key;unique_index;size:28;not null" json:"id"`
	Name      string    `gorm:"size:60" json:"name"`
	Email     string    `gorm:"unique_index;size:60;not null" json:"email"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"createdAt"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updatedAt"`
}
