package models

import (
	"time"
)

type Wallet struct {
	ID         string    `json:"id"`
	UserID     string    `json:"user_id"`
	Amount     float64   `json:"amount"`
	CurrencyID string    `json:"currency_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func (admin *Wallet) Validate() ([]error, bool) {
	return nil, true
}
