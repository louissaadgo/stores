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

func CreateStore(c *fiber.Ctx) error {
	store := models.Store{}
	err := c.BodyParser(&store)
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

	if _, isValid := store.Validate(); !isValid {
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
		store.ID = uuid.New().String()
		query := db.DB.QueryRow(`SELECT id FROM stores WHERE id = $1;`, store.ID)
		err = query.Scan(&store.ID)
		if err != nil {
			break
		}
	}

	query := db.DB.QueryRow(`SELECT name FROM stores WHERE name = $1;`, store.Name)
	err = query.Scan(&store.Name)
	if err == nil || err != sql.ErrNoRows {
		response := models.Response{
			Type: models.TypeErrorResponse,
			Data: views.Error{
				Error: "Store name already exists",
			},
		}
		c.Status(400)
		return c.JSON(response)
	}
	store.CreatedAt = time.Now().UTC()
	store.UpdatedAt = time.Now().UTC()
	store.AccessKey = uuid.New().String()
	store.MerchantID = c.GetRespHeader("request_user_id")

	_, err = db.DB.Exec(`INSERT INTO stores(id, merchant_id, name, description, phone, location, country, created_at, updated_at, access_key)
	VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10);`, store.ID, store.MerchantID, store.Name, store.Description, store.Phone, store.Location, store.Country, store.CreatedAt, store.UpdatedAt, store.AccessKey)
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
			Message: "Store added successfuly",
		},
	}

	return c.JSON(response)
}

func UpdateStore(c *fiber.Ctx) error {
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
