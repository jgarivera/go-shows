package tickets

import (
	"database/sql"

	"github.com/gorilla/mux"
)

func Register(db *sql.DB, router *mux.Router) {
	RegisterRoutes(router, &Handler{
		Database: db,
	})
}
