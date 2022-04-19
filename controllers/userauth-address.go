package controllers

import (
	"stores/db"
	"stores/models"
	"stores/views"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func CreateAddress(c *fiber.Ctx) error {
	address := models.Address{}
	err := c.BodyParser(&address)
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

	if _, isValid := address.Validate(); !isValid {
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
		address.ID = uuid.New().String()
		query := db.DB.QueryRow(`SELECT id FROM addresses WHERE id = $1;`, address.ID)
		err = query.Scan(&address.ID)
		if err != nil {
			break
		}
	}

	address.CreatedAt = time.Now().UTC()
	address.UpdatedAt = time.Now().UTC()
	address.UserID = c.GetRespHeader("request_user_id")

	_, err = db.DB.Exec(`INSERT INTO addresses(id, user_id, name, region, city, address, longitude, latitude, created_at, updated_at)
	VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10);`, address.ID, address.UserID, address.Name, address.Region, address.City, address.Address, address.Longitude, address.Latitude, address.CreatedAt, address.UpdatedAt)
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
			Message: "Address added successfuly",
		},
	}

	return c.JSON(response)
}

func UpdateAddress(c *fiber.Ctx) error {
	id := c.Params("id")

	address := models.Address{}
	err := c.BodyParser(&address)
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

	if _, isValid := address.Validate(); !isValid {
		response := models.Response{
			Type: models.TypeErrorResponse,
			Data: views.Error{
				Error: "Invalid Data",
			},
		}
		c.Status(400)
		return c.JSON(response)
	}

	query := db.DB.QueryRow(`SELECT id, user_id FROM addresses WHERE id = $1;`, id)
	var userID string
	err = query.Scan(&id, &userID)
	if err != nil {
		response := models.Response{
			Type: models.TypeErrorResponse,
			Data: views.Error{
				Error: "Invalid address ID",
			},
		}
		c.Status(400)
		return c.JSON(response)
	}

	if userID != c.GetRespHeader("request_user_id") {
		response := models.Response{
			Type: models.TypeErrorResponse,
			Data: views.Error{
				Error: "User can only edit his own address",
			},
		}
		c.Status(400)
		return c.JSON(response)
	}

	_, err = db.DB.Exec(`UPDATE addresses SET name = $1, region = $2, city = $3, address = $4, longitude = $5, latitude = $6, updated_at = $7 WHERE id = $8;`, address.Name, address.Region, address.City, address.Address, address.Longitude, address.Latitude, time.Now().UTC(), id)
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
			Message: "Address updated successfuly",
		},
	}

	return c.JSON(response)
}
