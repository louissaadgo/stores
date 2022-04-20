package models

import (
	"time"
)

type Transaction struct {
	ID         string    `json:"id"`
	UserID     string    `json:"user_id"`
	CurrencyID string    `json:"currency_id"`
	Amount     float64   `json:"amount"`
	CreatedAt  time.Time `json:"created_at"`
}

func (admin *Transaction) Validate() ([]error, bool) {
	return nil, true
}
