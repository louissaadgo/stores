package models

import (
	"time"
)

type Store struct {
	ID          string    `json:"id"`
	MerchantID  string    `json:"merchant_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Phone       string    `json:"phone"`
	Location    string    `json:"location"`
	Country     string    `json:"country"`
	AccessKey   string    `json:"access_key"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (admin *Store) Validate() ([]error, bool) {
	return nil, true
}
