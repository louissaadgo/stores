package models

import (
	"time"
)

type Transaction struct {
	ID         int       `json:"id"`
	UserID     int       `json:"user_id"`
	CurrencyID int       `json:"currency_id"`
	Amount     float64   `json:"amount"`
	CreatedAt  time.Time `json:"created_at"`
}

func (transaction *Transaction) Validate() ([]error, bool) {
	return []error{}, true
}
