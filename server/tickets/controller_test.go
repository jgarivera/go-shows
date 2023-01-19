package tickets

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gorilla/mux"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTest(t *testing.T, r *mux.Router) *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"))

	if err != nil {
		t.Error("Failed to create db")
	}

	Register(db, r)

	return db
}

type GetTicketsResponse struct {
	Message
	Data []Ticket `json:"data,omitempty"`
}

func TestGetEmptyTickets(t *testing.T) {
	w := httptest.NewRecorder()
	r := mux.NewRouter()

	setupTest(t, r)

	r.ServeHTTP(w, httptest.NewRequest(http.MethodGet, "/api/tickets", nil))

	if status := w.Result().StatusCode; status != http.StatusOK {
		t.Error("Invalid status code", status)
	}

	var response GetTicketsResponse

	json.NewDecoder(w.Result().Body).Decode(&response)

	if tickets := response.Data; tickets == nil {
		t.Error("Invalid response", tickets)
	}
}

func TestGetTickets(t *testing.T) {
	w := httptest.NewRecorder()
	r := mux.NewRouter()

	db := setupTest(t, r)

	ticket := Ticket{
		ID:          1,
		Name:        "Test",
		Price:       100.0,
		Description: "Test description",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	db.Create(&ticket)

	r.ServeHTTP(w, httptest.NewRequest(http.MethodGet, "/api/tickets", nil))

	if status := w.Result().StatusCode; status != http.StatusOK {
		t.Error("Invalid status code", status)
	}

	var response GetTicketsResponse

	json.NewDecoder(w.Result().Body).Decode(&response)

	tickets := response.Data

	if len(tickets) == 0 {
		t.Error("No tickets found", tickets)
	}

	responseTicket := tickets[0]

	if !responseTicket.equal(&ticket) {
		t.Error("Not the same ticket", responseTicket)
	}
}
