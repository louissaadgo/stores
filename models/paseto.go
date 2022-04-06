package models

import "time"

type PasetoTokenPayload struct {
	ID       string    `json:"id"`
	UserID   string    `json:"user_id"`
	UserType string    `json:"user_type"`
	IssuedAt time.Time `json:"issued_at"`
}
