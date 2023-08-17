package shows

import (
	"net/url"
	"time"
)

type Show struct {
	ID          uint       `json:"id"`
	Name        string     `json:"name"`
	Price       float64    `json:"price"`
	Description string     `json:"description"`
	CreatedAt   *time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt   *time.Time `json:"updatedAt" db:"updated_at"`
}

func (s *Show) validate() url.Values {
	errs := url.Values{}

	if s.Name == "" {
		errs.Add("name", "The name field is required")
	}

	if s.Price == 0 {
		errs.Add("price", "The price field is required")
	}

	if s.Description == "" {
		errs.Add("description", "The description field is required")
	}

	return errs
}

func (s *Show) equal(s2 *Show) bool {
	if s.ID != s2.ID {
		return false
	}

	if s.Name != s2.Name {
		return false
	}

	if s.Price != s2.Price {
		return false
	}

	if s.Description != s2.Description {
		return false
	}

	if c := s.CreatedAt; c != nil && !c.Equal(*s2.CreatedAt) {
		return false
	}

	if u := s.UpdatedAt; u != nil && !u.Equal(*s2.UpdatedAt) {
		return false
	}

	return true
}
