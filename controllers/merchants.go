package controllers

import (
	"stores/db"
	"stores/models"
	"stores/views"

	"github.com/gofiber/fiber/v2"
)

func GetAllMerchants(c *fiber.Ctx) error {
	var merchants []models.Merchant
	rows, err := db.DB.Query(`SELECT id, name, email, status, created_at, updated_at FROM merchants;`)
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

	var merchant models.Merchant
	for rows.Next() {
		if err := rows.Scan(&merchant.ID, &merchant.Name, &merchant.Email, &merchant.Status, &merchant.CreatedAt, &merchant.UpdatedAt); err != nil {
			response := models.Response{
				Type: models.TypeErrorResponse,
				Data: views.Error{
					Error: "Something went wrong please try again",
				},
			}
			c.Status(400)
			return c.JSON(response)
		}
		merchants = append(merchants, merchant)
	}

	response := models.Response{
		Type: models.TypeAllMerchants,
		Data: views.AllMerchants{
			Merchants: merchants,
		},
	}

	return c.JSON(response)
}
