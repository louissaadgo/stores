package routes

import (
	"stores/controllers"
	"stores/middlewares"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func Initialize(app *fiber.App) {
	app.Use(cors.New())

	//Unauthorized routes
	app.Post("/api/v1/auth/web/signup", controllers.WebSignup)
	app.Post("/api/v1/auth/web/login", controllers.WebLogin)
	app.Get("/api/v1/auth/web/current/user/type", controllers.WebCurrentUserType)

	app.Post("/api/v1/auth/user/signup", controllers.UserSignup)
	app.Post("/api/v1/auth/user/signin", controllers.UserSignin)
	app.Get("/api/v1/auth/user/reset/password/request", controllers.UserResetPasswordRequest)
	app.Post("/api/v1/auth/user/reset/password", controllers.UserResetPassword)

	app.Get("/api/v1/categories", controllers.GetAllCategories)
	app.Get("/api/v1/currencies", controllers.GetAllCurrencies)

	app.Get("/api/v1/attributes", controllers.GetAllAttributes)

	//User specific routes
	app.Use("/api/v1/user/", middlewares.UserMiddleware)
	app.Get("/api/v1/user/request/otp", controllers.UserRequestOTP)
	app.Post("/api/v1/user/verify/otp", controllers.UserVerifyOTP)

	// app.Post("/api/v1/user/addresses", controllers.CreateAddress)
	// app.Put("/api/v1/user/addresses/:id", controllers.UpdateAddress)
	// app.Delete("/api/v1/user/addresses/:id", controllers.DeleteAddress)

	// app.Post("/api/v1/user/favorites/:id", controllers.AddFavorite)
	// app.Delete("/api/v1/user/favorites/:id", controllers.DeleteFavorite)

	// app.Get("/api/v1/user/carts", controllers.GetCart)
	// app.Delete("/api/v1/user/carts", controllers.EmptyCart)
	// app.Post("/api/v1/user/carts/:id", controllers.AddToCart)
	// app.Delete("/api/v1/user/carts/:id", controllers.DeleteFromCart)

	// app.Post("/api/v1/user/orders", controllers.CreateOrder)

	// app.Get("/api/v1/user/wallets", controllers.GetAllWallets)

	//Merchant specific routes
	app.Use("/api/v1/merchant/", middlewares.MerchantMiddleware)

	//Add Get stores
	app.Post("/api/v1/merchant/stores", controllers.CreateStore)
	app.Put("/api/v1/merchant/stores/:id", controllers.UpdateStore)

	//get all links
	app.Post("/api/v1/merchant/links", controllers.CreateLink)
	app.Put("/api/v1/merchant/links/:id", controllers.UpdateLink)

	//get all items
	app.Post("/api/v1/merchant/items", controllers.CreateItem)
	app.Put("/api/v1/merchant/items/:id", controllers.UpdateItem)

	//Admin specific routes
	app.Use("/api/v1/admin/", middlewares.AdminMiddleware)

	app.Get("/api/v1/admin/current", controllers.GetCurrentAdmin)

	app.Post("/api/v1/admin/admins/", controllers.CreateAdmin)
	app.Get("/api/v1/admin/admins/", controllers.GetAllAdmins)

	app.Get("/api/v1/admin/merchants/", controllers.GetAllMerchants)
	app.Put("/api/v1/admin/ban/merchant/:id", controllers.BanMerchant)
	app.Put("/api/v1/admin/active/merchant/:id", controllers.ActivateMerchant)

	//Add all users route
	app.Put("/api/v1/admin/ban/user/:id", controllers.BanUser)
	app.Put("/api/v1/admin/activate/user/:id", controllers.ActivateUser)

	app.Post("/api/v1/admin/attributes", controllers.CreateAttribute)
	app.Put("/api/v1/admin/attributes/:id", controllers.UpdateAttribute)

	app.Post("/api/v1/admin/attributevalues", controllers.CreateAttributeValue)
	app.Put("/api/v1/admin/attributevalues/:id", controllers.UpdateAttributeValue)

	app.Post("/api/v1/admin/currencies", controllers.CreateCurrency)
	app.Put("/api/v1/admin/currencies/:id", controllers.UpdateCurrency)

	app.Post("/api/v1/admin/categories", controllers.CreateCategory)
	app.Put("/api/v1/admin/categories/:id", controllers.UpdateCategory)

	app.Post("/api/v1/admin/subcategories", controllers.CreateSubCategory)
	app.Put("/api/v1/admin/subcategories/:id", controllers.UpdateSubCategory)

	app.Post("/api/v1/admin/transactions", controllers.CreateTransaction)
}
