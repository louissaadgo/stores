package models

type ItemsOrder struct {
	ID              int     `json:"id"`
	OrderID         int     `json:"order_id"`
	ItemID          int     `json:"item_id"`
	StoreID         int     `json:"store_id"`
	Price           float64 `json:"price"`
	DiscountedPrice float64 `json:"discounted_price"`
	Payment         string  `json:"payment"`
	Status          string  `json:"status"`
}

func (itemsOrder *ItemsOrder) Validate() ([]error, bool) {
	return nil, true
}
