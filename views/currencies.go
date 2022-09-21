package views

import "stores/models"

type AllCurrencies struct {
	Currencies []models.Currency `json:"currencies"`
}
