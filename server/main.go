package main

import (
	"net/http"

	"github.com/jg-rivera/go-shows/config"
)

func main() {
	config.LoadEnv()
	config.LoadDatabase()

	http.ListenAndServe("localhost:80", nil)
}
