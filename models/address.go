package models

import (
	"time"

	"github.com/louissaadgo/go-checkif"
)

type Address struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	Name      string    `json:"name"`
	Region    string    `json:"region"`
	City      string    `json:"city"`
	Address   string    `json:"address"`
	Longitude string    `json:"longitude"`
	Latitude  string    `json:"latitude"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (address *Address) Validate() ([]error, bool) {
	name := checkif.StringObject{Data: address.Name}
	name.IsLongerThan(1).IsShorterThan(31)
	if name.IsInvalid {
		return name.Errors, false
	}

	region := checkif.StringObject{Data: address.Region}
	region.IsLongerThan(1).IsShorterThan(21)
	if region.IsInvalid {
		return region.Errors, false
	}

	city := checkif.StringObject{Data: address.City}
	city.IsLongerThan(1).IsShorterThan(21)
	if city.IsInvalid {
		return city.Errors, false
	}

	addressData := checkif.StringObject{Data: address.Address}
	addressData.IsLongerThan(9).IsShorterThan(361)
	if addressData.IsInvalid {
		return addressData.Errors, false
	}

	return []error{}, true
}
