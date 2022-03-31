package routes

import (
	"stores/controllers"

	"github.com/gofiber/fiber/v2"
)

func Initialize(app *fiber.App) {

	app.Post("/api/auth/store/signup", controllers.StoreSignup)
}
