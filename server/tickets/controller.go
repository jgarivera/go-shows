package tickets

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Message struct {
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

type Handler struct {
	Repository Repository
}

func (h *Handler) GetTicketById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.ParseUint(vars["id"], 10, 32)

	if err != nil {
		log.Fatal(err)
		respondWithError(w)
		return
	}

	ticketId := uint(id)

	w.Header().Set("Content-Type", "application/json")

	t, err := h.Repository.GetTicketById(ticketId)

	if err == sql.ErrNoRows {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(Message{Message: "Ticket not found"})
		return
	} else if err != nil {
		log.Fatal(err)
		respondWithError(w)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Message{Data: t})
}

func (h *Handler) GetTickets(w http.ResponseWriter, r *http.Request) {
	tickets, err := h.Repository.GetTickets()

	if err != nil {
		log.Fatal(err)
		respondWithError(w)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Message{Data: tickets})
}

func (h *Handler) CreateTicket(w http.ResponseWriter, r *http.Request) {
	var t Ticket

	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")

	if validErrs := t.validate(); len(validErrs) > 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Message{Message: "Ticket not created due to validation errors", Data: validErrs})
		return
	}

	_, err := h.Repository.CreateTicket(&t)

	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Message{Message: http.StatusText(http.StatusInternalServerError)})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(Message{Message: "Ticket created successfully", Data: t})
}

func (h *Handler) UpdateTicket(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.ParseUint(vars["id"], 10, 32)

	if err != nil {
		log.Fatal(err)
		respondWithError(w)
		return
	}

	ticketId := uint(id)

	w.Header().Set("Content-Type", "application/json")

	if _, err := h.Repository.DoesTicketExist(ticketId); err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(Message{Message: "Ticket not found"})
		return
	}

	t, err := h.Repository.GetTicketById(ticketId)

	if err != nil {
		log.Fatal(err)
		respondWithError(w)
		return
	}

	_, err = h.Repository.UpdateTicket(t)

	if err != nil {
		log.Fatal(err)
		respondWithError(w)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Message{Message: "Ticket updated successfully", Data: t})
}

func (h *Handler) DeleteTicket(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.ParseUint(vars["id"], 10, 32)

	if err != nil {
		log.Fatal(err)
		respondWithError(w)
		return
	}

	ticketId := uint(id)

	w.Header().Set("Content-Type", "application/json")

	if _, err := h.Repository.DoesTicketExist(ticketId); err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(Message{Message: "Ticket not found"})
		return
	}

	_, err = h.Repository.DeleteTicket(ticketId)

	if err != nil {
		log.Fatal(err)
		respondWithError(w)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Message{Message: "Ticket deleted successfully"})
}

func respondWithError(w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(Message{Message: http.StatusText(http.StatusInternalServerError)})
}

func RegisterRoutes(router *mux.Router, handler *Handler) {
	router.HandleFunc("/api/tickets", handler.GetTickets).Methods(http.MethodGet)
	router.HandleFunc("/api/tickets/{id}", handler.GetTicketById).Methods(http.MethodGet)
	router.HandleFunc("/api/tickets", handler.CreateTicket).Methods(http.MethodPost)
	router.HandleFunc("/api/tickets/{id}", handler.UpdateTicket).Methods(http.MethodPut)
	router.HandleFunc("/api/tickets/{id}", handler.DeleteTicket).Methods(http.MethodDelete)
}
