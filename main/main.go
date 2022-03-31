package main

import (
	"stores/routes"

	"github.com/gofiber/fiber/v2"
)

func init() {

}

func main() {
	app := fiber.New()
	routes.Initialize(app)

	app.Listen(":4000")
}
