package controllers

import (
	"stores/db"
	"stores/models"
	"stores/views"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func CreateSubCategory(c *fiber.Ctx) error {
	subCategory := models.SubCategory{}
	err := c.BodyParser(&subCategory)
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

	if _, isValid := subCategory.Validate(); !isValid {
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
		subCategory.ID = uuid.New().String()
		query := db.DB.QueryRow(`SELECT id FROM subcategories WHERE id = $1;`, subCategory.ID)
		err = query.Scan(&subCategory.ID)
		if err != nil {
			break
		}
	}

	query := db.DB.QueryRow(`SELECT id FROM categories WHERE id = $1;`, subCategory.CategoryID)
	var attributeID string
	err = query.Scan(&attributeID)
	if err != nil {
		response := models.Response{
			Type: models.TypeErrorResponse,
			Data: views.Error{
				Error: "Invalid category ID",
			},
		}
		c.Status(400)
		return c.JSON(response)
	}

	subCategory.CreatedAt = time.Now().UTC()
	subCategory.UpdatedAt = time.Now().UTC()

	_, err = db.DB.Exec(`INSERT INTO subcategories(id, name, category_id, created_at, updated_at)
	VALUES($1, $2, $3, $4, $5);`, subCategory.ID, subCategory.Name, subCategory.CategoryID, subCategory.CreatedAt, subCategory.UpdatedAt)
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
			Message: "SubCategory added successfuly",
		},
	}

	return c.JSON(response)
}

func UpdateSubCategory(c *fiber.Ctx) error {
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
