package models

import (
	"time"
)

type Address struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	Name      string    `json:"name"`
	Region    string    `json:"region"`
	City      string    `json:"city"`
	Address   string    `json:"address"`
	Longitude string    `json:"longitude"`
	Latitude  string    `json:"latitude"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (admin *Address) Validate() ([]error, bool) {
	return nil, true
}
