package models

import (
	"time"
)

type Order struct {
	ID              string    `json:"id"`
	Status          string    `json:"status"`
	Total           float64   `json:"total"`
	TotalDiscounted float64   `json:"total_discounted"`
	UserID          string    `json:"user_id"`
	AddressID       string    `json:"address_id"`
	CouponID        string    `json:"coupon_id"`
	CurrencyID      string    `json:"currency_id"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

func (order *Order) Validate() ([]error, bool) {
	return []error{}, true
}
