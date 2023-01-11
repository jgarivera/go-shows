package tickets

import (
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func Register(db *gorm.DB, router *mux.Router) {
	db.AutoMigrate(&Ticket{})

	RegisterRoutes(router, &Handler{
		Database: db,
	})
}
