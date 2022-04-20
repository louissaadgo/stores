package controllers

import (
	"stores/db"
	"stores/models"
	"stores/views"

	"github.com/gofiber/fiber/v2"
)

func GetAllWallets(c *fiber.Ctx) error {
	userID := c.GetRespHeader("request_user_id")
	var wallets []views.WalletResponse

	rows, err := db.DB.Query(`SELECT id, name, symbol FROM currencies;`)
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

	for rows.Next() {
		currency := views.CurrencyResponse{}
		if err := rows.Scan(&currency.ID, &currency.Name, &currency.Symbol); err != nil {
			response := models.Response{
				Type: models.TypeErrorResponse,
				Data: views.Error{
					Error: "Something went wrong please try again",
				},
			}
			c.Status(400)
			return c.JSON(response)
		}

		newRows, err := db.DB.Query(`SELECT id, amount, created_at FROM transactions WHERE currency_id = $1 AND user_id = $2;`, currency.ID, userID)
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
		defer newRows.Close()
		var transactions []views.TransactionResponse
		var total float64
		for newRows.Next() {
			var transaction views.TransactionResponse
			if err := newRows.Scan(&transaction.ID, &transaction.Amount, &transaction.CreatedAt); err != nil {
				response := models.Response{
					Type: models.TypeErrorResponse,
					Data: views.Error{
						Error: "Something went wrong please try again",
					},
				}
				c.Status(400)
				return c.JSON(response)
			}
			total += transaction.Amount
			transactions = append(transactions, transaction)
		}
		wallet := views.WalletResponse{
			Currency:     currency,
			Transactions: transactions,
			Total:        total,
		}
		wallets = append(wallets, wallet)
	}

	response := models.Response{
		Type: models.TypeAllWallets,
		Data: views.AllWallets{
			Wallets: wallets,
		},
	}

	return c.JSON(response)
}
