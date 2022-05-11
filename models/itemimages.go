package models

type ItemImages struct {
	ID     string `json:"id"`
	ItemID string `json:"item_id"`
	Url    string `json:"url"`
}

func (image *ItemImages) Validate() ([]error, bool) {
	return nil, true
}
