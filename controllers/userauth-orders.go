package controllers

import (
	"stores/db"
	"stores/models"
	"stores/views"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func CreateOrder(c *fiber.Ctx) error {
	order := models.Order{}

	err := c.BodyParser(&order)
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

	if _, isValid := order.Validate(); !isValid {
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
		order.ID = uuid.New().String()
		query := db.DB.QueryRow(`SELECT id FROM orders WHERE id = $1;`, order.ID)
		err := query.Scan(&order.ID)
		if err != nil {
			break
		}
	}
	order.Status = models.OrderStatusProccessing
	order.UserID = c.GetRespHeader("request_user_id")

	coupon := models.Coupon{}
	query := db.DB.QueryRow(`SELECT id, value, type, max_usage, used, end_date FROM coupons WHERE code = $1;`, order.CouponID)
	err = query.Scan(&coupon.ID, &coupon.Value, &coupon.Type, &coupon.MaxUsage, &coupon.Used, &coupon.EndDate)
	if err != nil {
		response := models.Response{
			Type: models.TypeErrorResponse,
			Data: views.Error{
				Error: "Invalid coupon ID",
			},
		}
		c.Status(400)
		return c.JSON(response)
	}
	order.CouponID = coupon.ID

	if coupon.Used >= coupon.MaxUsage {
		response := models.Response{
			Type: models.TypeErrorResponse,
			Data: views.Error{
				Error: "Coupon reached limit usage",
			},
		}
		c.Status(400)
		return c.JSON(response)
	}

	if time.Now().After(coupon.EndDate) {
		response := models.Response{
			Type: models.TypeErrorResponse,
			Data: views.Error{
				Error: "Coupon Date expired",
			},
		}
		c.Status(400)
		return c.JSON(response)
	}

	_, err = db.DB.Exec(`UPDATE coupons SET used = $1 WHERE id = $2;`, coupon.Used+1, coupon.ID)
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

	query = db.DB.QueryRow(`SELECT id FROM addresses WHERE id = $1;`, order.AddressID)
	err = query.Scan(&order.AddressID)
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

	rows, err := db.DB.Query(`SELECT item_id FROM carts WHERE user_id = $1;`, order.UserID)
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

	var items []views.OrderItem
	var total float64 = 0
	var item views.OrderItem
	//Getting all the items information
	for rows.Next() {
		if err := rows.Scan(&item.ID); err != nil {
			response := models.Response{
				Type: models.TypeErrorResponse,
				Data: views.Error{
					Error: "Something went wrong please try again",
				},
			}
			c.Status(400)
			return c.JSON(response)
		}
		query := db.DB.QueryRow(`SELECT id, name, sku, price, store_id FROM items WHERE id = $1;`, item.ID)
		err := query.Scan(&item.ID, &item.Name, &item.SKU, &item.Price, &item.StoreID)
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
		items = append(items, item)
	}

	for _, itemDB := range items {
		var idDB string
		for {
			idDB = uuid.New().String()
			query := db.DB.QueryRow(`SELECT id FROM items_order WHERE id = $1;`, idDB)
			err := query.Scan(&idDB)
			if err != nil {
				break
			}
		}
		_, err = db.DB.Exec(`INSERT INTO items_order(id, order_id, item_id, store_id, price)
			VALUES($1, $2, $3, $4, $5);`, idDB, order.ID, itemDB.ID, itemDB.StoreID, itemDB.Price)
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
	}

	order.Total = total
	order.TotalDiscounted = 0
	// if coupon.Type == models.CouponTypeFixed {
	// 	order.TotalDiscounted = total - coupon.Value
	// } else if coupon.Type == models.CouponTypePercentage {
	// 	order.TotalDiscounted = total - total*coupon.Value/100
	// } else {
	// 	response := models.Response{
	// 		Type: models.TypeErrorResponse,
	// 		Data: views.Error{
	// 			Error: "Invalid coupon type",
	// 		},
	// 	}
	// 	c.Status(400)
	// 	return c.JSON(response)
	// }

	order.CreatedAt = time.Now().UTC()
	order.UpdatedAt = time.Now().UTC()
	_, err = db.DB.Exec(`INSERT INTO orders(id, status, total, coupon_id, address_id, user_id, created_at, updated_at, total_discounted)
			VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9);`, order.ID, order.Status, order.Total, order.CouponID, order.AddressID, order.UserID, order.CreatedAt, order.UpdatedAt, order.TotalDiscounted)
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

	_, err = db.DB.Exec(`DELETE FROM carts WHERE user_id = $1;`, order.UserID)
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
			Message: "Order created successfully",
		},
	}

	return c.JSON(response)
}
