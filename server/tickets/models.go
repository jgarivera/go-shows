package tickets

import "gorm.io/gorm"

type Ticket struct {
	gorm.Model
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
}
