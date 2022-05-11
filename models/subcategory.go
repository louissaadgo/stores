package models

import (
	"time"

	"github.com/louissaadgo/go-checkif"
)

type SubCategory struct {
	ID         string    `json:"id"`
	Name       string    `json:"name"`
	CategoryID string    `json:"category_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func (subCategory *SubCategory) Validate() ([]error, bool) {
	name := checkif.StringObject{Data: subCategory.Name}
	name.IsLongerThan(1).IsShorterThan(31)
	if name.IsInvalid {
		return name.Errors, false
	}

	return []error{}, true
}
