package main

import (
	"log"
	"net/http"
	"os"

	"github.com/jg-rivera/go-shows/config"
)

func main() {
	config.LoadEnv()

	log.Printf("Database: %v", os.Getenv("DATABASE"))

	http.ListenAndServe("localhost:80", nil)
}
