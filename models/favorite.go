package models

type Favorite struct {
	ID     int `json:"id"`
	UserID int `json:"user_id"`
	ItemID int `json:"item_id"`
}

func (favorite *Favorite) Validate() ([]error, bool) {
	return nil, true
}
