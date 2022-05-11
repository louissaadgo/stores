package models

import (
	"time"

	"github.com/louissaadgo/go-checkif"
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

func (coupon *Coupon) Validate() ([]error, bool) {
	if coupon.Value < 0 {
		return []error{}, false
	}

	if coupon.MaxUsage < 0 {
		return []error{}, false
	}

	code := checkif.StringObject{Data: coupon.Code}
	code.IsLongerThan(2).IsShorterThan(16)
	if code.IsInvalid {
		return code.Errors, false
	}

	return []error{}, true
}
