package migration

import (
	"github.com/jinzhu/gorm"

	"github.com/Cyantosh0/go-gorm-automigrate/models"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(
		models.Users{},
		models.UserSettings{},
		models.Organizations{},
	)

	// UsersSettings: Add foreign keys
	db.Model(&models.UserSettings{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
}
