package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jg-rivera/go-shows/config"
	"github.com/jg-rivera/go-shows/tickets"
)

func main() {
	config.LoadEnv()
	config.LoadDatabase()

	db := config.Database
	router := mux.NewRouter().StrictSlash(true)

	tickets.Register(db, router)

	http.ListenAndServe("localhost:80", router)
}
