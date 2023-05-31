package config

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var Database *sql.DB

func LoadDatabase() {
	dbPath := os.Getenv("DB_DATABASE")

	db, err := sql.Open("sqlite3", dbPath)

	if err != nil {
		log.Fatal("Error loading database: ", err.Error())
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	Database = db
}
