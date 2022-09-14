package middlewares

import (
	"fmt"
	"stores/db"
	"stores/models"
	"stores/token"
	"stores/views"

	"github.com/gofiber/fiber/v2"
)

func AdminMiddleware(c *fiber.Ctx) error {

	tokenString := c.Query("token", "")

	payload, isValid := token.VerifyPasetoToken(tokenString)
	if !isValid {
		response := models.Response{
			Type: models.TypeErrorResponse,
			Data: views.Error{
				Error: "Unauthorized Paseto",
			},
		}
		c.Status(400)
		return c.JSON(response)
	}
	if payload.UserType != models.TypeAdmin {
		response := models.Response{
			Type: models.TypeErrorResponse,
			Data: views.Error{
				Error: "Unauthorized Type",
			},
		}
		c.Status(400)
		return c.JSON(response)
	}

	query := db.DB.QueryRow(`SELECT id, token_id FROM admins WHERE id = $1;`, payload.UserID)
	var tokenID string
	err := query.Scan(&payload.UserID, &tokenID)

	if err != nil || tokenID != payload.ID {
		response := models.Response{
			Type: models.TypeErrorResponse,
			Data: views.Error{
				Error: err.Error(),
			},
		}
		c.Status(400)
		return c.JSON(response)
	}

	c.Set("request_user_id", fmt.Sprint(payload.UserID))

	return c.Next()
}
