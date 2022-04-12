package controllers

import (
	"stores/db"
	"stores/models"
	"stores/views"

	"github.com/gofiber/fiber/v2"
)

func BanMerchant(c *fiber.Ctx) error {
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
	_, err = db.DB.Exec(`UPDATE merchants SET status = $1 WHERE id = $2;`, models.MerchantStatusBanned, id)
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
			Message: "Merchant banned successfuly",
		},
	}

	return c.JSON(response)
}

func BanUser(c *fiber.Ctx) error {
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
	_, err = db.DB.Exec(`UPDATE users SET status = $1 WHERE id = $2;`, models.UserStatusBanned, id)
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
			Message: "User banned successfuly",
		},
	}

	return c.JSON(response)
}
