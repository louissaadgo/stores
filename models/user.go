package models

import (
	"time"

	"github.com/louissaadgo/go-checkif"
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

func (user *User) Validate() ([]error, bool) {
	name := checkif.StringObject{Data: user.Name}
	name.IsLongerThan(1).IsShorterThan(31)
	if name.IsInvalid {
		return name.Errors, false
	}

	password := checkif.StringObject{Data: user.Password}
	password.IsLongerThan(7).IsShorterThan(61).ContainsLowerCaseLetter().ContainsUpperCaseLetter().ContainsNumber()
	if password.IsInvalid {
		return password.Errors, false
	}

	return []error{}, true
}
