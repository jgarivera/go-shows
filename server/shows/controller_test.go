package shows

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

func checkShowSame(t *testing.T, a *Show, b *Show) {
	if name := a.Name; name != b.Name {
		t.Error("Show name not the same", name)
	}

	if description := a.Description; description != b.Description {
		t.Error("Show description not the same", description)
	}

	if price := a.Price; price != b.Price {
		t.Error("Show price not the same", price)
	}
}

type GetShowsResponse struct {
	Message
	Data []Show `json:"data,omitempty"`
}

func TestGetEmptyShows(t *testing.T) {
	w := httptest.NewRecorder()
	r := mux.NewRouter()

	setupTest(t, r)

	r.ServeHTTP(w, httptest.NewRequest(http.MethodGet, "/api/shows", nil))

	if status := w.Result().StatusCode; status != http.StatusOK {
		t.Error("Invalid status code", status)
	}

	var response GetShowsResponse

	json.NewDecoder(w.Result().Body).Decode(&response)

	if shows := response.Data; shows == nil {
		t.Error("Invalid response", shows)
	}
}

func TestGetShows(t *testing.T) {
	w := httptest.NewRecorder()
	r := mux.NewRouter()

	repo := setupTest(t, r)

	show := Show{
		ID:          1,
		Name:        "Test",
		Price:       100.0,
		Description: "Test description",
	}

	repo.CreateShow(&show)

	r.ServeHTTP(w, httptest.NewRequest(http.MethodGet, "/api/shows", nil))

	if status := w.Result().StatusCode; status != http.StatusOK {
		t.Error("Invalid status code", status)
	}

	var response GetShowsResponse

	json.NewDecoder(w.Result().Body).Decode(&response)

	shows := response.Data

	if len(shows) == 0 {
		t.Error("No shows found", shows)
	}

	responseShow := shows[0]

	if !responseShow.equal(&show) {
		t.Error("Not the same show", responseShow)
	}
}

type GetShowResponse struct {
	Message
	Data Show `json:"data,omitempty"`
}

func TestGetShow(t *testing.T) {
	w := httptest.NewRecorder()
	r := mux.NewRouter()

	repo := setupTest(t, r)

	show := Show{
		Name:        "Test",
		Price:       100.0,
		Description: "Test description",
	}

	repo.CreateShow(&show)

	url := fmt.Sprintf("/api/shows/%d", show.ID)

	r.ServeHTTP(w, httptest.NewRequest(http.MethodGet, url, nil))

	if status := w.Result().StatusCode; status != http.StatusOK {
		t.Error("Invalid status code", status)
	}

	var response GetShowResponse

	json.NewDecoder(w.Result().Body).Decode(&response)

	responseShow := response.Data

	if !responseShow.equal(&show) {
		t.Error("Not the same show", responseShow)
	}
}

type CreateShowResponse struct {
	Message
	Data Show `json:"data,omitempty"`
}

func TestCreateShow(t *testing.T) {
	w := httptest.NewRecorder()
	r := mux.NewRouter()

	repo := setupTest(t, r)

	show := Show{
		Name:        "Test",
		Price:       100.0,
		Description: "Test description",
	}

	body, err := json.Marshal(show)

	if err != nil {
		t.Error(err)
	}

	r.ServeHTTP(w, httptest.NewRequest(http.MethodPost, "/api/shows", bytes.NewReader(body)))

	if status := w.Result().StatusCode; status != http.StatusCreated {
		t.Error("Invalid status code", status)
	}

	var response CreateShowResponse

	json.NewDecoder(w.Result().Body).Decode(&response)

	responseShow := response.Data

	checkShowSame(t, &responseShow, &show)

	savedShow, _ := repo.GetShowById(responseShow.ID)

	checkShowSame(t, savedShow, &show)
}
