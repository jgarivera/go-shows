package config

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading config from .env: ", err.Error())
		return
	}

	log.Println("Loaded config from .env")
}
