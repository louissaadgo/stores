package models

import (
	"time"
)

type AttributeValue struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	AttributeID string    `json:"attribute_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (admin *AttributeValue) Validate() ([]error, bool) {
	return nil, true
}
