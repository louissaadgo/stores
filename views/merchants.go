package views

import "stores/models"

type AllMerchants struct {
	Merchants []models.Merchant `json:"merchants"`
}
