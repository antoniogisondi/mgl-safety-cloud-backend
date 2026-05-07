package deadlinecategories

import (
	"github.com/gofiber/fiber/v2"

	"github.com/antoniogisondi/mgl-safety-cloud-backend/internal/database"
	"github.com/antoniogisondi/mgl-safety-cloud-backend/internal/models"
)

func GetCategories(c *fiber.Ctx) error {
	var categories []models.DeadlineCategory

	if err := database.DB.Find(&categories).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Errore recupero categorie",
		})
	}

	return c.JSON(categories)
}

func CreateCategory(c *fiber.Ctx) error {
	var category models.DeadlineCategory

	if err := c.BodyParser(&category); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Dati non validi",
		})
	}

	if err := database.DB.Create(&category).Error; err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Errore creazione categoria",
		})
	}

	return c.Status(201).JSON(category)
}
