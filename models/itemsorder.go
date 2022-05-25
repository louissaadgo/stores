package models

type ItemsOrder struct {
	ID              string  `json:"id"`
	OrderID         string  `json:"order_id"`
	ItemID          string  `json:"item_id"`
	StoreID         string  `json:"store_id"`
	Price           float64 `json:"price"`
	DiscountedPrice float64 `json:"discounted_price"`
	Payment         string  `json:"payment"`
	Status          string  `json:"status"`
}

func (itemsOrder *ItemsOrder) Validate() ([]error, bool) {
	return nil, true
}
