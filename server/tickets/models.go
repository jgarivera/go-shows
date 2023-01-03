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
