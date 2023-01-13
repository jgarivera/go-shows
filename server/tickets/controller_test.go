package tickets

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTest(t *testing.T, r *mux.Router) {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"))

	if err != nil {
		t.Error("Failed to create db")
	}

	Register(db, r)
}

func TestGetTickets(t *testing.T) {
	w := httptest.NewRecorder()
	r := mux.NewRouter()

	setupTest(t, r)

	r.ServeHTTP(w, httptest.NewRequest("GET", "/api/tickets", nil))

	if w.Code != http.StatusOK {
		t.Error("Invalid status code", w.Code)
	}

	var response Message

	json.Unmarshal(w.Body.Bytes(), &response)

	if response.Data == nil {
		t.Error("Invalid response", response.Data)
	}
}
