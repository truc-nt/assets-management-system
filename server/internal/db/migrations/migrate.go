package migrations

import "gorm.io/gorm"

func MigrateUp(DB *gorm.DB) {
	MigrateUpAssets(DB)
	MigrateUpUsers(DB)
}

func MigrationDown(DB *gorm.DB) {
	MigrateDownAssets(DB)
	MigrateDownUsers(DB)
}
