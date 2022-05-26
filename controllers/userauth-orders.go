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

	if !(order.CouponID == "") {
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
	}

	query := db.DB.QueryRow(`SELECT id FROM addresses WHERE id = $1;`, order.AddressID)
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

	var factor float64
	query = db.DB.QueryRow(`SELECT id, factor FROM currencies WHERE id = $1;`, order.CurrencyID)
	err = query.Scan(&order.CurrencyID, &factor)
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
	var totalDiscounted float64 = 0
	var totalPaidFromWallet float64 = 0
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

	if len(items) == 0 {
		response := models.Response{
			Type: models.TypeErrorResponse,
			Data: views.Error{
				Error: "No items in cart",
			},
		}
		c.Status(400)
		return c.JSON(response)
	}

	newRows, err := db.DB.Query(`SELECT amount FROM transactions WHERE currency_id = $1 AND user_id = $2;`, order.CurrencyID, order.UserID)
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

	var totalTransactions float64
	for newRows.Next() {
		var transaction float64
		if err := newRows.Scan(&transaction); err != nil {
			response := models.Response{
				Type: models.TypeErrorResponse,
				Data: views.Error{
					Error: "Something went wrong please try again",
				},
			}
			c.Status(400)
			return c.JSON(response)
		}
		totalTransactions += transaction
	}

	var itemOrdersWallet []models.ItemsOrder
	var itemOrdersCOD []models.ItemsOrder
	var totalItemOrders []models.ItemsOrder

	//Categorizing each item
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

		itemDB.Price = item.Price * factor
		total += itemDB.Price

		discountedPrice := itemDB.Price
		totalDiscounted += discountedPrice

		var itemOrder models.ItemsOrder
		itemOrder.ID = idDB
		itemOrder.OrderID = order.ID
		itemOrder.ItemID = itemDB.ID
		itemOrder.StoreID = itemDB.StoreID
		itemOrder.Price = itemDB.Price
		itemOrder.DiscountedPrice = discountedPrice

		var cod bool
		query := db.DB.QueryRow(`SELECT cash_on_delivery FROM stores WHERE id = $1;`, itemDB.StoreID)
		err := query.Scan(&cod)
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
		if cod {
			itemOrdersCOD = append(itemOrdersCOD, itemOrder)
		} else {
			itemOrdersWallet = append(itemOrdersWallet, itemOrder)
		}
	}

	//Checking Sufficiant wallet balance for Wallet items
	for _, item := range itemOrdersWallet {
		if totalTransactions-item.DiscountedPrice >= 0 {
			totalTransactions = totalTransactions - item.DiscountedPrice
			totalPaidFromWallet += item.DiscountedPrice
			item.Payment = models.PaymentWallet
			item.Status = models.ItemPaid
			totalItemOrders = append(totalItemOrders, item)
		} else {
			response := models.Response{
				Type: models.TypeErrorResponse,
				Data: views.Error{
					Error: "Insufficient wallet balance",
				},
			}
			c.Status(400)
			return c.JSON(response)
		}
	}

	//Checking Sufficiant wallet balance for COD items
	for _, item := range itemOrdersCOD {
		if totalTransactions-item.DiscountedPrice >= 0 {
			totalTransactions = totalTransactions - item.DiscountedPrice
			totalPaidFromWallet += item.DiscountedPrice
			item.Payment = models.PaymentWallet
			item.Status = models.ItemPaid
			totalItemOrders = append(totalItemOrders, item)
		} else {
			item.Payment = models.PaymentCOD
			item.Status = models.ItemUnpaid
			totalItemOrders = append(totalItemOrders, item)
		}
	}

	for _, itemOrder := range totalItemOrders {
		_, err = db.DB.Exec(`INSERT INTO items_order(id, order_id, item_id, store_id, price, discounted_price, payment, status)
			VALUES($1, $2, $3, $4, $5, $6, $7, $8);`, itemOrder.ID, itemOrder.OrderID, itemOrder.ItemID, itemOrder.StoreID, itemOrder.Price, itemOrder.DiscountedPrice, itemOrder.Payment, itemOrder.Status)
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
	order.TotalDiscounted = totalDiscounted

	transaction := models.Transaction{}
	transaction.UserID = order.UserID
	transaction.CurrencyID = order.CurrencyID
	transaction.CreatedAt = time.Now().UTC()
	for {
		transaction.ID = uuid.New().String()
		query := db.DB.QueryRow(`SELECT id FROM transactions WHERE id = $1;`, transaction.ID)
		err = query.Scan(&transaction.ID)
		if err != nil {
			break
		}
	}

	_, err = db.DB.Exec(`INSERT INTO transactions(id, user_id, currency_id, amount, created_at)
	VALUES($1, $2, $3, $4, $5);`, transaction.ID, transaction.UserID, transaction.CurrencyID, -totalPaidFromWallet, transaction.CreatedAt)
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
	_, err = db.DB.Exec(`INSERT INTO orders(id, status, total, coupon_id, address_id, user_id, created_at, updated_at, total_discounted, currency_id)
			VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10);`, order.ID, order.Status, order.Total, order.CouponID, order.AddressID, order.UserID, order.CreatedAt, order.UpdatedAt, order.TotalDiscounted, order.CurrencyID)
	if err != nil {
		response := models.Response{
			Type: models.TypeErrorResponse,
			Data: views.Error{
				Error: err.Error(),
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
