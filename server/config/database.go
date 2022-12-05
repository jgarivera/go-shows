package config

import (
	"log"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Database *gorm.DB

func LoadDatabase() {
	dbPath := os.Getenv("DB_DATABASE")

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})

	Database = db

	if err != nil {
		log.Fatal("Error loading database: ", err.Error())
	}
}
