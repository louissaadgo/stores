package models

import (
	"time"

	"github.com/louissaadgo/go-checkif"
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
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	CategoryID      string    `json:"category_id"`
	SubCategoryID   string    `json:"subcategory_id"`
	Stock           int       `json:"stock"`
	Status          string    `json:"status"`
}

func (item *Item) Validate() ([]error, bool) {
	name := checkif.StringObject{Data: item.Name}
	name.IsLongerThan(1).IsShorterThan(31)
	if name.IsInvalid {
		return name.Errors, false
	}

	sku := checkif.StringObject{Data: item.SKU}
	sku.IsLongerThan(1).IsShorterThan(31)
	if sku.IsInvalid {
		return sku.Errors, false
	}

	description := checkif.StringObject{Data: item.Description}
	description.IsLongerThan(19).IsShorterThan(51)
	if description.IsInvalid {
		return description.Errors, false
	}

	longDescription := checkif.StringObject{Data: item.LongDescription}
	longDescription.IsLongerThan(19).IsShorterThan(601)
	if longDescription.IsInvalid {
		return longDescription.Errors, false
	}

	if item.Price <= 0 {
		return []error{}, false
	}

	if item.Stock <= 0 {
		return []error{}, false
	}

	return []error{}, true
}
