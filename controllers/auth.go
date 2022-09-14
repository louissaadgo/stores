package controllers

import (
	"database/sql"
	"fmt"
	"stores/db"
	"stores/emailing"
	"stores/models"
	"stores/token"
	"stores/views"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func WebSignup(c *fiber.Ctx) error {

	merchant := models.Merchant{}
	err := c.BodyParser(&merchant)
	if err != nil {
		response := models.Response{
			Type: "invalid_data_types",
			Data: views.Error{
				Error: "Invalid Data Types",
			},
		}
		c.Status(400)
		return c.JSON(response)
	}

	if _, isValid := merchant.Validate(); !isValid {
		response := models.Response{
			Type: "invalid_data",
			Data: views.Error{
				Error: "Invalid Data",
			},
		}
		c.Status(400)
		return c.JSON(response)
	}

	query := db.DB.QueryRow(`SELECT email FROM merchants WHERE email = $1;`, merchant.Email)
	err = query.Scan(&merchant.Email)
	if err == nil || err != sql.ErrNoRows {
		response := models.Response{
			Type: "email_already_registered",
			Data: views.Error{
				Error: "email already registered",
			},
		}
		c.Status(400)
		return c.JSON(response)
	}

	query = db.DB.QueryRow(`SELECT email FROM admins WHERE email = $1;`, merchant.Email)
	err = query.Scan(&merchant.Email)
	if err == nil || err != sql.ErrNoRows {
		response := models.Response{
			Type: "email_already_registered",
			Data: views.Error{
				Error: "email already registered",
			},
		}
		c.Status(400)
		return c.JSON(response)
	}

	hashedPassword, isValid := HashPassword(merchant.Password)
	if !isValid {
		response := models.Response{
			Type: "error_hashing_password",
			Data: views.Error{
				Error: "Something went wrong while hashsing the password",
			},
		}
		c.Status(400)
		return c.JSON(response)
	}

	merchant.Password = hashedPassword
	merchant.Status = models.MerchantStatusInactive
	creationTime := time.Now().UTC()
	merchant.CreatedAt = creationTime
	merchant.UpdatedAt = creationTime

	tokenID := uuid.New().String()

	_, err = db.DB.Exec(`INSERT INTO merchants(email, password, name, status, created_at, updated_at, token_id)
		VALUES($1, $2, $3, $4, $5, $6, $7);`, merchant.Email, merchant.Password, merchant.Name, merchant.Status, merchant.CreatedAt, merchant.UpdatedAt, tokenID)
	if err != nil {
		response := models.Response{
			Type: "error_inserting_into_db",
			Data: views.Error{
				Error: "Something went wrong while inserting into the db",
			},
		}
		c.Status(400)
		return c.JSON(response)
	}

	query = db.DB.QueryRow(`SELECT id FROM merchants WHERE email = $1;`, merchant.Email)
	err = query.Scan(&merchant.ID)
	if err != nil {
		response := models.Response{
			Type: "errore_while_reading_id",
			Data: views.Error{
				Error: "Error while reading merchant id",
			},
		}
		c.Status(400)
		return c.JSON(response)
	}

	token, err := token.GeneratePasetoToken(tokenID, merchant.ID, models.TypeMerchant)
	if err != nil {
		response := models.Response{
			Type: "errorgenerating_paseto",
			Data: views.Error{
				Error: "Error while generating the paseto token",
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
		Data: views.AuthWeb{
			AuthToken: token,
		},
	}

	subject := fmt.Sprintf("Welcome %v", merchant.Name)
	message := fmt.Sprintf("Welcome to Aswak, %v.\nYou have been successfully registered as a merchant.", merchant.Name)
	go emailing.SendEmail(merchant.Email, subject, message)

	return c.JSON(response)
}

func WebCurrentUserType(c *fiber.Ctx) error {

	tokenString := c.Query("token", "")

	payload, isValid := token.VerifyPasetoToken(tokenString)
	if !isValid {
		response := models.Response{
			Type: "error_unauthenticated",
			Data: views.Error{
				Error: "Invalid token",
			},
		}
		c.Status(400)
		return c.JSON(response)
	}

	if payload.UserType != models.TypeMerchant && payload.UserType != models.TypeAdmin {
		response := models.Response{
			Type: "error_invalid_user_type",
			Data: views.Error{
				Error: "Error invalid user type",
			},
		}
		c.Status(400)
		return c.JSON(response)
	}

	var userStatus string
	var name string
	if payload.UserType == models.TypeMerchant {
		query := db.DB.QueryRow(`SELECT name, status FROM merchants WHERE id = $1;`, payload.UserID)
		err := query.Scan(&name, &userStatus)
		if err == sql.ErrNoRows {
			response := models.Response{
				Type: "error_unauthenticated",
				Data: views.Error{
					Error: "Invalid token",
				},
			}
			c.Status(400)
			return c.JSON(response)
		}
		if userStatus == models.MerchantStatusBanned {
			response := models.Response{
				Type: "error_merchant_banned",
				Data: views.Error{
					Error: "Merchant banned",
				},
			}
			c.Status(400)
			return c.JSON(response)
		}
		if userStatus == models.MerchantStatusInactive {
			response := models.Response{
				Type: "error_merchant_inactive",
				Data: views.Error{
					Error: "Merchant inactive",
				},
			}
			c.Status(400)
			return c.JSON(response)
		}
	}

	if payload.UserType == models.TypeAdmin {
		query := db.DB.QueryRow(`SELECT name FROM admins WHERE id = $1;`, payload.UserID)
		err := query.Scan(&name)
		if err == sql.ErrNoRows {
			response := models.Response{
				Type: "error_unauthenticated",
				Data: views.Error{
					Error: "Invalid token",
				},
			}
			c.Status(400)
			return c.JSON(response)
		}
	}

	response := models.Response{
		Type: "success",
		Data: views.CurrentTypeWeb{
			CurrentType: payload.UserType,
			Name:        name,
		},
	}

	return c.JSON(response)
}

func WebLogin(c *fiber.Ctx) error {

	emailQuery := c.Query("email", "")
	passwordQuery := c.Query("password", "")

	var password string
	var userID int

	query := db.DB.QueryRow(`SELECT id, password FROM merchants WHERE email = $1;`, emailQuery)
	err := query.Scan(&userID, &password)
	if err == nil || err != sql.ErrNoRows {
		if isValid := ValidatePassword(passwordQuery, password); !isValid {
			response := models.Response{
				Type: "invalid_credentials",
				Data: views.Error{
					Error: err.Error(),
				},
			}
			c.Status(400)
			return c.JSON(response)
		}
		tokenID := uuid.New().String()
		_, err = db.DB.Exec(`UPDATE merchants SET token_id = $1 WHERE email = $2;`, tokenID, emailQuery)
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
		token, err := token.GeneratePasetoToken(tokenID, userID, models.TypeMerchant)
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
			Data: views.AuthWeb{
				AuthToken: token,
			},
		}

		return c.JSON(response)
	}

	query = db.DB.QueryRow(`SELECT id, password FROM admins WHERE email = $1;`, emailQuery)
	err = query.Scan(&userID, &password)
	if err == nil || err != sql.ErrNoRows {
		if isValid := ValidatePassword(passwordQuery, password); !isValid {
			response := models.Response{
				Type: "invalid_credentials",
				Data: views.Error{
					Error: "Invalid Credentials",
				},
			}
			c.Status(400)
			return c.JSON(response)
		}
		tokenID := uuid.New().String()
		_, err = db.DB.Exec(`UPDATE admins SET token_id = $1 WHERE email = $2;`, tokenID, emailQuery)
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
		token, err := token.GeneratePasetoToken(tokenID, userID, models.TypeAdmin)
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
			Data: views.AuthWeb{
				AuthToken: token,
			},
		}

		return c.JSON(response)
	}

	response := models.Response{
		Type: "invalid_credentials",
		Data: views.Error{
			Error: "Invalid Credentials",
		},
	}
	c.Status(400)
	return c.JSON(response)
}
