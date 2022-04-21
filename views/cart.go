package views

type ItemResponse struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	SKU      string `json:"sku"`
	Price    string `json:"price"`
	Quantity string `json:"quantity"`
}

type CartResponse struct {
	Items []ItemResponse `json:"items"`
	Total float64        `json:"total"`
}
