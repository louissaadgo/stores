package models

import (
	"time"
)

type Link struct {
	ID        string    `json:"id"`
	StoreID   string    `json:"store_id"`
	Name      string    `json:"name"`
	URL       string    `json:"url"`
	Logo      string    `json:"logo"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (admin *Link) Validate() ([]error, bool) {
	return nil, true
}
