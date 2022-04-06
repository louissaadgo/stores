package models

import (
	"time"
)

type Review struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	StoreID   string    `json:"store_id"`
	OrderID   string    `json:"order_id"`
	Rating    int       `json:"rating"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (admin *Review) Validate() ([]error, bool) {
	return nil, true
}
