package shows

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

func (h *Handler) GetShowById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.ParseUint(vars["id"], 10, 32)

	if err != nil {
		log.Fatal(err)
		respondWithError(w)
		return
	}

	showId := uint(id)

	w.Header().Set("Content-Type", "application/json")

	t, err := h.Repository.GetShowById(showId)

	if err == sql.ErrNoRows {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(Message{Message: "Show not found"})
		return
	} else if err != nil {
		log.Fatal(err)
		respondWithError(w)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Message{Data: t})
}

func (h *Handler) GetShows(w http.ResponseWriter, r *http.Request) {
	shows, err := h.Repository.GetShows()

	if err != nil {
		log.Fatal(err)
		respondWithError(w)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Message{Data: shows})
}

func (h *Handler) CreateShow(w http.ResponseWriter, r *http.Request) {
	var t Show

	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")

	if validErrs := t.validate(); len(validErrs) > 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Message{Message: "Show not created due to validation errors", Data: validErrs})
		return
	}

	_, err := h.Repository.CreateShow(&t)

	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Message{Message: http.StatusText(http.StatusInternalServerError)})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(Message{Message: "Show created successfully", Data: t})
}

func (h *Handler) UpdateShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.ParseUint(vars["id"], 10, 32)

	if err != nil {
		log.Fatal(err)
		respondWithError(w)
		return
	}

	showId := uint(id)

	w.Header().Set("Content-Type", "application/json")

	if _, err := h.Repository.DoesShowExist(showId); err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(Message{Message: "Show not found"})
		return
	}

	t, err := h.Repository.GetShowById(showId)

	if err != nil {
		log.Fatal(err)
		respondWithError(w)
		return
	}

	_, err = h.Repository.UpdateShow(t)

	if err != nil {
		log.Fatal(err)
		respondWithError(w)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Message{Message: "Show updated successfully", Data: t})
}

func (h *Handler) DeleteShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.ParseUint(vars["id"], 10, 32)

	if err != nil {
		log.Fatal(err)
		respondWithError(w)
		return
	}

	showId := uint(id)

	w.Header().Set("Content-Type", "application/json")

	if _, err := h.Repository.DoesShowExist(showId); err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(Message{Message: "Show not found"})
		return
	}

	_, err = h.Repository.DeleteShow(showId)

	if err != nil {
		log.Fatal(err)
		respondWithError(w)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Message{Message: "Show deleted successfully"})
}

func respondWithError(w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(Message{Message: http.StatusText(http.StatusInternalServerError)})
}

func RegisterRoutes(router *mux.Router, handler *Handler) {
	router.HandleFunc("/api/shows", handler.GetShows).Methods(http.MethodGet)
	router.HandleFunc("/api/shows/{id}", handler.GetShowById).Methods(http.MethodGet)
	router.HandleFunc("/api/shows", handler.CreateShow).Methods(http.MethodPost)
	router.HandleFunc("/api/shows/{id}", handler.UpdateShow).Methods(http.MethodPut)
	router.HandleFunc("/api/shows/{id}", handler.DeleteShow).Methods(http.MethodDelete)
}
