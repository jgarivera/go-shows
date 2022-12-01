package config

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Instance *gorm.DB

func LoadDatabase() {
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})

	Instance = db

	if err != nil {
		log.Fatal("Error loading database: ", err.Error())
	}

	db.AutoMigrate()
}
