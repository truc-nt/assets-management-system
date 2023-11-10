package migrations

import (
	"server/internal/models"

	"gorm.io/gorm"
)

func MigrateUpAsset(db *gorm.DB) {
	db.AutoMigrate(&models.Asset{})
}

func MigrateDownAsset(db *gorm.DB) {
	db.Migrator().DropTable("assets")
}
