package views

type OrderItem struct {
	ID      string  `json:"id"`
	Name    string  `json:"name"`
	SKU     string  `json:"sku"`
	Price   float64 `json:"price"`
	StoreID string  `json:"store_id"`
}
