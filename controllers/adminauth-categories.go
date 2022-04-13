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

func CreateCategory(c *fiber.Ctx) error {
	category := models.Category{}
	err := c.BodyParser(&category)
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

	if _, isValid := category.Validate(); !isValid {
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
		category.ID = uuid.New().String()
		query := db.DB.QueryRow(`SELECT id FROM categories WHERE id = $1;`, category.ID)
		err = query.Scan(&category.ID)
		if err != nil {
			break
		}
	}

	query := db.DB.QueryRow(`SELECT name FROM categories WHERE name = $1;`, category.Name)
	err = query.Scan(&category.Name)
	if err == nil || err != sql.ErrNoRows {
		response := models.Response{
			Type: models.TypeErrorResponse,
			Data: views.Error{
				Error: "Category name already exists",
			},
		}
		c.Status(400)
		return c.JSON(response)
	}
	category.CreatedAt = time.Now().UTC()
	category.UpdatedAt = time.Now().UTC()

	_, err = db.DB.Exec(`INSERT INTO categories(id, name, created_at, updated_at)
	VALUES($1, $2, $3, $4);`, category.ID, category.Name, category.CreatedAt, category.UpdatedAt)
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
			Message: "Category added successfuly",
		},
	}

	return c.JSON(response)
}

func UpdateCategory(c *fiber.Ctx) error {
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
