package models

import (
	"time"
)

type SubCategory struct {
	ID         string    `json:"id"`
	Name       string    `json:"name"`
	CategoryID string    `json:"category_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func (admin *SubCategory) Validate() ([]error, bool) {
	return nil, true
}
