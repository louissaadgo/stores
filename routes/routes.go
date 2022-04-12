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

	//User specific routes
	app.Use("/api/v1/user/", middlewares.UserMiddleware)

	//Merchant specific routes
	app.Use("/api/v1/merchant/", middlewares.MerchantMiddleware)

	//Admin specific routes
	app.Use("/api/v1/admin/", middlewares.AdminMiddleware)
	app.Post("/api/v1/admin/admins/", controllers.CreateAdmin)
	app.Get("/api/v1/admin/admins/", controllers.GetAllAdmins)

	app.Post("/api/v1/admin/ban/merchant/:id", controllers.BanMerchant)
	app.Post("/api/v1/admin/activate/merchant/:id", controllers.ActivateMerchant)
}
