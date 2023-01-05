package tickets

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jg-rivera/go-shows/config"
)

type Message struct {
	Message string `json:"message"`
}

func GetTicketById(w http.ResponseWriter, r *http.Request) {
	ticketId := mux.Vars(r)["id"]

	w.Header().Set("Content-Type", "application/json")

	if !doesTicketExist(ticketId) {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(Message{Message: "hi"})
		return
	}

	var ticket Ticket

	config.Database.First(&ticket, ticketId)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ticket)
}

func GetTickets(w http.ResponseWriter, r *http.Request) {
	var tickets []Ticket

	config.Database.Find(&tickets)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tickets)
}

func CreateTicket(w http.ResponseWriter, r *http.Request) {
	var ticket Ticket

	if err := json.NewDecoder(r.Body).Decode(&ticket); err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")

	if validErrs := ticket.validate(); len(validErrs) > 0 {
		err := map[string]interface{}{"validationError": validErrs}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}

	config.Database.Create(&ticket)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(ticket)
}

func UpdateTicket(w http.ResponseWriter, r *http.Request) {
	ticketId := mux.Vars(r)["id"]

	w.Header().Set("Content-Type", "application/json")

	if !doesTicketExist(ticketId) {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(Message{Message: "Ticket not found"})
		return
	}

	var ticket Ticket

	config.Database.First(&ticket, ticketId)
	json.NewDecoder(r.Body).Decode(&ticket)

	config.Database.Save(&ticket)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ticket)
}

func DeleteTicket(w http.ResponseWriter, r *http.Request) {
	ticketId := mux.Vars(r)["id"]

	w.Header().Set("Content-Type", "application/json")

	if !doesTicketExist(ticketId) {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(Message{Message: "Ticket not found"})
		return
	}

	var ticket Ticket

	config.Database.Delete(&ticket, ticketId)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Message{Message: "Ticket deleted successfully"})
}

func doesTicketExist(ticketId string) bool {
	var ticket Ticket

	config.Database.First(&ticket, ticketId)

	if ticket.ID == 0 {
		return false
	} else {
		return true
	}
}

func RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/api/tickets", GetTickets).Methods("GET")
	router.HandleFunc("/api/tickets/{id}", GetTicketById).Methods("GET")
	router.HandleFunc("/api/tickets", CreateTicket).Methods("POST")
	router.HandleFunc("/api/tickets/{id}", UpdateTicket).Methods("PUT")
	router.HandleFunc("/api/tickets/{id}", DeleteTicket).Methods("DELETE")
}
