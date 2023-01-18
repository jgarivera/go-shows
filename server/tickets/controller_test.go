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
	Data []Ticket
}

func TestGetEmptyTickets(t *testing.T) {
	w := httptest.NewRecorder()
	r := mux.NewRouter()

	setupTest(t, r)

	r.ServeHTTP(w, httptest.NewRequest("GET", "/api/tickets", nil))

	if w.Code != http.StatusOK {
		t.Error("Invalid status code", w.Code)
	}

	var response GetTicketsResponse

	json.Unmarshal(w.Body.Bytes(), &response)

	if response.Data == nil {
		t.Error("Invalid response", response.Data)
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

	r.ServeHTTP(w, httptest.NewRequest("GET", "/api/tickets", nil))

	if w.Code != http.StatusOK {
		t.Error("Invalid status code", w.Code)
	}

	var response GetTicketsResponse

	json.Unmarshal(w.Body.Bytes(), &response)

	tickets := response.Data

	if len(tickets) == 0 {
		t.Error("No tickets found", tickets)
	}

	responseTicket := tickets[0]

	if !responseTicket.equal(&ticket) {
		t.Error("Not the same ticket", responseTicket)
	}
}
