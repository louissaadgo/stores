package models

import (
	"time"
)

type Cart struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	ItemIDs   []string  `json:"item_ids"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (admin *Cart) Validate() ([]error, bool) {
	return nil, true
}
