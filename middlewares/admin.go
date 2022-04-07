package middlewares

import (
	"stores/models"
	"stores/token"
	"stores/views"

	"github.com/gofiber/fiber/v2"
)

func AdminMiddleware(c *fiber.Ctx) error {
	platform := models.Platform{}
	err := c.BodyParser(&platform)

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
	var authToken string
	if platform.Platform == models.PlatformWeb {
		authToken = c.Cookies("token")
	} else if platform.Platform == models.PlatformMobile {
		authToken = platform.AuthToken
	} else {
		response := models.Response{
			Type: models.TypeErrorResponse,
			Data: views.Error{
				Error: "Invalid platform",
			},
		}
		c.Status(400)
		return c.JSON(response)
	}

	payload, isValid := token.VerifyPasetoToken(authToken)
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
	if payload.UserType != models.TypeAdmin {
		response := models.Response{
			Type: models.TypeErrorResponse,
			Data: views.Error{
				Error: "Unauthorized",
			},
		}
		c.Status(400)
		return c.JSON(response)
	}
	//check if admin in db then pass request
}
