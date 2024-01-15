package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jgarivera/go-shows/config"
	"github.com/jgarivera/go-shows/shows"
)

func main() {
	config.LoadEnv()
	config.LoadDatabase()

	db := config.Database
	router := mux.NewRouter().StrictSlash(true)

	shows.Register(db, router)

	http.ListenAndServe("localhost:8080", router)
}
