package main

import (
	"log"
	"service-api/infra/http/routers"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	routers.SetupRoutes(app)

	log.Println("Server starting on :3000")
	if err := app.Listen(":3000"); err != nil {
		log.Fatal(err)
	}
}
