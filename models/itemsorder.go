package models

type ItemsOrder struct {
	ID      string  `json:"id"`
	OrderID string  `json:"order_id"`
	ItemID  string  `json:"item_id"`
	StoreID string  `json:"store_id"`
	Price   float64 `json:"price"`
}

func (itemsOrder *ItemsOrder) Validate() ([]error, bool) {
	return nil, true
}
