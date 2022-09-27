package controllers

import (
	"database/sql"
	"stores/db"
	"stores/models"
	"stores/views"
	"time"

	"github.com/gofiber/fiber/v2"
)

func GetAllCategories(c *fiber.Ctx) error {
	var categories []views.CategoryResponse
	rows, err := db.DB.Query(`SELECT id, name, created_at, updated_at FROM categories;`)
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

	var category views.CategoryResponse
	for rows.Next() {
		if err := rows.Scan(&category.ID, &category.Name, &category.CreatedAt, &category.UpdatedAt); err != nil {
			response := models.Response{
				Type: models.TypeErrorResponse,
				Data: views.Error{
					Error: "Something went wrong please try again",
				},
			}
			c.Status(400)
			return c.JSON(response)
		}
		newRows, err := db.DB.Query(`SELECT id, name, category_id, created_at, updated_at FROM subcategories WHERE category_id = $1;`, category.ID)
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
		var subCategories []views.SubCategoryResponse
		for newRows.Next() {
			var subCategory views.SubCategoryResponse
			if err := newRows.Scan(&subCategory.ID, &subCategory.Name, &subCategory.CategoryID, &subCategory.CreatedAt, &subCategory.UpdatedAt); err != nil {
				response := models.Response{
					Type: models.TypeErrorResponse,
					Data: views.Error{
						Error: "Something went wrong please try again",
					},
				}
				c.Status(400)
				return c.JSON(response)
			}
			subCategories = append(subCategories, subCategory)
		}
		category.SubCategories = subCategories
		categories = append(categories, category)
	}

	response := models.Response{
		Type: models.TypeAllCategories,
		Data: views.AllCategories{
			Categories: categories,
		},
	}

	return c.JSON(response)
}

func CreateCategory(c *fiber.Ctx) error {
	category := models.Category{}
	err := c.BodyParser(&category)
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

	if _, isValid := category.Validate(); !isValid {
		response := models.Response{
			Type: models.TypeErrorResponse,
			Data: views.Error{
				Error: "Invalid Data",
			},
		}
		c.Status(400)
		return c.JSON(response)
	}

	query := db.DB.QueryRow(`SELECT name FROM categories WHERE name = $1;`, category.Name)
	err = query.Scan(&category.Name)
	if err == nil || err != sql.ErrNoRows {
		response := models.Response{
			Type: "category_already_exists",
			Data: views.Error{
				Error: "Category name already exists",
			},
		}
		c.Status(400)
		return c.JSON(response)
	}
	category.CreatedAt = time.Now().UTC()
	category.UpdatedAt = time.Now().UTC()

	_, err = db.DB.Exec(`INSERT INTO categories(name, created_at, updated_at)
	VALUES($1, $2, $3);`, category.Name, category.CreatedAt, category.UpdatedAt)
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
			Message: "Category added successfuly",
		},
	}

	return c.JSON(response)
}

func UpdateCategory(c *fiber.Ctx) error {
	id := c.Params("id")

	category := models.Category{}
	err := c.BodyParser(&category)
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

	if _, isValid := category.Validate(); !isValid {
		response := models.Response{
			Type: models.TypeErrorResponse,
			Data: views.Error{
				Error: "Invalid Data",
			},
		}
		c.Status(400)
		return c.JSON(response)
	}

	query := db.DB.QueryRow(`SELECT id FROM categories WHERE id = $1;`, id)
	err = query.Scan(&id)
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

	query = db.DB.QueryRow(`SELECT name FROM categories WHERE name = $1;`, category.Name)
	err = query.Scan(&category.Name)
	if err == nil || err != sql.ErrNoRows {
		response := models.Response{
			Type: models.TypeErrorResponse,
			Data: views.Error{
				Error: "category_name_already_exists",
			},
		}
		c.Status(400)
		return c.JSON(response)
	}

	_, err = db.DB.Exec(`UPDATE categories SET name = $1, updated_at = $2 WHERE id = $3;`, category.Name, time.Now().UTC(), id)
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
			Message: "Category updated successfuly",
		},
	}

	return c.JSON(response)
}

func CreateSubCategory(c *fiber.Ctx) error {
	subCategory := models.SubCategory{}
	err := c.BodyParser(&subCategory)
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

	if _, isValid := subCategory.Validate(); !isValid {
		response := models.Response{
			Type: models.TypeErrorResponse,
			Data: views.Error{
				Error: "Invalid Data",
			},
		}
		c.Status(400)
		return c.JSON(response)
	}

	query := db.DB.QueryRow(`SELECT id FROM categories WHERE id = $1;`, subCategory.CategoryID)
	var attributeID string
	err = query.Scan(&attributeID)
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

	subCategory.CreatedAt = time.Now().UTC()
	subCategory.UpdatedAt = time.Now().UTC()

	_, err = db.DB.Exec(`INSERT INTO subcategories(name, category_id, created_at, updated_at)
	VALUES($1, $2, $3, $4);`, subCategory.Name, subCategory.CategoryID, subCategory.CreatedAt, subCategory.UpdatedAt)
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
			Message: "SubCategory added successfuly",
		},
	}

	return c.JSON(response)
}

func UpdateSubCategory(c *fiber.Ctx) error {
	id := c.Params("id")

	subCategory := models.SubCategory{}
	err := c.BodyParser(&subCategory)
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

	if _, isValid := subCategory.Validate(); !isValid {
		response := models.Response{
			Type: models.TypeErrorResponse,
			Data: views.Error{
				Error: "Invalid Data",
			},
		}
		c.Status(400)
		return c.JSON(response)
	}

	query := db.DB.QueryRow(`SELECT id FROM subcategories WHERE id = $1;`, id)
	err = query.Scan(&id)
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

	_, err = db.DB.Exec(`UPDATE subcategories SET name = $1, updated_at = $2 WHERE id = $3;`, subCategory.Name, time.Now().UTC(), id)
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
			Message: "Subcategory updated successfuly",
		},
	}

	return c.JSON(response)
}
