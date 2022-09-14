package models

import (
	"time"

	"github.com/louissaadgo/go-checkif"
)

type Store struct {
	ID             int       `json:"id"`
	MerchantID     int       `json:"merchant_id"`
	Name           string    `json:"name"`
	Description    string    `json:"description"`
	Phone          string    `json:"phone"`
	Location       string    `json:"location"`
	Country        string    `json:"country"`
	AccessKey      string    `json:"access_key"`
	CashOnDelivery bool      `json:"cash_on_delivery"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

func (store *Store) Validate() ([]error, bool) {
	name := checkif.StringObject{Data: store.Name}
	name.IsLongerThan(1).IsShorterThan(31)
	if name.IsInvalid {
		return name.Errors, false
	}

	description := checkif.StringObject{Data: store.Description}
	description.IsLongerThan(19).IsShorterThan(601)
	if description.IsInvalid {
		return description.Errors, false
	}

	return []error{}, true
}
