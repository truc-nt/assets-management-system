package migrations

import (
	"server/internal/models"

	"gorm.io/gorm"
)

func MigrateUpDepartment(db *gorm.DB) {
	db.AutoMigrate(&models.Department{})
}

func MigrateDownDepartment(db *gorm.DB) {
	db.Migrator().DropTable("departments")
}
