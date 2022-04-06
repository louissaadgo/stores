package models

import (
	"time"
)

type Coupon struct {
	ID        string    `json:"id"`
	Value     float64   `json:"value"`
	Type      string    `json:"type"`
	MaxUsage  int       `json:"max_usage"`
	Used      int       `json:"used"`
	Code      string    `json:"code"`
	EndDate   time.Time `json:"end_date"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (admin *Coupon) Validate() ([]error, bool) {
	return nil, true
}
