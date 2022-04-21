package controllers

import (
	"database/sql"
	"stores/db"
	"stores/models"
	"stores/views"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func AddFavorite(c *fiber.Ctx) error {
	itemID := c.Params("id")
	favorite := models.Favorite{}

	query := db.DB.QueryRow(`SELECT id FROM favorites WHERE user_id = $1 AND item_id = $2;`, c.GetRespHeader("request_user_id"), itemID)
	err := query.Scan(&favorite.ID)
	if err == nil || err != sql.ErrNoRows {
		response := models.Response{
			Type: models.TypeErrorResponse,
			Data: views.Error{
				Error: "Already in favorites",
			},
		}
		c.Status(400)
		return c.JSON(response)
	}

	for {
		favorite.ID = uuid.New().String()
		query := db.DB.QueryRow(`SELECT id FROM favorites WHERE id = $1;`, favorite.ID)
		err := query.Scan(&favorite.ID)
		if err != nil {
			break
		}
	}
	favorite.UserID = c.GetRespHeader("request_user_id")

	query = db.DB.QueryRow(`SELECT id FROM items WHERE id = $1;`, itemID)
	err = query.Scan(&itemID)
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

	_, err = db.DB.Exec(`INSERT INTO favorites(id, user_id, item_id)
	VALUES($1, $2, $3);`, favorite.ID, favorite.UserID, itemID)
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
			Message: "Favorite added successfuly",
		},
	}

	return c.JSON(response)
}

func DeleteFavorite(c *fiber.Ctx) error {
	id := c.Params("id")

	query := db.DB.QueryRow(`SELECT id, user_id FROM favorites WHERE id = $1;`, id)
	var userID string
	err := query.Scan(&id, &userID)
	if err != nil {
		response := models.Response{
			Type: models.TypeErrorResponse,
			Data: views.Error{
				Error: "Invalid favorite ID",
			},
		}
		c.Status(400)
		return c.JSON(response)
	}

	if userID != c.GetRespHeader("request_user_id") {
		response := models.Response{
			Type: models.TypeErrorResponse,
			Data: views.Error{
				Error: "User can only delete his own favorites",
			},
		}
		c.Status(400)
		return c.JSON(response)
	}

	_, err = db.DB.Exec(`DELETE FROM favorites WHERE id = $1;`, id)
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
			Message: "Favorite deleted successfuly",
		},
	}

	return c.JSON(response)
}
