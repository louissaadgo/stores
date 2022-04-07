package routes

import (
	"stores/controllers"

	"github.com/gofiber/fiber/v2"
)

func Initialize(app *fiber.App) {
	app.Post("/api/v1/auth/admin/signin", controllers.AdminSignin)
}
