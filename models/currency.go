package models

import "github.com/louissaadgo/go-checkif"

type Currency struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Symbol string  `json:"symbol"`
	Factor float64 `json:"factor"`
}

func (currency *Currency) Validate() ([]error, bool) {
	name := checkif.StringObject{Data: currency.Name}
	name.IsLongerThan(1).IsShorterThan(31)
	if name.IsInvalid {
		return name.Errors, false
	}

	symbol := checkif.StringObject{Data: currency.Symbol}
	symbol.IsLongerThan(1).IsShorterThan(5)
	if symbol.IsInvalid {
		return symbol.Errors, false
	}

	if currency.Factor <= 0 {
		return []error{}, false
	}

	return []error{}, true
}
