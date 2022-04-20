package views

import "time"

type TransactionResponse struct {
	ID         string    `json:"id"`
	CurrencyID string    `json:"currency_id"`
	Amount     float64   `json:"amount"`
	CreatedAt  time.Time `json:"created_at"`
}

type WalletResponse struct {
	ID           string                `json:"id"`
	CurrencyID   string                `json:"currency_id"`
	Transactions []TransactionResponse `json:"transactions"`
	Total        string                `json:"total"`
}

type AllWallets struct {
	Wallets []WalletResponse `json:"wallets"`
}
