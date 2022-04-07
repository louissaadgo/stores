package controllers

import (
	"stores/db"
	"stores/models"
	"stores/token"
	"stores/views"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func AdminSignin(c *fiber.Ctx) error {
	admin := models.Admin{}
	err := c.BodyParser(&admin)

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

	if _, isValid := admin.Validate(); !isValid {
		response := models.Response{
			Type: models.TypeErrorResponse,
			Data: views.Error{
				Error: "Invalid Data",
			},
		}
		c.Status(400)
		return c.JSON(response)
	}

	password := admin.Password
	query := db.DB.QueryRow(`SELECT id, password FROM admins WHERE email = $1;`, admin.Email)
	err = query.Scan(&admin.ID, &admin.Password)
	if err != nil {
		response := models.Response{
			Type: models.TypeErrorResponse,
			Data: views.Error{
				Error: "Invalid Credentials",
			},
		}
		c.Status(400)
		return c.JSON(response)
	}

	if isValid := ValidatePassword(password, admin.Password); !isValid {
		response := models.Response{
			Type: models.TypeErrorResponse,
			Data: views.Error{
				Error: "Invalid Credentials",
			},
		}
		c.Status(400)
		return c.JSON(response)
	}

	tokenID := uuid.New().String()
	if tokenID == "" {
		response := models.Response{
			Type: models.TypeErrorResponse,
			Data: views.Error{
				Error: "Something went wrong please try again",
			},
		}
		c.Status(400)
		return c.JSON(response)
	}

	token, err := token.GeneratePasetoToken(tokenID, admin.ID, models.TypeAdmin)
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

	cookie := fiber.Cookie{
		Name:  "token",
		Value: token,
	}
	c.Cookie(&cookie)

	response := models.Response{
		Type: models.TypeAuthResponse,
		Data: views.Auth{
			AuthToken: token,
		},
	}

	return c.JSON(response)
}
