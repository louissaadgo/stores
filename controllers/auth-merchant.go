package controllers

import (
	"database/sql"
	"stores/db"
	"stores/models"
	"stores/token"
	"stores/views"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func MerchantSignup(c *fiber.Ctx) error {
	merchant := models.Merchant{}
	err := c.BodyParser(&merchant)
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

	if _, isValid := merchant.Validate(); !isValid {
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
		merchant.ID = uuid.New().String()
		query := db.DB.QueryRow(`SELECT id FROM merchants WHERE id = $1;`, merchant.ID)
		err = query.Scan(&merchant.ID)
		if err != nil {
			break
		}
	}

	query := db.DB.QueryRow(`SELECT email FROM merchants WHERE email = $1;`, merchant.Email)
	err = query.Scan(&merchant.Email)
	if err == nil || err != sql.ErrNoRows {
		response := models.Response{
			Type: models.TypeErrorResponse,
			Data: views.Error{
				Error: "Email already exists",
			},
		}
		c.Status(400)
		return c.JSON(response)
	}

	hashedPassword, isValid := HashPassword(merchant.Password)
	if !isValid {
		response := models.Response{
			Type: models.TypeErrorResponse,
			Data: views.Error{
				Error: "Something went wrong please try again",
			},
		}
		c.Status(400)
		return c.JSON(response)
	}
	merchant.Password = hashedPassword
	merchant.Status = models.MerchantStatusActive
	merchant.CreatedAt = time.Now().UTC()
	merchant.UpdatedAt = time.Now().UTC()

	tokenID := uuid.New().String()

	_, err = db.DB.Exec(`INSERT INTO merchants(id, email, password, name, status, created_at, updated_at, token_id)
	VALUES($1, $2, $3, $4, $5, $6, $7, $8);`, merchant.ID, merchant.Email, merchant.Password, merchant.Name, merchant.Status, merchant.CreatedAt, merchant.UpdatedAt, tokenID)
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

	if tokenID == "" {
		response := models.Response{
			Type: models.TypeErrorResponse,
			Data: views.Error{
				Error: "Something went wrong please try again",
			},
		}
		c.Status(400)
		return c.JSON(response)
	}

	token, err := token.GeneratePasetoToken(tokenID, merchant.ID, models.TypeMerchant)
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

	cookie := fiber.Cookie{
		Name:  "token",
		Value: token,
	}
	c.Cookie(&cookie)

	response := models.Response{
		Type: models.TypeAuthResponse,
		Data: views.Auth{
			AuthToken: token,
		},
	}

	return c.JSON(response)
}

func MerchantSignin(c *fiber.Ctx) error {
	merchant := models.Merchant{}
	err := c.BodyParser(&merchant)

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

	if _, isValid := merchant.Validate(); !isValid {
		response := models.Response{
			Type: models.TypeErrorResponse,
			Data: views.Error{
				Error: "Invalid Data",
			},
		}
		c.Status(400)
		return c.JSON(response)
	}

	password := merchant.Password
	query := db.DB.QueryRow(`SELECT id, password FROM merchants WHERE email = $1;`, merchant.Email)
	err = query.Scan(&merchant.ID, &merchant.Password)
	if err != nil {
		response := models.Response{
			Type: models.TypeErrorResponse,
			Data: views.Error{
				Error: "Invalid Credentials",
			},
		}
		c.Status(400)
		return c.JSON(response)
	}

	if isValid := ValidatePassword(password, merchant.Password); !isValid {
		response := models.Response{
			Type: models.TypeErrorResponse,
			Data: views.Error{
				Error: "Invalid Credentials",
			},
		}
		c.Status(400)
		return c.JSON(response)
	}

	tokenID := uuid.New().String()
	if tokenID == "" {
		response := models.Response{
			Type: models.TypeErrorResponse,
			Data: views.Error{
				Error: "Something went wrong please try again",
			},
		}
		c.Status(400)
		return c.JSON(response)
	}

	_, err = db.DB.Exec(`UPDATE merchants SET token_id = $1 WHERE id = $2;`, tokenID, merchant.ID)
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

	token, err := token.GeneratePasetoToken(tokenID, merchant.ID, models.TypeMerchant)
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

	cookie := fiber.Cookie{
		Name:  "token",
		Value: token,
	}
	c.Cookie(&cookie)

	response := models.Response{
		Type: models.TypeAuthResponse,
		Data: views.Auth{
			AuthToken: token,
		},
	}

	return c.JSON(response)
}
