package migrations

import "gorm.io/gorm"

func MigrateUp(DB *gorm.DB) {
	MigrateUpAsset(DB)
	MigrateUpDepartment(DB)
}

func MigrationDown(DB *gorm.DB) {
	MigrateDownAsset(DB)
	MigrateDownDepartment(DB)
}
