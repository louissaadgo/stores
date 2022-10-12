package controllers

import (
	"database/sql"
	"fmt"
	"stores/core"
	"stores/db"

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
			AuthToken:  token,
			ExpiryDate: time.Now().Add(time.Hour * 2),
		},
	}

	subject := fmt.Sprintf("Welcome %v", merchant.Name)
	message := fmt.Sprintf("Welcome to Aswak, %v.\nYou have been successfully registered as a merchant.", merchant.Name)
	go core.SendEmail(merchant.Email, subject, message)

	return c.JSON(response)
}

func WebCurrentUserType(c *fiber.Ctx) error {

	tokenString := models.Token{}
	err := c.BodyParser(&tokenString.Token)
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

	payload, isValid := token.VerifyPasetoToken(tokenString.Token)
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

	admin := models.AdminLogin{}
	err := c.BodyParser(&admin)
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

	emailQuery := admin.Email
	passwordQuery := admin.Password

	var password string
	var userID int

	query := db.DB.QueryRow(`SELECT id, password FROM merchants WHERE email = $1;`, emailQuery)
	err = query.Scan(&userID, &password)
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
				AuthToken:  token,
				ExpiryDate: time.Now().Add(time.Hour * 2),
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

	var verifiedPhone bool = false

	query := db.DB.QueryRow(`SELECT phone, verified_phone FROM users WHERE phone = $1;`, user.Phone)
	err = query.Scan(&user.Phone, &verifiedPhone)

	if err == nil || err != sql.ErrNoRows {
		response := models.Response{
			Type: "phone_already_registered",
			Data: views.Error{
				Error: "phone already registered",
			},
		}
		c.Status(400)
		return c.JSON(response)
	}

	hashedPassword, _ := HashPassword(user.Password)
	user.Password = hashedPassword
	user.Status = models.UserStatusActive
	user.VerifiedPhone = false
	user.OTP = ""
	user.CreatedAt = time.Now().UTC()
	user.UpdatedAt = time.Now().UTC()

	tokenID := uuid.New().String()

	_, err = db.DB.Exec(`INSERT INTO users(name, phone, image, verified_phone, otp, password, token_id, country, status, created_at, updated_at)
	VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11);`, user.Name, user.Phone, "", user.VerifiedPhone, user.OTP, user.Password, tokenID, user.Country, user.Status, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		response := models.Response{
			Type: models.TypeErrorResponse,
			Data: views.Error{
				Error: "Error occured while inserting into db",
			},
		}
		c.Status(400)
		return c.JSON(response)
	}

	oneRow := db.DB.QueryRow(`SELECT id FROM users WHERE phone = $1;`, user.Phone)
	err = oneRow.Scan(&user.ID)
	if err != nil {
		response := models.Response{
			Type: models.TypeErrorResponse,
			Data: views.Error{
				Error: "Error occured while inserting into db",
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
		Data: views.UserAuth{
			AuthToken: token,
			UserData: views.UserAuthData{
				Name:          user.Name,
				Phone:         user.Phone,
				Country:       user.Country,
				Status:        user.Status,
				VerifiedPhone: user.VerifiedPhone,
			},
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

	password := user.Password

	query := db.DB.QueryRow(`SELECT id, password, name, phone, country, status, verified_phone FROM users WHERE phone = $1;`, user.Phone)
	err = query.Scan(&user.ID, &user.Password, &user.Name, &user.Phone, &user.Country, &user.Status, &user.VerifiedPhone)
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
				Error: "Error modifying token_id",
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
				Error: "error while generating paseto token",
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
		Data: views.UserAuth{
			AuthToken: token,
			UserData: views.UserAuthData{
				Name:          user.Name,
				Phone:         user.Phone,
				Country:       user.Country,
				Status:        user.Status,
				VerifiedPhone: user.VerifiedPhone,
			},
		},
	}

	return c.JSON(response)
}

func UserRequestOTP(c *fiber.Ctx) error {

	userID := c.GetRespHeader("request_user_id")
	OTP := core.GenerateRandomNumber()

	_, err := db.DB.Exec(`UPDATE users SET otp = $1 WHERE id = $2;`, OTP, userID)
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

	// if !otp.SendOTP(OTP) {
	// 	response := models.Response{
	// 		Type: models.TypeErrorResponse,
	// 		Data: views.Error{
	// 			Error: "Something went wrong please try again",
	// 		},
	// 	}
	// 	c.Status(400)
	// 	return c.JSON(response)
	// }

	return c.SendString("Success")
}

func UserVerifyOTP(c *fiber.Ctx) error {

	userID := c.GetRespHeader("request_user_id")

	testOTP := c.Params("otp")

	otpToken := models.OTP{}
	otpToken.OTPToken = testOTP

	var OTP string
	query := db.DB.QueryRow(`SELECT otp FROM users WHERE id = $1;`, userID)
	err := query.Scan(&OTP)
	if err != nil {
		response := models.Response{
			Type: "error",
			Data: views.Error{
				Error: "something went wrong",
			},
		}
		c.Status(400)
		return c.JSON(response)
	}

	if otpToken.OTPToken != OTP {
		response := models.Response{
			Type: "invalid_otp",
			Data: views.Error{
				Error: "invalid otp",
			},
		}
		c.Status(400)
		return c.JSON(response)
	}

	_, err = db.DB.Exec(`UPDATE users SET verified_phone = $1 WHERE id = $2;`, true, userID)
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

	return c.SendString("Success")
}

func UserResetPasswordRequest(c *fiber.Ctx) error {

	phoneNumber := models.PhoneNumber{}
	err := c.BodyParser(&phoneNumber)
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

	OTP := core.GenerateRandomNumber()

	query := db.DB.QueryRow(`SELECT phone FROM users WHERE phone = $1;`, phoneNumber.Phone)
	err = query.Scan(&phoneNumber.Phone)
	if err == sql.ErrNoRows {
		response := models.Response{
			Type: "Error",
			Data: views.Error{
				Error: "Phone not found",
			},
		}
		c.Status(400)
		return c.JSON(response)
	}

	_, err = db.DB.Exec(`UPDATE users SET otp = $1 WHERE phone = $2;`, OTP, phoneNumber.Phone)
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

	// if !otp.SendOTP(OTP) {
	// 	response := models.Response{
	// 		Type: models.TypeErrorResponse,
	// 		Data: views.Error{
	// 			Error: "Something went wrong please try again",
	// 		},
	// 	}
	// 	c.Status(400)
	// 	return c.JSON(response)
	// }

	return c.SendString("Success")
}

func UserResetPassword(c *fiber.Ctx) error {

	testOTP := c.Params("otp")

	otpAndPhone := models.OTPAndPhoneAndPassword{}
	otpAndPhone.OTPToken = testOTP

	err := c.BodyParser(&otpAndPhone)
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

	var OTP string
	query := db.DB.QueryRow(`SELECT otp FROM users WHERE phone = $1;`, otpAndPhone.Phone)
	err = query.Scan(&OTP)
	if err != nil {
		response := models.Response{
			Type: "error",
			Data: views.Error{
				Error: "something went wrong",
			},
		}
		c.Status(400)
		return c.JSON(response)
	}

	if otpAndPhone.OTPToken != OTP {
		response := models.Response{
			Type: "invalid_otp",
			Data: views.Error{
				Error: "invalid otp",
			},
		}
		c.Status(400)
		return c.JSON(response)
	}

	OTP = core.GenerateRandomNumber()
	newPass, _ := HashPassword(otpAndPhone.Password)

	_, err = db.DB.Exec(`UPDATE users SET otp = $1, password = $2 WHERE phone = $3;`, OTP, newPass, otpAndPhone.Phone)
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

	return c.SendString("success")

}

func UserAddPictureAndName(c *fiber.Ctx) error {
	userID := c.GetRespHeader("request_user_id")

	user := models.UserImageAndName{}
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

	_, err = db.DB.Exec(`UPDATE users SET name = $1, image = $2 WHERE id = $3;`, user.Name, user.Image, userID)
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

	return c.SendString("Success")
}

func UserVerifyOTPForPassword(c *fiber.Ctx) error {

	testOTP := c.Params("otp")

	otpToken := models.OTPWithPhone{}
	err := c.BodyParser(&otpToken)
	if err != nil {
		response := models.Response{
			Type: "error",
			Data: views.Error{
				Error: "something went wrong",
			},
		}
		c.Status(400)
		return c.JSON(response)
	}
	otpToken.OTPToken = testOTP

	var OTP string
	query := db.DB.QueryRow(`SELECT otp FROM users WHERE phone = $1;`, otpToken.Phone)
	err = query.Scan(&OTP)
	if err != nil {
		response := models.Response{
			Type: "error",
			Data: views.Error{
				Error: "something went wrong",
			},
		}
		c.Status(400)
		return c.JSON(response)
	}

	if otpToken.OTPToken != OTP {
		response := models.Response{
			Type: "invalid_otp",
			Data: views.Error{
				Error: "invalid otp",
			},
		}
		c.Status(400)
		return c.JSON(response)
	}

	return c.SendString("Success")
}
