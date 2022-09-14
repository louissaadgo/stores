package models

type ItemImages struct {
	ID     int    `json:"id"`
	ItemID int    `json:"item_id"`
	Url    string `json:"url"`
}

func (image *ItemImages) Validate() ([]error, bool) {
	return nil, true
}
