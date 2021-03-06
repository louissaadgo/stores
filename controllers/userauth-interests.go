package controllers

import (
	"database/sql"
	"stores/db"
	"stores/models"
	"stores/views"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func AddInterest(c *fiber.Ctx) error {
	categoryID := c.Params("id")
	interest := models.Interest{}

	query := db.DB.QueryRow(`SELECT id FROM interests WHERE user_id = $1 AND category_id = $2;`, c.GetRespHeader("request_user_id"), categoryID)
	err := query.Scan(&interest.ID)
	if err == nil || err != sql.ErrNoRows {
		response := models.Response{
			Type: models.TypeErrorResponse,
			Data: views.Error{
				Error: "Already in interests",
			},
		}
		c.Status(400)
		return c.JSON(response)
	}

	for {
		interest.ID = uuid.New().String()
		query := db.DB.QueryRow(`SELECT id FROM interests WHERE id = $1;`, interest.ID)
		err := query.Scan(&interest.ID)
		if err != nil {
			break
		}
	}
	interest.UserID = c.GetRespHeader("request_user_id")

	query = db.DB.QueryRow(`SELECT id FROM categories WHERE id = $1;`, categoryID)
	err = query.Scan(&categoryID)
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

	_, err = db.DB.Exec(`INSERT INTO interests(id, user_id, category_id)
	VALUES($1, $2, $3);`, interest.ID, interest.UserID, categoryID)
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
			Message: "Interest added successfuly",
		},
	}

	return c.JSON(response)
}

func DeleteInterest(c *fiber.Ctx) error {
	id := c.Params("id")

	query := db.DB.QueryRow(`SELECT id, user_id FROM interests WHERE id = $1;`, id)
	var userID string
	err := query.Scan(&id, &userID)
	if err != nil {
		response := models.Response{
			Type: models.TypeErrorResponse,
			Data: views.Error{
				Error: "Invalid interest ID",
			},
		}
		c.Status(400)
		return c.JSON(response)
	}

	if userID != c.GetRespHeader("request_user_id") {
		response := models.Response{
			Type: models.TypeErrorResponse,
			Data: views.Error{
				Error: "User can only delete his own interests",
			},
		}
		c.Status(400)
		return c.JSON(response)
	}

	_, err = db.DB.Exec(`DELETE FROM interests WHERE id = $1;`, id)
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
			Message: "Interest deleted successfuly",
		},
	}

	return c.JSON(response)
}
