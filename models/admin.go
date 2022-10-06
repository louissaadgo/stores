package models

import (
	"time"

	"github.com/louissaadgo/go-checkif"
)

type Admin struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	TokenID   string    `json:"token_id"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type AdminLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (admin *Admin) Validate() ([]error, bool) {
	email := checkif.StringObject{Data: admin.Email}
	email.IsEmail()
	if email.IsInvalid {
		return email.Errors, false
	}

	name := checkif.StringObject{Data: admin.Name}
	name.IsLongerThan(1).IsShorterThan(21)
	if name.IsInvalid {
		return name.Errors, false
	}

	password := checkif.StringObject{Data: admin.Password}
	password.IsLongerThan(6).IsShorterThan(61)
	if password.IsInvalid {
		return password.Errors, false
	}

	return []error{}, true
}
