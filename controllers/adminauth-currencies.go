package controllers

import (
	"database/sql"
	"stores/db"
	"stores/models"
	"stores/views"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func CreateCurrency(c *fiber.Ctx) error {
	currency := models.Currency{}
	err := c.BodyParser(&currency)
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

	if _, isValid := currency.Validate(); !isValid {
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
		currency.ID = uuid.New().String()
		query := db.DB.QueryRow(`SELECT id FROM currencies WHERE id = $1;`, currency.ID)
		err = query.Scan(&currency.ID)
		if err != nil {
			break
		}
	}

	query := db.DB.QueryRow(`SELECT name FROM currencies WHERE name = $1;`, currency.Name)
	err = query.Scan(&currency.Name)
	if err == nil || err != sql.ErrNoRows {
		response := models.Response{
			Type: models.TypeErrorResponse,
			Data: views.Error{
				Error: "Currency name already exists",
			},
		}
		c.Status(400)
		return c.JSON(response)
	}

	_, err = db.DB.Exec(`INSERT INTO currencies(id, name, symbol, factor)
	VALUES($1, $2, $3, $4);`, currency.ID, currency.Name, currency.Symbol, currency.Factor)
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
			Message: "Currency added successfuly",
		},
	}

	return c.JSON(response)
}
