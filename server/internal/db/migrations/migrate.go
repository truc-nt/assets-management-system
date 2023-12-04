package migrations

import "gorm.io/gorm"

func MigrateUp(DB *gorm.DB) {
	MigrateUpUser(DB)
	MigrateUpAsset(DB)
}

func MigrationDown(DB *gorm.DB) {
	MigrateDownUser(DB)
	MigrateDownAsset(DB)
}
