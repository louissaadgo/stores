package controllers

import (
	"stores/db"
	"stores/models"
	"stores/views"

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
