package main

import (
	"log"
	"service-api/infra/config"
)

func main() {
	app := config.SetupApp()

	log.Println("Server starting on :3000")
	if err := app.Listen(":3000"); err != nil {
		log.Fatal(err)
	}
}
