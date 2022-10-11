package main

import (
	"log"
	"os"
	"stores/db"
	"stores/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	os.Mkdir("storage", os.ModePerm)

	err, isValid := db.InitializeDB()
	if !isValid {
		log.Fatalln(err)
	}
	defer db.DB.Close()

	app := fiber.New()
	routes.Initialize(app)
	app.Listen(":4000")
}
