package models

import (
	"time"

	"github.com/louissaadgo/go-checkif"
)

type Category struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (category *Category) Validate() ([]error, bool) {
	name := checkif.StringObject{Data: category.Name}
	name.IsLongerThan(1).IsShorterThan(31)
	if name.IsInvalid {
		return name.Errors, false
	}

	return []error{}, true
}
