package tickets

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	_ "github.com/mattn/go-sqlite3"

	"github.com/gorilla/mux"
)

func setupTest(t *testing.T, r *mux.Router) Repository {
	db, err := sql.Open("sqlite3", "file::memory:?cache=shared")

	if err != nil {
		t.Error("Failed to create db")
	}

	schema, err := os.ReadFile("schema.sql")

	if err != nil {
		t.Fatalf("failed to read SQL file: %v", err)
	}

	_, err = db.Exec(string(schema))

	if err != nil {
		t.Fatalf("failed to create table: %v", err)
	}

	Register(db, r)

	return &SqlRepository{
		Database: db,
	}
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

	repo := setupTest(t, r)

	ticket := Ticket{
		ID:          1,
		Name:        "Test",
		Price:       100.0,
		Description: "Test description",
	}

	repo.CreateTicket(&ticket)

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

	repo := setupTest(t, r)

	ticket := Ticket{
		Name:        "Test",
		Price:       100.0,
		Description: "Test description",
	}

	repo.CreateTicket(&ticket)

	url := fmt.Sprintf("/api/tickets/%d", ticket.ID)

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

	repo := setupTest(t, r)

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

	savedTicket, _ := repo.GetTicketById(responseTicket.ID)

	checkTicketSame(t, savedTicket, &ticket)
}
