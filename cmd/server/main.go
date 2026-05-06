package main

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"github.com/joho/godotenv"

	"github.com/antoniogisondi/mgl-safety-cloud-backend/internal/database"

	"github.com/antoniogisondi/mgl-safety-cloud-backend/internal/auth"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Errore caricamento .env")
	}

	database.Connect()
	database.Migrate()

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "API Online",
		})
	})

	api := app.Group("/api")

	api.Post("/auth/register", auth.Register)
	api.Post("/auth/login", auth.Login)

	log.Fatal(app.Listen(":8080"))
}
