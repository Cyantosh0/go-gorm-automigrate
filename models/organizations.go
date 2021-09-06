package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Organizations struct {
	ID        BinaryUUID `gorm:"primary_key;type:binary(16)" json:"id"`
	Name      string     `gorm:"size:100;not null" json:"name" binding:"required,lte=100"`
	Address   string     `gorm:"size:200" json:"address" binding:"required,lte=200"`
	Phone     string     `gorm:"size:20" json:"phone" binding:"required,lte=20"`
	CreatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"createdAt"`
	UpdatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"updatedAt"`
}

func (o *Organizations) BeforeCreate(db *gorm.DB) error {
	id, err := uuid.NewRandom()
	o.ID = BinaryUUID(id)
	return err
}
