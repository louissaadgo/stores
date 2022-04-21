package controllers

import (
	"stores/db"
	"stores/models"
	"stores/views"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func CreateOrder(c *fiber.Ctx) error {
	order := models.Order{}
	for {
		order.ID = uuid.New().String()
		query := db.DB.QueryRow(`SELECT id FROM orders WHERE id = $1;`, order.ID)
		err := query.Scan(&order.ID)
		if err != nil {
			break
		}
	}
	order.UserID = c.GetRespHeader("request_user_id")
	var total float64 = 0

	rows, err := db.DB.Query(`SELECT item_id FROM carts WHERE user_id = $1;`, c.GetRespHeader("request_user_id"))
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
		if err := rows.Scan(&itemID); err != nil {
			response := models.Response{
				Type: models.TypeErrorResponse,
				Data: views.Error{
					Error: "Something went wrong please try again",
				},
			}
			c.Status(400)
			return c.JSON(response)
		}
		query := db.DB.QueryRow(`SELECT id, name, sku, price FROM items WHERE id = $1;`, itemID)
		err := query.Scan(&item.ID, &item.Name, &item.SKU, &item.Price)
		if err != nil {
			response := models.Response{
				Type: models.TypeErrorResponse,
				Data: views.Error{
					Error: "Invalid item ID",
				},
			}
			c.Status(400)
			return c.JSON(response)
		}
		total += item.Price
		item.Quantity = 1
		items = append(items, item)
	}
}
