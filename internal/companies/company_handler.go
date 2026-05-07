package companies

import (
	"github.com/gofiber/fiber/v2"

	"github.com/antoniogisondi/mgl-safety-cloud-backend/internal/database"
	"github.com/antoniogisondi/mgl-safety-cloud-backend/internal/models"
)

func GetCompanies(c *fiber.Ctx) error {
	var companies []models.Company

	if err := database.DB.Preload("Client").Find(&companies).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Errore recupero aziende",
		})
	}

	return c.JSON(companies)
}

func GetCompany(c *fiber.Ctx) error {
	id := c.Params("id")

	var company models.Company

	if err := database.DB.Preload("Client").First(&company, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "Azienda non trovata",
		})
	}

	return c.JSON(company)
}

func CreateCompany(c *fiber.Ctx) error {
	var company models.Company

	if err := c.BodyParser(&company); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Dati non validi",
		})
	}

	var client models.Client

	if err := database.DB.First(&client, company.ClientID).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "Cliente collegato non trovato",
		})
	}

	if err := database.DB.Create(&company).Error; err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Errore creazione azienda",
		})
	}

	database.DB.Preload("Client").First(&company, company.ID)

	return c.Status(201).JSON(company)
}

func UpdateCompany(c *fiber.Ctx) error {
	id := c.Params("id")

	var company models.Company

	if err := database.DB.First(&company, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "Azienda non trovata",
		})
	}

	if err := c.BodyParser(&company); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Dati non validi",
		})
	}

	var client models.Client

	if err := database.DB.First(&client, company.ClientID).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "Cliente collegato non trovato",
		})
	}

	if err := database.DB.Save(&company).Error; err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Errore aggiornamento azienda",
		})
	}

	database.DB.Preload("Client").First(&company, company.ID)

	return c.JSON(company)
}

func DeleteCompany(c *fiber.Ctx) error {
	id := c.Params("id")

	var company models.Company

	if err := database.DB.First(&company, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "Azienda non trovata",
		})
	}

	if err := database.DB.Delete(&company).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Errore eliminazione azienda",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Azienda eliminata correttamente",
	})
}
