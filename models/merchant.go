package models

import (
	"time"

	"github.com/louissaadgo/go-checkif"
)

type Merchant struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	TokenID   string    `json:"token_id"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (merchant *Merchant) Validate() ([]error, bool) {
	name := checkif.StringObject{Data: merchant.Name}
	name.IsLongerThan(0).IsShorterThan(26)
	if name.IsInvalid {
		return name.Errors, false
	}

	email := checkif.StringObject{Data: merchant.Email}
	email.IsEmail()
	if email.IsInvalid {
		return email.Errors, false
	}

	password := checkif.StringObject{Data: merchant.Password}
	password.IsLongerThan(7).IsShorterThan(61)
	if password.IsInvalid {
		return password.Errors, false
	}

	return []error{}, true
}
