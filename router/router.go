package router

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func Initialize() {
	app := fiber.New()
	initializeRoutes(app)
	log.Fatal(app.Listen(":8080"))
}
