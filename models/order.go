package models

import (
	"time"
)

type Order struct {
	ID              int       `json:"id"`
	Status          string    `json:"status"`
	Total           float64   `json:"total"`
	TotalDiscounted float64   `json:"total_discounted"`
	UserID          int       `json:"user_id"`
	AddressID       int       `json:"address_id"`
	CurrencyID      int       `json:"currency_id"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

func (order *Order) Validate() ([]error, bool) {
	return []error{}, true
}
