package migration

import (
	"github.com/jinzhu/gorm"

	"github.com/Cyantosh0/go-gorm-automigrate/models"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(
		models.Users{},
		models.UserSettings{},
		// models.Organizations{},
	)

	// UsersSettings: Add foreign keys
	db.Model(&models.UserSettings{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")

	// Users: Drop 'language' column
	db.Model(&models.Users{}).DropColumn("language")

	// UsersSettings: Change size from 255 to 200
	db.Model(&models.UserSettings{}).ModifyColumn("value", "varchar(200)")

	db.DropTable(&models.Organizations{}) // Optional: db.DropTable("organizations")
}
