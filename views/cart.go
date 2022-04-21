package views

type ItemResponse struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	SKU      string  `json:"sku"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}

type CartResponse struct {
	Items []ItemResponse `json:"items"`
	Total float64        `json:"total"`
}
