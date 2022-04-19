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

func UserSignup(c *fiber.Ctx) error {
	user := models.User{}
	err := c.BodyParser(&user)
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

	if _, isValid := user.Validate(); !isValid {
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
		user.ID = uuid.New().String()
		query := db.DB.QueryRow(`SELECT id FROM users WHERE id = $1;`, user.ID)
		err = query.Scan(&user.ID)
		if err != nil {
			break
		}
	}

	query := db.DB.QueryRow(`SELECT phone FROM users WHERE phone = $1;`, user.Phone)
	err = query.Scan(&user.Phone)
	if err == nil || err != sql.ErrNoRows {
		response := models.Response{
			Type: models.TypeErrorResponse,
			Data: views.Error{
				Error: "phone already exists",
			},
		}
		c.Status(400)
		return c.JSON(response)
	}

	if user.SignType == models.SignTypeNative {
		hashedPassword, isValid := HashPassword(user.Password)
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
		user.Password = hashedPassword
		user.SignID = ""
	} else if user.SignType == models.SignTypeParty {
		user.Password = ""
	} else {
		response := models.Response{
			Type: models.TypeErrorResponse,
			Data: views.Error{
				Error: "invalid sign_type",
			},
		}
		c.Status(400)
		return c.JSON(response)
	}
	user.Status = models.UserStatusActive
	user.LoyalityPoints = 0
	user.CreatedAt = time.Now().UTC()
	user.UpdatedAt = time.Now().UTC()

	tokenID := uuid.New().String()

	_, err = db.DB.Exec(`INSERT INTO users(id, name, phone, password, sign_type, sign_id, bday, image, country, status, loyality_points, created_at, updated_at, token_id)
	VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14);`, user.ID, user.Name, user.Phone, user.Password, user.SignType, user.SignID, user.Bday, user.Image, user.Country, user.Status, user.LoyalityPoints, user.CreatedAt, user.UpdatedAt, tokenID)
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

	token, err := token.GeneratePasetoToken(tokenID, user.ID, models.TypeUser)
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

func UserSignin(c *fiber.Ctx) error {
	user := models.User{}
	err := c.BodyParser(&user)

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

	if _, isValid := user.Validate(); !isValid {
		response := models.Response{
			Type: models.TypeErrorResponse,
			Data: views.Error{
				Error: "Invalid Data",
			},
		}
		c.Status(400)
		return c.JSON(response)
	}

	password := user.Password
	signID := user.SignID
	query := db.DB.QueryRow(`SELECT id, password, sign_id FROM users WHERE phone = $1;`, user.Phone)
	err = query.Scan(&user.ID, &user.Password, &user.SignID)
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

	if user.SignType == models.SignTypeNative {
		if isValid := ValidatePassword(password, user.Password); !isValid {
			response := models.Response{
				Type: models.TypeErrorResponse,
				Data: views.Error{
					Error: "Invalid Credentials",
				},
			}
			c.Status(400)
			return c.JSON(response)
		}
	} else if user.SignType == models.SignTypeParty {
		if signID != user.SignID {
			response := models.Response{
				Type: models.TypeErrorResponse,
				Data: views.Error{
					Error: "Invalid Credentials",
				},
			}
			c.Status(400)
			return c.JSON(response)
		}
	} else {
		response := models.Response{
			Type: models.TypeErrorResponse,
			Data: views.Error{
				Error: "Invalid sign_type",
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
	_, err = db.DB.Exec(`UPDATE users SET token_id = $1 WHERE id = $2;`, tokenID, user.ID)
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

	token, err := token.GeneratePasetoToken(tokenID, user.ID, models.TypeUser)
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
