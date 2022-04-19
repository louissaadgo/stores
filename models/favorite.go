package models

type Favorite struct {
	ID     string `json:"id"`
	UserID string `json:"user_id"`
	ItemID string `json:"item_id"`
}

func (admin *Favorite) Validate() ([]error, bool) {
	return nil, true
}
