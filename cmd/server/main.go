package main

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"github.com/joho/godotenv"

	"github.com/antoniogisondi/mgl-safety-cloud-backend/internal/database"

	"github.com/antoniogisondi/mgl-safety-cloud-backend/internal/auth"

	"github.com/antoniogisondi/mgl-safety-cloud-backend/internal/middleware"

	"github.com/antoniogisondi/mgl-safety-cloud-backend/internal/clients"

	"github.com/antoniogisondi/mgl-safety-cloud-backend/internal/companies"

	"github.com/antoniogisondi/mgl-safety-cloud-backend/internal/workers"
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

	protected := api.Group("/", middleware.Protected())

	protected.Get("/me", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"user_id": c.Locals("user_id"),
			"email":   c.Locals("email"),
			"role":    c.Locals("role"),
		})
	})

	protected.Get("/clients", clients.GetClients)
	protected.Get("/clients/:id", clients.GetClient)
	protected.Post("/clients", clients.CreateClient)
	protected.Put("/clients/:id", clients.UpdateClient)
	protected.Delete("/clients/:id", clients.DeleteClient)

	protected.Get("/companies", companies.GetCompanies)
	protected.Get("/companies/:id", companies.GetCompany)
	protected.Post("/companies", companies.CreateCompany)
	protected.Put("/companies/:id", companies.UpdateCompany)
	protected.Delete("/companies/:id", companies.DeleteCompany)

	protected.Get("/workers", workers.GetWorkers)
	protected.Get("/workers/:id", workers.GetWorker)
	protected.Post("/workers", workers.CreateWorker)
	protected.Put("/workers/:id", workers.UpdateWorker)
	protected.Delete("/workers/:id", workers.DeleteWorker)

	log.Fatal(app.Listen(":8080"))
}
