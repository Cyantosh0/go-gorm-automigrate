package models

import "time"

type UserSettings struct {
	UserID    string    `gorm:"unique_index;size:28;not null" json:"userId"`
	Name      string    `gorm:"unique_index;size:45;not null" json:"name"`
	Value     string    `gorm:"size:200;not null" json:"value"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"createdAt"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updatedAt"`
}
