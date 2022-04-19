package models

import (
	"time"

	"github.com/louissaadgo/go-checkif"
)

type Admin struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	TokenID   string    `json:"token_id"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (admin *Admin) Validate() ([]error, bool) {
	o := checkif.StringObject{Data: admin.Email}
	o.IsEmail()
	return o.Errors, !o.IsInvalid
}
