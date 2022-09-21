package controllers

import (
	"database/sql"
	"stores/db"
	"stores/models"
	"stores/views"

	"github.com/gofiber/fiber/v2"
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

	query := db.DB.QueryRow(`SELECT name FROM currencies WHERE name = $1;`, currency.Name)
	err = query.Scan(&currency.Name)
	if err == nil || err != sql.ErrNoRows {
		response := models.Response{
			Type: "currency_already_exists",
			Data: views.Error{
				Error: "Currency name already exists",
			},
		}
		c.Status(400)
		return c.JSON(response)
	}

	_, err = db.DB.Exec(`INSERT INTO currencies(name, symbol, factor)
	VALUES($1, $2, $3);`, currency.Name, currency.Symbol, currency.Factor)
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

func UpdateCurrency(c *fiber.Ctx) error {
	id := c.Params("id")

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

	var findID int
	query := db.DB.QueryRow(`SELECT id FROM currencies WHERE id = $1;`, id)
	err = query.Scan(&findID)
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

	query = db.DB.QueryRow(`SELECT name FROM currencies WHERE name = $1;`, currency.Name)
	err = query.Scan(&currency.Name)
	if err == nil || err != sql.ErrNoRows {
		response := models.Response{
			Type: "currency_already_exists",
			Data: views.Error{
				Error: "Currency name already exists",
			},
		}
		c.Status(400)
		return c.JSON(response)
	}

	_, err = db.DB.Exec(`UPDATE currencies SET name = $1, symbol = $2, factor = $3 WHERE id = $4;`, currency.Name, currency.Symbol, currency.Factor, id)
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
			Message: "Currency updated successfuly",
		},
	}

	return c.JSON(response)
}

func GetAllCurrencies(c *fiber.Ctx) error {
	var currencies []models.Currency
	rows, err := db.DB.Query(`SELECT id, name, symbol, factor FROM currencies;`)
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

	var currency models.Currency
	for rows.Next() {
		if err := rows.Scan(&currency.ID, &currency.Name, &currency.Symbol, &currency.Factor); err != nil {
			response := models.Response{
				Type: models.TypeErrorResponse,
				Data: views.Error{
					Error: "Something went wrong please try again",
				},
			}
			c.Status(400)
			return c.JSON(response)
		}
		currencies = append(currencies, currency)
	}

	response := models.Response{
		Type: models.TypeAllCurrencies,
		Data: views.AllCurrencies{
			Currencies: currencies,
		},
	}

	return c.JSON(response)
}
