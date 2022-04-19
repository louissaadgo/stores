package controllers

import (
	"stores/db"
	"stores/models"
	"stores/views"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func AddToCart(c *fiber.Ctx) error {
	itemID := c.Params("id")
	cart := models.Cart{}

	for {
		cart.ID = uuid.New().String()
		query := db.DB.QueryRow(`SELECT id FROM carts WHERE id = $1;`, cart.ID)
		err := query.Scan(&cart.ID)
		if err != nil {
			break
		}
	}
	cart.UserID = c.GetRespHeader("request_user_id")

	query := db.DB.QueryRow(`SELECT id FROM items WHERE id = $1;`, itemID)
	err := query.Scan(&itemID)
	if err != nil {
		response := models.Response{
			Type: models.TypeErrorResponse,
			Data: views.Error{
				Error: "Invalid item ID",
			},
		}
		c.Status(400)
		return c.JSON(response)
	}

	_, err = db.DB.Exec(`INSERT INTO carts(id, user_id, item_id)
	VALUES($1, $2, $3);`, cart.ID, cart.UserID, itemID)
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
			Message: "Item added to cart successfuly",
		},
	}

	return c.JSON(response)
}

func DeleteFromCart(c *fiber.Ctx) error {
	id := c.Params("id")

	query := db.DB.QueryRow(`SELECT id, user_id FROM carts WHERE id = $1;`, id)
	var userID string
	err := query.Scan(&id, &userID)
	if err != nil {
		response := models.Response{
			Type: models.TypeErrorResponse,
			Data: views.Error{
				Error: "Invalid cart ID",
			},
		}
		c.Status(400)
		return c.JSON(response)
	}

	if userID != c.GetRespHeader("request_user_id") {
		response := models.Response{
			Type: models.TypeErrorResponse,
			Data: views.Error{
				Error: "User can only delete from his own cart",
			},
		}
		c.Status(400)
		return c.JSON(response)
	}

	_, err = db.DB.Exec(`DELETE FROM carts WHERE id = $1;`, id)
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
			Message: "Cart item deleted successfuly",
		},
	}

	return c.JSON(response)
}
