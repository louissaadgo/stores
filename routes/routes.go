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

	app.Get("/api/v1/attributes", controllers.GetAllAttributes)
	app.Get("/api/v1/categories", controllers.GetAllCategories)

	//User specific routes
	app.Use("/api/v1/user/", middlewares.UserMiddleware)

	app.Post("/api/v1/user/addresses", controllers.CreateAddress)
	app.Put("/api/v1/user/addresses/:id", controllers.UpdateAddress)
	app.Delete("/api/v1/user/addresses/:id", controllers.DeleteAddress)

	app.Post("/api/v1/user/favorites/:id", controllers.AddFavorite)
	app.Delete("/api/v1/user/favorites/:id", controllers.DeleteFavorite)

	app.Post("/api/v1/user/interests/:id", controllers.AddInterest)
	app.Delete("/api/v1/user/interests/:id", controllers.DeleteInterest)

	app.Get("/api/v1/user/carts", controllers.GetCart)
	app.Post("/api/v1/user/carts/:id", controllers.AddToCart)
	app.Delete("/api/v1/user/carts/:id", controllers.DeleteFromCart)

	app.Get("/api/v1/user/wallets", controllers.GetAllWallets)

	//Merchant specific routes
	app.Use("/api/v1/merchant/", middlewares.MerchantMiddleware)

	app.Post("/api/v1/merchant/stores", controllers.CreateStore)
	app.Put("/api/v1/merchant/stores/:id", controllers.UpdateStore)

	app.Post("/api/v1/merchant/links", controllers.CreateLink)
	app.Put("/api/v1/merchant/links/:id", controllers.UpdateLink)

	app.Post("/api/v1/merchant/items", controllers.CreateItem)
	app.Put("/api/v1/merchant/items/:id", controllers.UpdateItem)

	//Admin specific routes
	app.Use("/api/v1/admin/", middlewares.AdminMiddleware)

	app.Post("/api/v1/admin/admins/", controllers.CreateAdmin)
	app.Get("/api/v1/admin/admins/", controllers.GetAllAdmins)

	app.Put("/api/v1/admin/ban/merchant/:id", controllers.BanMerchant)
	app.Put("/api/v1/admin/ban/user/:id", controllers.BanUser)

	app.Put("/api/v1/admin/activate/merchant/:id", controllers.ActivateMerchant)
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

	app.Get("/api/v1/admin/coupons", controllers.GetAllCoupons)
	app.Post("/api/v1/admin/coupons", controllers.CreateCoupon)
	app.Put("/api/v1/admin/coupons/:id", controllers.UpdateCoupon)

	app.Post("/api/v1/admin/transactions", controllers.CreateTransaction)
}
