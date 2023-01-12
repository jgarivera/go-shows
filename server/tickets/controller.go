package tickets

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type Message struct {
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

type Handler struct {
	Database *gorm.DB
}

func (h *Handler) GetTicketById(w http.ResponseWriter, r *http.Request) {
	ticketId := mux.Vars(r)["id"]

	w.Header().Set("Content-Type", "application/json")

	if !h.doesTicketExist(ticketId) {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(Message{Message: "Ticket not found"})
		return
	}

	var ticket Ticket

	h.Database.First(&ticket, ticketId)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ticket)
}

func (h *Handler) GetTickets(w http.ResponseWriter, r *http.Request) {
	var tickets []Ticket

	h.Database.Find(&tickets)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tickets)
}

func (h *Handler) CreateTicket(w http.ResponseWriter, r *http.Request) {
	var ticket Ticket

	if err := json.NewDecoder(r.Body).Decode(&ticket); err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")

	if validErrs := ticket.validate(); len(validErrs) > 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Message{Message: "Ticket not created due to validation errors", Data: validErrs})
		return
	}

	h.Database.Create(&ticket)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(ticket)
}

func (h *Handler) UpdateTicket(w http.ResponseWriter, r *http.Request) {
	ticketId := mux.Vars(r)["id"]

	w.Header().Set("Content-Type", "application/json")

	if !h.doesTicketExist(ticketId) {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(Message{Message: "Ticket not found"})
		return
	}

	var ticket Ticket

	h.Database.First(&ticket, ticketId)
	json.NewDecoder(r.Body).Decode(&ticket)

	h.Database.Save(&ticket)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ticket)
}

func (h *Handler) DeleteTicket(w http.ResponseWriter, r *http.Request) {
	ticketId := mux.Vars(r)["id"]

	w.Header().Set("Content-Type", "application/json")

	if !h.doesTicketExist(ticketId) {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(Message{Message: "Ticket not found"})
		return
	}

	var ticket Ticket

	h.Database.Delete(&ticket, ticketId)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Message{Message: "Ticket deleted successfully"})
}

func (h *Handler) doesTicketExist(ticketId string) bool {
	var ticket Ticket

	h.Database.First(&ticket, ticketId)

	return ticket.ID != 0
}

func RegisterRoutes(router *mux.Router, handler *Handler) {
	router.HandleFunc("/api/tickets", handler.GetTickets).Methods("GET")
	router.HandleFunc("/api/tickets/{id}", handler.GetTicketById).Methods("GET")
	router.HandleFunc("/api/tickets", handler.CreateTicket).Methods("POST")
	router.HandleFunc("/api/tickets/{id}", handler.UpdateTicket).Methods("PUT")
	router.HandleFunc("/api/tickets/{id}", handler.DeleteTicket).Methods("DELETE")
}
