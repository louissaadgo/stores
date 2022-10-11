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

	tokenString := models.Token{}
	err := c.BodyParser(&tokenString)
	if err != nil {
		response := models.Response{
			Type: models.TypeErrorResponse,
			Data: views.Error{
				Error: err.Error(),
			},
		}
		c.Status(400)
		return c.JSON(response)
	}

	payload, isValid := token.VerifyPasetoToken(tokenString.Token)
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

	query := db.DB.QueryRow(`SELECT token_id FROM admins WHERE id = $1;`, payload.UserID)
	var tokenID string
	err = query.Scan(&tokenID)

	if tokenID != payload.ID {
		response := models.Response{
			Type: "error",
			Data: views.Error{
				Error: "invalid token",
			},
		}
		c.Status(400)
		return c.JSON(response)
	}

	c.Set("request_user_id", fmt.Sprint(payload.UserID))

	return c.Next()
}
