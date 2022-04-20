package controllers

import (
	"stores/db"
	"stores/models"
	"stores/views"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func CreateTransaction(c *fiber.Ctx) error {
	transaction := models.Transaction{}
	err := c.BodyParser(&transaction)
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

	if _, isValid := transaction.Validate(); !isValid {
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
		transaction.ID = uuid.New().String()
		query := db.DB.QueryRow(`SELECT id FROM transactions WHERE id = $1;`, transaction.ID)
		err = query.Scan(&transaction.ID)
		if err != nil {
			break
		}
	}

	query := db.DB.QueryRow(`SELECT id FROM users WHERE id = $1;`, transaction.UserID)
	err = query.Scan(&transaction.UserID)
	if err != nil {
		response := models.Response{
			Type: models.TypeErrorResponse,
			Data: views.Error{
				Error: "Invalid user ID",
			},
		}
		c.Status(400)
		return c.JSON(response)
	}

	query = db.DB.QueryRow(`SELECT id FROM currencies WHERE id = $1;`, transaction.CurrencyID)
	err = query.Scan(&transaction.CurrencyID)
	if err != nil {
		response := models.Response{
			Type: models.TypeErrorResponse,
			Data: views.Error{
				Error: "Invalid currency ID",
			},
		}
		c.Status(400)
		return c.JSON(response)
	}
	transaction.CreatedAt = time.Now().UTC()

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
