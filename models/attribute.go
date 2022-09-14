package models

import (
	"time"

	"github.com/louissaadgo/go-checkif"
)

type Attribute struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (attribute *Attribute) Validate() ([]error, bool) {
	name := checkif.StringObject{Data: attribute.Name}
	name.IsLongerThan(1).IsShorterThan(21)
	if name.IsInvalid {
		return name.Errors, false
	}

	return []error{}, true
}
