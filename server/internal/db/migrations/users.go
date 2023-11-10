package migrations

import (
	"server/internal/models"

	"gorm.io/gorm"
)

func MigrateUpUsers(db *gorm.DB) {
	db.AutoMigrate(&models.User{})
}

func MigrateDownUsers(db *gorm.DB) {
	db.Migrator().DropTable("users")
}
