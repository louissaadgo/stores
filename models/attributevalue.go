package models

import (
	"time"

	"github.com/louissaadgo/go-checkif"
)

type AttributeValue struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	AttributeID int       `json:"attribute_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (attributeValue *AttributeValue) Validate() ([]error, bool) {
	name := checkif.StringObject{Data: attributeValue.Name}
	name.IsLongerThan(0).IsShorterThan(21)
	if name.IsInvalid {
		return name.Errors, false
	}

	return []error{}, true
}
