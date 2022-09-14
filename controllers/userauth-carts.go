package controllers

// import (
// 	"stores/db"
// 	"stores/models"
// 	"stores/views"

// 	"github.com/gofiber/fiber/v2"
// 	"github.com/google/uuid"
// )

// func AddToCart(c *fiber.Ctx) error {
// 	itemID := c.Params("id")
// 	cart := models.Cart{}

// 	for {
// 		cart.ID = uuid.New().String()
// 		query := db.DB.QueryRow(`SELECT id FROM carts WHERE id = $1;`, cart.ID)
// 		err := query.Scan(&cart.ID)
// 		if err != nil {
// 			break
// 		}
// 	}
// 	cart.UserID = c.GetRespHeader("request_user_id")

// 	query := db.DB.QueryRow(`SELECT id FROM items WHERE id = $1;`, itemID)
// 	err := query.Scan(&itemID)
// 	if err != nil {
// 		response := models.Response{
// 			Type: models.TypeErrorResponse,
// 			Data: views.Error{
// 				Error: "Invalid item ID",
// 			},
// 		}
// 		c.Status(400)
// 		return c.JSON(response)
// 	}

// 	_, err = db.DB.Exec(`INSERT INTO carts(id, user_id, item_id)
// 	VALUES($1, $2, $3);`, cart.ID, cart.UserID, itemID)
// 	if err != nil {
// 		response := models.Response{
// 			Type: models.TypeErrorResponse,
// 			Data: views.Error{
// 				Error: "Something went wrong please try again",
// 			},
// 		}
// 		c.Status(400)
// 		return c.JSON(response)
// 	}

// 	response := models.Response{
// 		Type: models.TypeSuccessResponse,
// 		Data: views.Success{
// 			Message: "Item added to cart successfuly",
// 		},
// 	}

// 	return c.JSON(response)
// }

// func DeleteFromCart(c *fiber.Ctx) error {
// 	itemID := c.Params("id")
// 	userID := c.GetRespHeader("request_user_id")

// 	_, err := db.DB.Exec(`DELETE FROM carts WHERE id = any (SELECT id FROM carts WHERE item_id = $1 AND user_id = $2 LIMIT 1);`, itemID, userID)
// 	if err != nil {
// 		response := models.Response{
// 			Type: models.TypeErrorResponse,
// 			Data: views.Error{
// 				Error: "Something went wrong please try again",
// 			},
// 		}
// 		c.Status(400)
// 		return c.JSON(response)
// 	}

// 	response := models.Response{
// 		Type: models.TypeSuccessResponse,
// 		Data: views.Success{
// 			Message: "Cart item deleted successfuly",
// 		},
// 	}

// 	return c.JSON(response)
// }

// func GetCart(c *fiber.Ctx) error {
// 	var items []views.ItemResponse
// 	var total float64 = 0

// 	rows, err := db.DB.Query(`SELECT item_id FROM carts WHERE user_id = $1;`, c.GetRespHeader("request_user_id"))
// 	if err != nil {
// 		response := models.Response{
// 			Type: models.TypeErrorResponse,
// 			Data: views.Error{
// 				Error: "Something went wrong please try again",
// 			},
// 		}
// 		c.Status(400)
// 		return c.JSON(response)
// 	}
// 	defer rows.Close()

// 	var item views.ItemResponse
// 	var itemID string
// 	for rows.Next() {
// 		if err := rows.Scan(&itemID); err != nil {
// 			response := models.Response{
// 				Type: models.TypeErrorResponse,
// 				Data: views.Error{
// 					Error: "Something went wrong please try again",
// 				},
// 			}
// 			c.Status(400)
// 			return c.JSON(response)
// 		}
// 		query := db.DB.QueryRow(`SELECT id, name, sku, price FROM items WHERE id = $1;`, itemID)
// 		err := query.Scan(&item.ID, &item.Name, &item.SKU, &item.Price)
// 		if err != nil {
// 			response := models.Response{
// 				Type: models.TypeErrorResponse,
// 				Data: views.Error{
// 					Error: "Invalid item ID",
// 				},
// 			}
// 			c.Status(400)
// 			return c.JSON(response)
// 		}
// 		total += item.Price
// 		item.Quantity = 1
// 		items = append(items, item)
// 	}

// 	var itemsResponse []views.ItemResponse
// 	var found bool
// 	for _, item := range items {
// 		found = false

// 		for i, item2 := range itemsResponse {
// 			if item.ID == item2.ID {
// 				found = true
// 				item2.Quantity++
// 				itemsResponse[i] = item2
// 				break
// 			}
// 		}

// 		if found {
// 			continue
// 		}
// 		itemsResponse = append(itemsResponse, item)
// 	}

// 	response := models.Response{
// 		Type: models.TypeUserCart,
// 		Data: views.CartResponse{
// 			Items: itemsResponse,
// 			Total: total,
// 		},
// 	}

// 	return c.JSON(response)
// }

// func EmptyCart(c *fiber.Ctx) error {
// 	userID := c.GetRespHeader("request_user_id")

// 	_, err := db.DB.Exec(`DELETE FROM carts WHERE user_id = $1;`, userID)
// 	if err != nil {
// 		response := models.Response{
// 			Type: models.TypeErrorResponse,
// 			Data: views.Error{
// 				Error: "Something went wrong please try again",
// 			},
// 		}
// 		c.Status(400)
// 		return c.JSON(response)
// 	}

// 	response := models.Response{
// 		Type: models.TypeSuccessResponse,
// 		Data: views.Success{
// 			Message: "Cart emptied successfuly",
// 		},
// 	}

// 	return c.JSON(response)
// }
