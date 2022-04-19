package models

import (
	"time"
)

type User struct {
	ID             string    `json:"id"`
	Name           string    `json:"name"`
	Phone          string    `json:"phone"`
	Password       string    `json:"password"`
	SignType       string    `json:"sign_type"`
	SignID         string    `json:"sign_id"`
	TokenID        string    `json:"token_id"`
	Bday           time.Time `json:"bday"`
	Image          string    `json:"image"`
	Country        string    `json:"country"`
	Status         string    `json:"status"`
	LoyalityPoints int       `json:"loyality_points"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

func (admin *User) Validate() ([]error, bool) {
	return nil, true
}
