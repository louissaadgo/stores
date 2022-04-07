package routes

import (
	"stores/controllers"
	"stores/middlewares"

	"github.com/gofiber/fiber/v2"
)

func Initialize(app *fiber.App) {
	//Unauthorized routes
	app.Post("/api/v1/auth/admin/signin", controllers.AdminSignin)
	app.Post("/api/v1/auth/merchant/signup", controllers.MerchantSignup)
	app.Post("/api/v1/auth/merchant/signin", controllers.MerchantSignin)
	app.Post("/api/v1/auth/user/signup", controllers.UserSignup)
	app.Post("/api/v1/auth/user/signin", controllers.UserSignin)

	//Admin specific routes
	app.Use(middlewares.AdminMiddleware)
}
