package models

type Cart struct {
	ID     string `json:"id"`
	UserID string `json:"user_id"`
	ItemID string `json:"item_id"`
}

func (admin *Cart) Validate() ([]error, bool) {
	return nil, true
}
