package tickets

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

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

func checkTicketSame(t *testing.T, a *Ticket, b *Ticket) {
	if name := a.Name; name != b.Name {
		t.Error("Ticket name not the same", name)
	}

	if description := a.Description; description != b.Description {
		t.Error("Ticket description not the same", description)
	}

	if price := a.Price; price != b.Price {
		t.Error("Ticket price not the same", price)
	}
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

type GetTicketResponse struct {
	Message
	Data Ticket `json:"data,omitempty"`
}

func TestGetTicket(t *testing.T) {
	w := httptest.NewRecorder()
	r := mux.NewRouter()

	db := setupTest(t, r)

	ticket := Ticket{
		Name:        "Test",
		Price:       100.0,
		Description: "Test description",
	}

	db.Create(&ticket)

	url := "/api/tickets/" + fmt.Sprint(ticket.ID)

	r.ServeHTTP(w, httptest.NewRequest(http.MethodGet, url, nil))

	if status := w.Result().StatusCode; status != http.StatusOK {
		t.Error("Invalid status code", status)
	}

	var response GetTicketResponse

	json.NewDecoder(w.Result().Body).Decode(&response)

	responseTicket := response.Data

	if !responseTicket.equal(&ticket) {
		t.Error("Not the same ticket", responseTicket)
	}
}

type CreateTicketResponse struct {
	Message
	Data Ticket `json:"data,omitempty"`
}

func TestCreateTicket(t *testing.T) {
	w := httptest.NewRecorder()
	r := mux.NewRouter()

	db := setupTest(t, r)

	ticket := Ticket{
		Name:        "Test",
		Price:       100.0,
		Description: "Test description",
	}

	body, err := json.Marshal(ticket)

	if err != nil {
		t.Error(err)
	}

	r.ServeHTTP(w, httptest.NewRequest(http.MethodPost, "/api/tickets", bytes.NewReader(body)))

	if status := w.Result().StatusCode; status != http.StatusCreated {
		t.Error("Invalid status code", status)
	}

	var response CreateTicketResponse

	json.NewDecoder(w.Result().Body).Decode(&response)

	responseTicket := response.Data

	checkTicketSame(t, &responseTicket, &ticket)

	var savedTicket Ticket

	db.First(&savedTicket, responseTicket.ID)

	checkTicketSame(t, &savedTicket, &ticket)
}
