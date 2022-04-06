package models

import (
	"time"
)

type Transaction struct {
	ID        string    `json:"id"`
	WalletID  string    `json:"wallet_id"`
	Amount    float64   `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (admin *Transaction) Validate() ([]error, bool) {
	return nil, true
}
