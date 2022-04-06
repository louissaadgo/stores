package models

type Currency struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Symbol string  `json:"symbol"`
	Factor float64 `json:"factor"`
}

func (admin *Currency) Validate() ([]error, bool) {
	return nil, true
}
