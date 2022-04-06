package models

import (
	"time"
)

type Interest struct {
	ID          string    `json:"id"`
	UserID      string    `json:"user_id"`
	CategoryIDs []string  `json:"category_ids"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (admin *Interest) Validate() ([]error, bool) {
	return nil, true
}
