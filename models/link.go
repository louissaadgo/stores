package models

import (
	"time"

	"github.com/louissaadgo/go-checkif"
)

type Link struct {
	ID        string    `json:"id"`
	StoreID   string    `json:"store_id"`
	Name      string    `json:"name"`
	URL       string    `json:"url"`
	Logo      string    `json:"logo"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (link *Link) Validate() ([]error, bool) {
	name := checkif.StringObject{Data: link.Name}
	name.IsLongerThan(1).IsShorterThan(21)
	if name.IsInvalid {
		return name.Errors, false
	}

	url := checkif.StringObject{Data: link.URL}
	url.IsLongerThan(9).IsShorterThan(201)
	if url.IsInvalid {
		return url.Errors, false
	}

	return []error{}, true
}
