package tickets

import (
	"net/url"
	"time"
)

type Ticket struct {
	ID          uint      `json:"id" gorm:"primarykey"`
	Name        string    `json:"name"`
	Price       float64   `json:"price"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func (t *Ticket) validate() url.Values {
	errs := url.Values{}

	if t.Name == "" {
		errs.Add("name", "The name field is required")
	}

	if t.Price == 0 {
		errs.Add("price", "The price field is required")
	}

	if t.Description == "" {
		errs.Add("description", "The description field is required")
	}

	return errs
}

func (t *Ticket) equal(t2 *Ticket) bool {
	if t.ID != t2.ID {
		return false
	}

	if t.Name != t2.Name {
		return false
	}

	if t.Price != t2.Price {
		return false
	}

	if t.Description != t2.Description {
		return false
	}

	if !t.CreatedAt.Equal(t2.CreatedAt) {
		return false
	}

	if !t.UpdatedAt.Equal(t2.UpdatedAt) {
		return false
	}

	return true
}
