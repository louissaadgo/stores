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

func CreateAttribute(c *fiber.Ctx) error {
	attribute := models.Attribute{}
	err := c.BodyParser(&attribute)
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

	if _, isValid := attribute.Validate(); !isValid {
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
		attribute.ID = uuid.New().String()
		query := db.DB.QueryRow(`SELECT id FROM attributes WHERE id = $1;`, attribute.ID)
		err = query.Scan()
		if err != nil {
			break
		}
	}

	query := db.DB.QueryRow(`SELECT name FROM attributes WHERE name = $1;`, attribute.Name)
	err = query.Scan(&attribute.Name)
	if err == nil || err != sql.ErrNoRows {
		response := models.Response{
			Type: models.TypeErrorResponse,
			Data: views.Error{
				Error: "Attribute name already exists",
			},
		}
		c.Status(400)
		return c.JSON(response)
	}
	attribute.CreatedAt = time.Now().UTC()
	attribute.UpdatedAt = time.Now().UTC()

	_, err = db.DB.Exec(`INSERT INTO attributes(id, name, created_at, updated_at)
	VALUES($1, $2, $3, $4);`, attribute.ID, attribute.Name, attribute.CreatedAt, attribute.UpdatedAt)
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
			Message: "Attribute added successfuly",
		},
	}

	return c.JSON(response)
}

func UpdateAttribute(c *fiber.Ctx) error {
	id := c.Params("id")

	attribute := models.Attribute{}
	err := c.BodyParser(&attribute)
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

	if _, isValid := attribute.Validate(); !isValid {
		response := models.Response{
			Type: models.TypeErrorResponse,
			Data: views.Error{
				Error: "Invalid Data",
			},
		}
		c.Status(400)
		return c.JSON(response)
	}

	query := db.DB.QueryRow(`SELECT id FROM attributes WHERE id = $1;`, id)
	err = query.Scan(&id)
	if err != nil {
		response := models.Response{
			Type: models.TypeErrorResponse,
			Data: views.Error{
				Error: "Invalid attribute ID",
			},
		}
		c.Status(400)
		return c.JSON(response)
	}

	query = db.DB.QueryRow(`SELECT name FROM attributes WHERE name = $1;`, attribute.Name)
	err = query.Scan(&attribute.Name)
	if err == nil || err != sql.ErrNoRows {
		response := models.Response{
			Type: models.TypeErrorResponse,
			Data: views.Error{
				Error: "Attribute name already exists",
			},
		}
		c.Status(400)
		return c.JSON(response)
	}

	_, err = db.DB.Exec(`UPDATE attributes SET name = $1, updated_at = $2 WHERE id = $3;`, attribute.Name, time.Now().UTC(), id)
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
			Message: "Attribute updated successfuly",
		},
	}

	return c.JSON(response)
}
