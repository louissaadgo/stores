package views

import "time"

type CurrencyResponse struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
}
type TransactionResponse struct {
	ID        string    `json:"id"`
	Amount    float64   `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}

type WalletResponse struct {
	Currency     CurrencyResponse      `json:"currency"`
	Transactions []TransactionResponse `json:"transactions"`
	Total        float64               `json:"total"`
}

type AllWallets struct {
	Wallets []WalletResponse `json:"wallets"`
}
