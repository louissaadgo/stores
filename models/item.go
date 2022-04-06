package models

import (
	"time"
)

type Item struct {
	ID              string    `json:"id"`
	Name            string    `json:"name"`
	SKU             string    `json:"sku"`
	Description     string    `json:"description"`
	LongDescription string    `json:"long_description"`
	Price           float64   `json:"price"`
	Images          []string  `json:"images"`
	StoreID         string    `json:"store_id"`
	AttributeIDs    []string  `json:"attribute_ids"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	CategoryID      string    `json:"category_id"`
	SubCategoryID   string    `json:"subcategory_id"`
	Stock           int       `json:"stock"`
	Status          string    `json:"status"`
}

func (admin *Item) Validate() ([]error, bool) {
	return nil, true
}
