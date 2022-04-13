package controllers

import (
	"database/sql"
	"stores/db"
	"stores/models"
	"stores/views"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func CreateCoupon(c *fiber.Ctx) error {
	coupon := models.Coupon{}
	err := c.BodyParser(&coupon)
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

	if _, isValid := coupon.Validate(); !isValid {
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
		coupon.ID = uuid.New().String()
		query := db.DB.QueryRow(`SELECT id FROM coupons WHERE id = $1;`, coupon.ID)
		err = query.Scan(&coupon.ID)
		if err != nil {
			break
		}
	}

	if coupon.Type != models.CouponTypeFixed && coupon.Type != models.CouponTypePercentage {
		response := models.Response{
			Type: models.TypeErrorResponse,
			Data: views.Error{
				Error: "Invalid coupon type",
			},
		}
		c.Status(400)
		return c.JSON(response)
	}
	coupon.Used = 0
	coupon.CreatedAt = time.Now().UTC()
	coupon.UpdatedAt = time.Now().UTC()

	_, err = db.DB.Exec(`INSERT INTO coupons(id, value, type, max_usage, used, code, end_date, created_at, updated_at)
	VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9);`, coupon.ID, coupon.Value, coupon.Type, coupon.MaxUsage, coupon.Used, coupon.Code, coupon.EndDate, coupon.CreatedAt, coupon.UpdatedAt)
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
			Message: "Coupon created successfuly",
		},
	}

	return c.JSON(response)
}

func UpdateCoupon(c *fiber.Ctx) error {
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
				Error: "Category name already exists",
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
