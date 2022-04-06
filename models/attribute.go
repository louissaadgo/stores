package models

import (
	"time"
)

type Attribute struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (admin *Attribute) Validate() ([]error, bool) {
	return nil, true
}
