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

	db.AutoMigrate(&tickets.Ticket{})

	router := mux.NewRouter().StrictSlash(true)

	tickets.RegisterRoutes(router)

	http.ListenAndServe("localhost:80", router)
}
