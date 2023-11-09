package db

import (
	"log"
	"os"
	"server/internal/db/migrations"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	dsn := os.Getenv("CONNECTION_URL")
	db_name := os.Getenv("DB_NAME")
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to db")
	}

	err = DB.Exec("CREATE DATABASE IF NOT EXISTS " + db_name).Error
	if err != nil {
		log.Fatal("Failed to create database")
	}

	dsn = dsn + db_name
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	migrations.MigrateUp(DB)
}
