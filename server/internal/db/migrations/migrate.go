package migrations

import "gorm.io/gorm"

func MigrateUp(DB *gorm.DB) {
	MigrateUpAsset(DB)
	MigrateUpUsers(DB)
}

func MigrationDown(DB *gorm.DB) {
	MigrateDownAsset(DB)
	MigrateDownUsers(DB)
}
