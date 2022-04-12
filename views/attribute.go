package views

import "stores/models"

type AllAttributes struct {
	Attributes []models.Attribute `json:"attributes"`
}
