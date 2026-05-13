package companies

import (
	"github.com/gofiber/fiber/v2"

	"github.com/antoniogisondi/mgl-safety-cloud-backend/internal/database"
	"github.com/antoniogisondi/mgl-safety-cloud-backend/internal/models"
)

func GetCompanies(c *fiber.Ctx) error {
	var companies []models.Company

	if err := database.DB.
		Preload("Workers.Deadlines").
		Preload("Deadlines").
		Find(&companies).Error; err != nil {

		return c.Status(500).JSON(fiber.Map{
			"error": "Errore recupero aziende",
		})
	}

	for i := range companies {

		if companies[i].Workers == nil {
			companies[i].Workers = []models.Worker{}
		}

		if companies[i].Deadlines == nil {
			companies[i].Deadlines = []models.Deadline{}
		}

		for j := range companies[i].Workers {

			if companies[i].Workers[j].Deadlines == nil {
				companies[i].Workers[j].Deadlines = []models.Deadline{}
			}
		}
	}

	return c.JSON(companies)
}

func GetCompany(c *fiber.Ctx) error {
	id := c.Params("id")

	var company models.Company

	if err := database.DB.
		Preload("Workers.Deadlines").
		Preload("Deadlines").
		First(&company, id).Error; err != nil {

		return c.Status(404).JSON(fiber.Map{
			"error": "Azienda non trovata",
		})
	}

	if company.Workers == nil {
		company.Workers = []models.Worker{}
	}

	if company.Deadlines == nil {
		company.Deadlines = []models.Deadline{}
	}

	for i := range company.Workers {

		if company.Workers[i].Deadlines == nil {
			company.Workers[i].Deadlines = []models.Deadline{}
		}
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
