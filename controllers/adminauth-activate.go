package controllers

import (
	"stores/db"
	"stores/models"
	"stores/views"
	"time"

	"github.com/gofiber/fiber/v2"
)

func ActivateMerchant(c *fiber.Ctx) error {
	id := c.Params("id")
	query := db.DB.QueryRow(`SELECT id FROM merchants WHERE id = $1;`, id)
	err := query.Scan(&id)
	if err != nil {
		response := models.Response{
			Type: models.TypeErrorResponse,
			Data: views.Error{
				Error: "Invalid merchant ID",
			},
		}
		c.Status(400)
		return c.JSON(response)
	}
	_, err = db.DB.Exec(`UPDATE merchants SET status = $1, updated_at = $2 WHERE id = $3;`, models.MerchantStatusActive, time.Now().UTC(), id)
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
			Message: "Merchant activated successfuly",
		},
	}

	return c.JSON(response)
}

func ActivateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	query := db.DB.QueryRow(`SELECT id FROM users WHERE id = $1;`, id)
	err := query.Scan(&id)
	if err != nil {
		response := models.Response{
			Type: models.TypeErrorResponse,
			Data: views.Error{
				Error: "Invalid user ID",
			},
		}
		c.Status(400)
		return c.JSON(response)
	}
	_, err = db.DB.Exec(`UPDATE users SET status = $1, updated_at = $2 WHERE id = $3;`, models.UserStatusActive, time.Now().UTC(), id)
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
			Message: "User activated successfuly",
		},
	}

	return c.JSON(response)
}
