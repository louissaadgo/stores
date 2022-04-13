package views

import "stores/models"

type AllCoupons struct {
	Coupons []models.Coupon `json:"coupons"`
}
