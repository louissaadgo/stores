package models

type Interest struct {
	ID         string `json:"id"`
	UserID     string `json:"user_id"`
	CategoryID string `json:"category_id"`
}

func (admin *Interest) Validate() ([]error, bool) {
	return nil, true
}
