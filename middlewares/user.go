package middlewares

import (
	"fmt"
	"stores/db"
	"stores/models"
	"stores/token"
	"stores/views"

	"github.com/gofiber/fiber/v2"
)

func UserMiddleware(c *fiber.Ctx) error {

	tokenString := models.Token{}
	err := c.BodyParser(&tokenString.Token)
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

	payload, isValid := token.VerifyPasetoToken(tokenString.Token)
	if !isValid {
		response := models.Response{
			Type: models.TypeErrorResponse,
			Data: views.Error{
				Error: "Unauthorized",
			},
		}
		c.Status(400)
		return c.JSON(response)
	}
	if payload.UserType != models.TypeUser {
		response := models.Response{
			Type: models.TypeErrorResponse,
			Data: views.Error{
				Error: "Unauthorized",
			},
		}
		c.Status(400)
		return c.JSON(response)
	}

	query := db.DB.QueryRow(`SELECT id, status, token_id FROM users WHERE id = $1;`, payload.UserID)
	var userStatus string
	var tokenID string
	err = query.Scan(&payload.UserID, &userStatus, &tokenID)
	if err != nil || tokenID != payload.ID {
		response := models.Response{
			Type: models.TypeErrorResponse,
			Data: views.Error{
				Error: "Unauthorized",
			},
		}
		c.Status(400)
		return c.JSON(response)
	}
	if userStatus == models.UserStatusBanned {
		response := models.Response{
			Type: models.TypeBannedResponse,
			Data: views.Banned{
				UserType: payload.UserType,
				Status:   userStatus,
			},
		}
		c.Status(400)
		return c.JSON(response)
	}
	c.Set("request_user_id", fmt.Sprint(payload.UserID))

	return c.Next()
}
