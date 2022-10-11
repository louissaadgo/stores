package models

import (
	"time"

	"github.com/louissaadgo/go-checkif"
)

type User struct {
	ID            int       `json:"id"`
	Name          string    `json:"name"`
	Phone         string    `json:"phone"`
	Image         string    `json:"image"`
	VerifiedPhone bool      `json:"verified_phone"`
	OTP           string    `json:"otp"`
	Password      string    `json:"password"`
	TokenID       string    `json:"token_id"`
	Country       string    `json:"country"`
	Status        string    `json:"status"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
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
