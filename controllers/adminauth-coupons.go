package controllers

import (
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

	query := db.DB.QueryRow(`SELECT id FROM coupons WHERE id = $1;`, id)
	err = query.Scan(&id)
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

	_, err = db.DB.Exec(`UPDATE coupons SET value = $1, type = $2, max_usage = $3, code = $4, end_date = $5, updated_at = $6  WHERE id = $7;`, coupon.Value, coupon.Type, coupon.MaxUsage, coupon.Code, coupon.EndDate, time.Now().UTC(), id)
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
			Message: "Coupon updated successfuly",
		},
	}

	return c.JSON(response)
}

func GetAllCoupons(c *fiber.Ctx) error {
	var coupons []models.Coupon
	rows, err := db.DB.Query(`SELECT id, value, type, max_usage, used, code, end_date, created_at, updated_at FROM coupons;`)
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

	var coupon models.Coupon
	for rows.Next() {
		if err := rows.Scan(&coupon.ID, &coupon.Value, &coupon.Type, &coupon.MaxUsage, &coupon.Used, &coupon.Code, &coupon.EndDate, &coupon.CreatedAt, &coupon.UpdatedAt); err != nil {
			response := models.Response{
				Type: models.TypeErrorResponse,
				Data: views.Error{
					Error: "Something went wrong please try again",
				},
			}
			c.Status(400)
			return c.JSON(response)
		}
		coupons = append(coupons, coupon)
	}

	response := models.Response{
		Type: models.TypeAllCoupons,
		Data: views.AllCoupons{
			Coupons: coupons,
		},
	}

	return c.JSON(response)
}
