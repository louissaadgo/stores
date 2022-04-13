package views

type SubCategoryResponse struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	CategoryID string `json:"category_id"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}

type CategoryResponse struct {
	ID            string                `json:"id"`
	Name          string                `json:"name"`
	SubCategories []SubCategoryResponse `json:"subcategories"`
	CreatedAt     string                `json:"created_at"`
	UpdatedAt     string                `json:"updated_at"`
}

type AllCategories struct {
	Categories []CategoryResponse `json:"categories"`
}
