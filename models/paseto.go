package models

import "time"

type PasetoTokenPayload struct {
	ID        string    `json:"id"`
	UserID    int       `json:"user_id"`
	UserType  string    `json:"user_type"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiresAt time.Time `json:"expires_at"`
}

type Token struct {
	Token string `json:"token"`
}
