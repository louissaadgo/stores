package controllers

import (
	"database/sql"
	"stores/db"
	"stores/models"
	"stores/token"
	"stores/views"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func CreateAdmin(c *fiber.Ctx) error {
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

	for {
		admin.ID = uuid.New().String()
		query := db.DB.QueryRow(`SELECT id FROM admins WHERE id = $1;`, admin.ID)
		err = query.Scan()
		if err != nil {
			break
		}
	}

	query := db.DB.QueryRow(`SELECT email FROM admins WHERE email = $1;`, admin.Email)
	err = query.Scan(&admin.Email)
	if err == nil || err != sql.ErrNoRows {
		response := models.Response{
			Type: models.TypeErrorResponse,
			Data: views.Error{
				Error: "Email already exists",
			},
		}
		c.Status(400)
		return c.JSON(response)
	}

	hashedPassword, isValid := HashPassword(admin.Password)
	if !isValid {
		response := models.Response{
			Type: models.TypeErrorResponse,
			Data: views.Error{
				Error: "Something went wrong please try again",
			},
		}
		c.Status(400)
		return c.JSON(response)
	}
	admin.Password = hashedPassword
	admin.CreatedAt = time.Now().UTC()
	admin.UpdatedAt = time.Now().UTC()

	_, err = db.DB.Exec(`INSERT INTO admins(id, name, email, password, created_at, updated_at)
	VALUES($1, $2, $3, $4, $5, $6);`, admin.ID, admin.Name, admin.Email, admin.Password, admin.CreatedAt, admin.UpdatedAt)
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
