package models

import (
	"time"

	"github.com/louissaadgo/go-checkif"
)

type Review struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	ItemID    string    `json:"item_id"`
	OrderID   string    `json:"order_id"`
	Rating    int       `json:"rating"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (review *Review) Validate() ([]error, bool) {
	if review.Rating < 0 || review.Rating > 5 {
		return []error{}, false
	}

	content := checkif.StringObject{Data: review.Content}
	content.IsShorterThan(361)
	if content.IsInvalid {
		return content.Errors, false
	}

	return []error{}, true
}
