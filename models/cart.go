package models

type Cart struct {
	ID     int `json:"id"`
	UserID int `json:"user_id"`
	ItemID int `json:"item_id"`
}

func (admin *Cart) Validate() ([]error, bool) {
	return nil, true
}
