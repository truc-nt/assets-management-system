package migrations

import (
	"server/internal/models"

	"gorm.io/gorm"
)

func MigrateUpUser(db *gorm.DB) {
	db.AutoMigrate(&models.User{})
}

func MigrateDownUser(db *gorm.DB) {
	db.Migrator().DropTable("user")
}
