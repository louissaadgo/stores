package controllers

import (
	"stores/db"
	"stores/models"
	"stores/views"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

func CreateItem(c *fiber.Ctx) error {
	item := models.Item{}
	err := c.BodyParser(&item)
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

	if _, isValid := item.Validate(); !isValid {
		response := models.Response{
			Type: models.TypeErrorResponse,
			Data: views.Error{
				Error: "Invalid Data",
			},
		}
		c.Status(400)
		return c.JSON(response)
	}

	query := db.DB.QueryRow(`SELECT id, merchant_id FROM stores WHERE id = $1;`, item.StoreID)
	var merchantID int
	err = query.Scan(&item.StoreID, &merchantID)
	if err != nil {
		response := models.Response{
			Type: models.TypeErrorResponse,
			Data: views.Error{
				Error: "Invalid store ID",
			},
		}
		c.Status(400)
		return c.JSON(response)
	}

	realMerchantID, _ := strconv.Atoi(c.GetRespHeader("request_user_id"))
	if merchantID != realMerchantID {
		response := models.Response{
			Type: models.TypeErrorResponse,
			Data: views.Error{
				Error: "Merchant can only add items to his own stores",
			},
		}
		c.Status(400)
		return c.JSON(response)
	}

	query = db.DB.QueryRow(`SELECT id FROM categories WHERE id = $1;`, item.CategoryID)
	err = query.Scan(&item.CategoryID)
	if err != nil {
		response := models.Response{
			Type: models.TypeErrorResponse,
			Data: views.Error{
				Error: "Invalid category ID",
			},
		}
		c.Status(400)
		return c.JSON(response)
	}

	query = db.DB.QueryRow(`SELECT id FROM subcategories WHERE id = $1;`, item.SubCategoryID)
	err = query.Scan(&item.SubCategoryID)
	if err != nil {
		response := models.Response{
			Type: models.TypeErrorResponse,
			Data: views.Error{
				Error: "Invalid subcategory ID",
			},
		}
		c.Status(400)
		return c.JSON(response)
	}

	item.CreatedAt = time.Now().UTC()
	item.UpdatedAt = time.Now().UTC()
	item.Status = models.ItemStatusActive

	_, err = db.DB.Exec(`INSERT INTO items(id, name, sku, description, long_description, price, store_id, created_at, updated_at, category_id, subcategory_id, stock, status)
	VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13);`, item.ID, item.Name, item.SKU, item.Description, item.LongDescription, item.Price, item.StoreID, item.CreatedAt, item.UpdatedAt, item.CategoryID, item.SubCategoryID, item.Stock, item.Status)
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

	for _, image := range item.Images {
		_, err = db.DB.Exec(`INSERT INTO item_images(item_id, url)
	VALUES($1, $2);`, item.ID, image)
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

	response := models.Response{
		Type: models.TypeSuccessResponse,
		Data: views.Success{
			Message: "Item added successfuly",
		},
	}

	return c.JSON(response)
}

func UpdateItem(c *fiber.Ctx) error {
	id := c.Params("id")

	item := models.Item{}
	err := c.BodyParser(&item)
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

	if _, isValid := item.Validate(); !isValid {
		response := models.Response{
			Type: models.TypeErrorResponse,
			Data: views.Error{
				Error: "Invalid Data",
			},
		}
		c.Status(400)
		return c.JSON(response)
	}

	query := db.DB.QueryRow(`SELECT id, store_id FROM items WHERE id = $1;`, id)
	var storeID string
	err = query.Scan(&id, &storeID)
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

	query = db.DB.QueryRow(`SELECT merchant_id FROM stores WHERE id = $1;`, storeID)
	var merchantID int
	err = query.Scan(&merchantID)
	if err != nil {
		response := models.Response{
			Type: models.TypeErrorResponse,
			Data: views.Error{
				Error: "Invalid store ID",
			},
		}
		c.Status(400)
		return c.JSON(response)
	}

	realMerchantID, _ := strconv.Atoi(c.GetRespHeader("request_user_id"))
	if merchantID != realMerchantID {
		response := models.Response{
			Type: models.TypeErrorResponse,
			Data: views.Error{
				Error: "Merchant can only edit his own items",
			},
		}
		c.Status(400)
		return c.JSON(response)
	}

	query = db.DB.QueryRow(`SELECT id FROM categories WHERE id = $1;`, item.CategoryID)
	err = query.Scan(&item.CategoryID)
	if err != nil {
		response := models.Response{
			Type: models.TypeErrorResponse,
			Data: views.Error{
				Error: "Invalid category ID",
			},
		}
		c.Status(400)
		return c.JSON(response)
	}

	query = db.DB.QueryRow(`SELECT id FROM subcategories WHERE id = $1;`, item.SubCategoryID)
	err = query.Scan(&item.SubCategoryID)
	if err != nil {
		response := models.Response{
			Type: models.TypeErrorResponse,
			Data: views.Error{
				Error: "Invalid subcategory ID",
			},
		}
		c.Status(400)
		return c.JSON(response)
	}

	_, err = db.DB.Exec(`UPDATE items SET name = $1, sku = $2, description = $3, long_description = $4, price = $5, updated_at = $6, category_id = $7, subcategory_id = $8, stock = $9 WHERE id = $10;`, item.Name, item.SKU, item.Description, item.LongDescription, item.Price, time.Now().UTC(), item.CategoryID, item.SubCategoryID, item.Stock, id)
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
			Message: "Item updated successfuly",
		},
	}

	return c.JSON(response)
}
