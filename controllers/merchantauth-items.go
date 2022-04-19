package controllers

import (
	"stores/db"
	"stores/models"
	"stores/views"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
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

	for {
		item.ID = uuid.New().String()
		query := db.DB.QueryRow(`SELECT id FROM items WHERE id = $1;`, item.ID)
		err = query.Scan(&item.ID)
		if err != nil {
			break
		}
	}

	query := db.DB.QueryRow(`SELECT id, merchant_id FROM stores WHERE id = $1;`, item.StoreID)
	var merchantID string
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

	if merchantID != c.GetRespHeader("request_user_id") {
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

	_, err = db.DB.Exec(`INSERT INTO items(id, name, sku, description, long_description, price, images, store_id, created_at, updated_at, category_id, subcategory_id, stock, status)
	VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14);`, item.ID, item.Name, item.SKU, item.Description, item.LongDescription, item.Price, item.Images, item.StoreID, item.CreatedAt, item.UpdatedAt, item.CategoryID, item.SubCategoryID, item.Stock, item.Status)
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
			Message: "Item added successfuly",
		},
	}

	return c.JSON(response)
}

func UpdateItem(c *fiber.Ctx) error {
	id := c.Params("id")

	store := models.Store{}
	err := c.BodyParser(&store)
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

	if _, isValid := store.Validate(); !isValid {
		response := models.Response{
			Type: models.TypeErrorResponse,
			Data: views.Error{
				Error: "Invalid Data",
			},
		}
		c.Status(400)
		return c.JSON(response)
	}

	query := db.DB.QueryRow(`SELECT id, merchant_id FROM stores WHERE id = $1;`, id)
	var merchantID string
	err = query.Scan(&id, &merchantID)
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

	if merchantID != c.GetRespHeader("request_user_id") {
		response := models.Response{
			Type: models.TypeErrorResponse,
			Data: views.Error{
				Error: "Merchant can only edits his own stores",
			},
		}
		c.Status(400)
		return c.JSON(response)
	}

	_, err = db.DB.Exec(`UPDATE stores SET description = $1, phone = $2, location = $3, country = $4, updated_at = $5 WHERE id = $6;`, store.Description, store.Phone, store.Location, store.Country, time.Now().UTC(), id)
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
			Message: "Store updated successfuly",
		},
	}

	return c.JSON(response)
}
