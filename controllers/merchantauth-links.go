package controllers

import (
	"stores/db"
	"stores/models"
	"stores/views"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func CreateLink(c *fiber.Ctx) error {
	link := models.Link{}
	err := c.BodyParser(&link)
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

	if _, isValid := link.Validate(); !isValid {
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
		link.ID = uuid.New().String()
		query := db.DB.QueryRow(`SELECT id FROM links WHERE id = $1;`, link.ID)
		err = query.Scan(&link.ID)
		if err != nil {
			break
		}
	}

	query := db.DB.QueryRow(`SELECT merchant_id FROM stores WHERE id = $1;`, link.StoreID)
	var merchantID string
	err = query.Scan(&merchantID)
	if err != nil {
		response := models.Response{
			Type: models.TypeErrorResponse,
			Data: views.Error{
				Error: "Invalid store ID",
			},
		}
		c.Status(400)
		return c.JSON(response)
	}

	if merchantID != c.GetRespHeader("request_user_id") {
		response := models.Response{
			Type: models.TypeErrorResponse,
			Data: views.Error{
				Error: "Merchant can only add links to his own stores",
			},
		}
		c.Status(400)
		return c.JSON(response)
	}

	link.CreatedAt = time.Now().UTC()
	link.UpdatedAt = time.Now().UTC()

	_, err = db.DB.Exec(`INSERT INTO links(id, store_id, name, url, logo, created_at, updated_at)
	VALUES($1, $2, $3, $4, $5, $6, $7);`, link.ID, link.StoreID, link.Name, link.URL, link.Logo, link.CreatedAt, link.UpdatedAt)
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
			Message: "Link added successfuly",
		},
	}

	return c.JSON(response)
}

func UpdateLink(c *fiber.Ctx) error {
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
