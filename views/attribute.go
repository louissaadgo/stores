package views

type AttributeValueResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	AttributeID string `json:"attribute_id"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type AttributeResponse struct {
	ID        string                   `json:"id"`
	Name      string                   `json:"name"`
	Values    []AttributeValueResponse `json:"values"`
	CreatedAt string                   `json:"created_at"`
	UpdatedAt string                   `json:"updated_at"`
}

type AllAttributes struct {
	Attributes []AttributeResponse `json:"attributes"`
}
