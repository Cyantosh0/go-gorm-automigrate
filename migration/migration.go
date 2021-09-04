package migration

import (
	"github.com/jinzhu/gorm"

	"github.com/Cyantosh0/go-gorm-automigrate/models"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(models.Users{})
}
