package main

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"github.com/joho/godotenv"

	"github.com/antoniogisondi/mgl-safety-cloud-backend/internal/database"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Errore caricamento .env")
	}

	database.Connect()

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "API Online",
		})
	})

	log.Fatal(app.Listen(":8080"))
}
