package controllers

import (
	"database/sql"
	"stores/db"
	"stores/models"
	"stores/views"
	"time"

	"github.com/gofiber/fiber/v2"
)

func CreateAttribute(c *fiber.Ctx) error {
	attribute := models.Attribute{}
	err := c.BodyParser(&attribute)
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

	if _, isValid := attribute.Validate(); !isValid {
		response := models.Response{
			Type: models.TypeErrorResponse,
			Data: views.Error{
				Error: "Invalid Data",
			},
		}
		c.Status(400)
		return c.JSON(response)
	}

	query := db.DB.QueryRow(`SELECT name FROM attributes WHERE name = $1;`, attribute.Name)
	err = query.Scan(&attribute.Name)
	if err == nil || err != sql.ErrNoRows {
		response := models.Response{
			Type: "attribute_name_already_exists",
			Data: views.Error{
				Error: "Attribute name already exists",
			},
		}
		c.Status(400)
		return c.JSON(response)
	}
	attribute.CreatedAt = time.Now().UTC()
	attribute.UpdatedAt = time.Now().UTC()

	_, err = db.DB.Exec(`INSERT INTO attributes(name, created_at, updated_at)
	VALUES($1, $2, $3);`, attribute.Name, attribute.CreatedAt, attribute.UpdatedAt)
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
			Message: "Attribute added successfuly",
		},
	}

	return c.JSON(response)
}

func UpdateAttribute(c *fiber.Ctx) error {
	id := c.Params("id")

	attribute := models.Attribute{}
	err := c.BodyParser(&attribute)
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

	if _, isValid := attribute.Validate(); !isValid {
		response := models.Response{
			Type: models.TypeErrorResponse,
			Data: views.Error{
				Error: "Invalid Data",
			},
		}
		c.Status(400)
		return c.JSON(response)
	}

	query := db.DB.QueryRow(`SELECT id FROM attributes WHERE id = $1;`, id)
	err = query.Scan(&id)
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

	query = db.DB.QueryRow(`SELECT name FROM attributes WHERE name = $1;`, attribute.Name)
	err = query.Scan(&attribute.Name)
	if err == nil || err != sql.ErrNoRows {
		response := models.Response{
			Type: "attribute_name_already_exists",
			Data: views.Error{
				Error: "Attribute name already exists",
			},
		}
		c.Status(400)
		return c.JSON(response)
	}

	_, err = db.DB.Exec(`UPDATE attributes SET name = $1, updated_at = $2 WHERE id = $3;`, attribute.Name, time.Now().UTC(), id)
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
			Message: "Attribute updated successfuly",
		},
	}

	return c.JSON(response)
}

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

	_, err = db.DB.Exec(`INSERT INTO attribute_values(name, attribute_id, created_at, updated_at)
	VALUES($1, $2, $3, $4);`, attributeValue.Name, attributeValue.AttributeID, attributeValue.CreatedAt, attributeValue.UpdatedAt)
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

func UpdateAttributeValue(c *fiber.Ctx) error {
	id := c.Params("id")

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

	query := db.DB.QueryRow(`SELECT id FROM attribute_values WHERE id = $1;`, id)
	err = query.Scan(&id)
	if err != nil {
		response := models.Response{
			Type: models.TypeErrorResponse,
			Data: views.Error{
				Error: "Invalid attribute value ID",
			},
		}
		c.Status(400)
		return c.JSON(response)
	}

	_, err = db.DB.Exec(`UPDATE attribute_values SET name = $1, updated_at = $2 WHERE id = $3;`, attributeValue.Name, time.Now().UTC(), id)
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
			Message: "Attribute value updated successfuly",
		},
	}

	return c.JSON(response)
}
