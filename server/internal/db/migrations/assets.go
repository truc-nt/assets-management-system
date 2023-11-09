package migrations

import (
	"server/internal/models"

	"gorm.io/gorm"
)

func MigrateUpAssets(db *gorm.DB) {
	db.AutoMigrate(&models.Assets{})
}

func MigrateDownAssets(db *gorm.DB) {
	db.Migrator().DropTable("assets")
}
