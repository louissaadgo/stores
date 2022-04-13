package controllers

import (
	"stores/db"
	"stores/models"
	"stores/views"

	"github.com/gofiber/fiber/v2"
)

func GetAllAttributes(c *fiber.Ctx) error {
	var attributes []views.AttributeResponse
	rows, err := db.DB.Query(`SELECT id, name, created_at, updated_at FROM attributes;`)
	if err != nil {
		response := models.Response{
			Type: models.TypeErrorResponse,
			Data: views.Error{
				Error: "Something went wrong please try again",
			},
		}
		c.Status(400)
		return c.JSON(response)
	}
	defer rows.Close()

	var attribute views.AttributeResponse
	for rows.Next() {
		if err := rows.Scan(&attribute.ID, &attribute.Name, &attribute.CreatedAt, &attribute.UpdatedAt); err != nil {
			response := models.Response{
				Type: models.TypeErrorResponse,
				Data: views.Error{
					Error: "Something went wrong please try again",
				},
			}
			c.Status(400)
			return c.JSON(response)
		}
		newRows, err := db.DB.Query(`SELECT id, name, attribute_id,created_at, updated_at FROM attribute_values WHERE attribute_id = $1;`, attribute.ID)
		if err != nil {
			response := models.Response{
				Type: models.TypeErrorResponse,
				Data: views.Error{
					Error: "Something went wrong please try again",
				},
			}
			c.Status(400)
			return c.JSON(response)
		}
		defer newRows.Close()
		var attValues []views.AttributeValueResponse
		for newRows.Next() {
			var attValue views.AttributeValueResponse
			if err := newRows.Scan(&attValue.ID, &attValue.Name, &attValue.AttributeID, &attValue.CreatedAt, &attValue.UpdatedAt); err != nil {
				response := models.Response{
					Type: models.TypeErrorResponse,
					Data: views.Error{
						Error: "Something went wrong please try again",
					},
				}
				c.Status(400)
				return c.JSON(response)
			}
			attValues = append(attValues, attValue)
		}
		attribute.Values = attValues
		attributes = append(attributes, attribute)
	}

	response := models.Response{
		Type: models.TypeAllAttributes,
		Data: views.AllAttributes{
			Attributes: attributes,
		},
	}

	return c.JSON(response)
}
