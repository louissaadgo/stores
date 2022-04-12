package controllers

import (
	"stores/db"
	"stores/models"
	"stores/views"

	"github.com/gofiber/fiber/v2"
)

func GetAllAttributes(c *fiber.Ctx) error {
	var attributes []models.Attribute
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

	var attribute models.Attribute
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
