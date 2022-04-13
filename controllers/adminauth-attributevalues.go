package controllers

import (
	"stores/db"
	"stores/models"
	"stores/views"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func CreateAttributeValue(c *fiber.Ctx) error {
	attributeValue := models.AttributeValue{}
	err := c.BodyParser(&attributeValue)
	if err != nil {
		response := models.Response{
			Type: models.TypeErrorResponse,
			Data: views.Error{
				Error: "Invalid Data Types",
			},
		}
		c.Status(400)
		return c.JSON(response)
	}

	if _, isValid := attributeValue.Validate(); !isValid {
		response := models.Response{
			Type: models.TypeErrorResponse,
			Data: views.Error{
				Error: "Invalid Data",
			},
		}
		c.Status(400)
		return c.JSON(response)
	}

	for {
		attributeValue.ID = uuid.New().String()
		query := db.DB.QueryRow(`SELECT id FROM attribute_values WHERE id = $1;`, attributeValue.ID)
		err = query.Scan()
		if err != nil {
			break
		}
	}

	query := db.DB.QueryRow(`SELECT id FROM attributes WHERE id = $1;`, attributeValue.AttributeID)
	var attributeID string
	err = query.Scan(&attributeID)
	if err != nil {
		response := models.Response{
			Type: models.TypeErrorResponse,
			Data: views.Error{
				Error: "Invalid attribute ID",
			},
		}
		c.Status(400)
		return c.JSON(response)
	}

	attributeValue.CreatedAt = time.Now().UTC()
	attributeValue.UpdatedAt = time.Now().UTC()

	_, err = db.DB.Exec(`INSERT INTO attribute_values(id, name, attribute_id, created_at, updated_at)
	VALUES($1, $2, $3, $4, $5);`, attributeValue.ID, attributeValue.Name, attributeValue.AttributeID, attributeValue.CreatedAt, attributeValue.UpdatedAt)
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

	response := models.Response{
		Type: models.TypeSuccessResponse,
		Data: views.Success{
			Message: "Attribute value added successfuly",
		},
	}

	return c.JSON(response)
}
