package controllers

import (
	"database/sql"
	"stores/db"
	"stores/models"
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
		err = query.Scan(&admin.ID)
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

	response := models.Response{
		Type: models.TypeSuccessResponse,
		Data: views.Success{
			Message: "Admin created successfuly",
		},
	}

	return c.JSON(response)
}

func GetAllAdmins(c *fiber.Ctx) error {
	var admins []models.Admin
	rows, err := db.DB.Query(`SELECT id, name, email, created_at, updated_at FROM admins;`)
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
	defer rows.Close()

	var admin models.Admin
	for rows.Next() {
		if err := rows.Scan(&admin.ID, &admin.Name, &admin.Email, &admin.CreatedAt, &admin.UpdatedAt); err != nil {
			response := models.Response{
				Type: models.TypeErrorResponse,
				Data: views.Error{
					Error: "Something went wrong please try again",
				},
			}
			c.Status(400)
			return c.JSON(response)
		}
		admins = append(admins, admin)
	}

	response := models.Response{
		Type: models.TypeAllAdmins,
		Data: views.AllAdmins{
			Admins: admins,
		},
	}

	return c.JSON(response)
}
