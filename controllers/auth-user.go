package controllers

// func UserSignup(c *fiber.Ctx) error {
// 	user := models.User{}
// 	err := c.BodyParser(&user)
// 	if err != nil {
// 		response := models.Response{
// 			Type: models.TypeErrorResponse,
// 			Data: views.Error{
// 				Error: "Invalid Data Types",
// 			},
// 		}
// 		c.Status(400)
// 		return c.JSON(response)
// 	}

// 	if _, isValid := user.Validate(); !isValid {
// 		response := models.Response{
// 			Type: models.TypeErrorResponse,
// 			Data: views.Error{
// 				Error: "Invalid Data",
// 			},
// 		}
// 		c.Status(400)
// 		return c.JSON(response)
// 	}

// 	for {
// 		user.ID = uuid.New().String()
// 		query := db.DB.QueryRow(`SELECT id FROM users WHERE id = $1;`, user.ID)
// 		err = query.Scan(&user.ID)
// 		if err != nil {
// 			break
// 		}
// 	}

// 	query := db.DB.QueryRow(`SELECT phone FROM users WHERE phone = $1;`, user.Phone)
// 	err = query.Scan(&user.Phone)
// 	if err == nil || err != sql.ErrNoRows {
// 		response := models.Response{
// 			Type: models.TypeErrorResponse,
// 			Data: views.Error{
// 				Error: "phone already exists",
// 			},
// 		}
// 		c.Status(400)
// 		return c.JSON(response)
// 	}

// 	query = db.DB.QueryRow(`SELECT email FROM users WHERE email = $1;`, user.Email)
// 	err = query.Scan(&user.Email)
// 	if err == nil || err != sql.ErrNoRows {
// 		response := models.Response{
// 			Type: models.TypeErrorResponse,
// 			Data: views.Error{
// 				Error: "email already exists",
// 			},
// 		}
// 		c.Status(400)
// 		return c.JSON(response)
// 	}

// 	hashedPassword, _ := HashPassword(user.Password)
// 	user.Password = hashedPassword
// 	user.Status = models.UserStatusActive
// 	user.VerifiedEmail = false
// 	user.VerifiedPhone = false
// 	user.CreatedAt = time.Now().UTC()
// 	user.UpdatedAt = time.Now().UTC()

// 	tokenID := uuid.New().String()

// 	_, err = db.DB.Exec(`INSERT INTO users(id, name, phone, verified_phone, email, verified_email, password, token_id, country, status, created_at, updated_at)
// 	VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12);`, user.ID, user.Name, user.Phone, user.VerifiedPhone, user.Email, user.VerifiedEmail, user.Password, user.TokenID, user.Country, user.Status, user.CreatedAt, user.UpdatedAt)
// 	if err != nil {
// 		response := models.Response{
// 			Type: models.TypeErrorResponse,
// 			Data: views.Error{
// 				Error: "Error occured while inserting into db",
// 			},
// 		}
// 		c.Status(400)
// 		return c.JSON(response)
// 	}

// 	token, err := token.GeneratePasetoToken(tokenID, user.ID, models.TypeUser)
// 	if err != nil {
// 		response := models.Response{
// 			Type: models.TypeErrorResponse,
// 			Data: views.Error{
// 				Error: "Something went wrong please try again",
// 			},
// 		}
// 		c.Status(400)
// 		return c.JSON(response)
// 	}

// 	cookie := fiber.Cookie{
// 		Name:  "token",
// 		Value: token,
// 	}
// 	c.Cookie(&cookie)

// 	response := models.Response{
// 		Type: models.TypeAuthResponse,
// 		Data: views.UserAuth{
// 			AuthToken: token,
// 			UserData: views.UserAuthData{
// 				Name:          user.Name,
// 				Phone:         user.Phone,
// 				Country:       user.Country,
// 				Status:        user.Status,
// 				Email:         user.Email,
// 				VerifiedEmail: user.VerifiedEmail,
// 				VerifiedPhone: user.VerifiedPhone,
// 			},
// 		},
// 	}

// 	emailing.SendEmail(user.Email, "Welcome to Aswak", "Hi "+user.Name+", welcome to aswak!!")

// 	return c.JSON(response)
// }

// func UserSignin(c *fiber.Ctx) error {
// 	user := models.User{}
// 	err := c.BodyParser(&user)

// 	if err != nil {
// 		response := models.Response{
// 			Type: models.TypeErrorResponse,
// 			Data: views.Error{
// 				Error: "Invalid Data Types",
// 			},
// 		}
// 		c.Status(400)
// 		return c.JSON(response)
// 	}

// 	password := user.Password

// 	query := db.DB.QueryRow(`SELECT id, password, email, name, phone, country, status, verified_email, verified_phone FROM users WHERE phone = $1;`, user.Phone)
// 	err = query.Scan(&user.ID, &user.Password, &user.Email, &user.Name, &user.Phone, &user.Country, &user.Status, &user.VerifiedEmail, &user.VerifiedPhone)
// 	if err != nil {
// 		response := models.Response{
// 			Type: models.TypeErrorResponse,
// 			Data: views.Error{
// 				Error: "Invalid Credentials",
// 			},
// 		}
// 		c.Status(400)
// 		return c.JSON(response)
// 	}

// 	if isValid := ValidatePassword(password, user.Password); !isValid {
// 		response := models.Response{
// 			Type: models.TypeErrorResponse,
// 			Data: views.Error{
// 				Error: "Invalid Credentials",
// 			},
// 		}
// 		c.Status(400)
// 		return c.JSON(response)
// 	}

// 	tokenID := uuid.New().String()
// 	if tokenID == "" {
// 		response := models.Response{
// 			Type: models.TypeErrorResponse,
// 			Data: views.Error{
// 				Error: "Something went wrong please try again",
// 			},
// 		}
// 		c.Status(400)
// 		return c.JSON(response)
// 	}

// 	_, err = db.DB.Exec(`UPDATE users SET token_id = $1 WHERE id = $2;`, tokenID, user.ID)
// 	if err != nil {
// 		response := models.Response{
// 			Type: models.TypeErrorResponse,
// 			Data: views.Error{
// 				Error: "Error modifying token_id",
// 			},
// 		}
// 		c.Status(400)
// 		return c.JSON(response)
// 	}

// 	token, err := token.GeneratePasetoToken(tokenID, user.ID, models.TypeUser)
// 	if err != nil {
// 		response := models.Response{
// 			Type: models.TypeErrorResponse,
// 			Data: views.Error{
// 				Error: "error while generating paseto token",
// 			},
// 		}
// 		c.Status(400)
// 		return c.JSON(response)
// 	}

// 	cookie := fiber.Cookie{
// 		Name:  "token",
// 		Value: token,
// 	}
// 	c.Cookie(&cookie)

// 	response := models.Response{
// 		Type: models.TypeAuthResponse,
// 		Data: views.UserAuth{
// 			AuthToken: token,
// 			UserData: views.UserAuthData{
// 				Name:          user.Name,
// 				Phone:         user.Phone,
// 				Country:       user.Country,
// 				Status:        user.Status,
// 				Email:         user.Email,
// 				VerifiedEmail: user.VerifiedEmail,
// 				VerifiedPhone: user.VerifiedPhone,
// 			},
// 		},
// 	}

// 	return c.JSON(response)
// }
