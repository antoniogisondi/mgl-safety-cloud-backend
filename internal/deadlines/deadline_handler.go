package deadlines

import (
	"time"

	"github.com/gofiber/fiber/v2"

	"github.com/antoniogisondi/mgl-safety-cloud-backend/internal/database"
	"github.com/antoniogisondi/mgl-safety-cloud-backend/internal/models"
)

func calculateStatus(expirationDate *time.Time) string {
	if expirationDate == nil {
		return "no_expiration"
	}

	now := time.Now()
	daysLeft := int(expirationDate.Sub(now).Hours() / 24)

	if daysLeft < 0 {
		return "expired"
	}

	if daysLeft <= 30 {
		return "expiring"
	}

	return "valid"
}

func GetDeadlines(c *fiber.Ctx) error {
	var deadlines []models.Deadline

	if err := database.DB.Find(&deadlines).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Errore recupero scadenze",
		})
	}

	for i := range deadlines {
		deadlines[i].Status = calculateStatus(deadlines[i].ExpirationDate)
	}

	return c.JSON(deadlines)
}

func GetDeadline(c *fiber.Ctx) error {
	id := c.Params("id")

	var deadline models.Deadline

	if err := database.DB.First(&deadline, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "Scadenza non trovata",
		})
	}

	deadline.Status = calculateStatus(deadline.ExpirationDate)

	return c.JSON(deadline)
}

func CreateDeadline(c *fiber.Ctx) error {
	var deadline models.Deadline

	if err := c.BodyParser(&deadline); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Dati non validi",
		})
	}

	var company models.Company
	if err := database.DB.First(&company, deadline.CompanyID).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "Azienda collegata non trovata",
		})
	}

	if deadline.WorkerID != nil {
		var worker models.Worker
		if err := database.DB.First(&worker, *deadline.WorkerID).Error; err != nil {
			return c.Status(404).JSON(fiber.Map{
				"error": "Lavoratore collegato non trovato",
			})
		}

		if worker.CompanyID != deadline.CompanyID {
			return c.Status(400).JSON(fiber.Map{
				"error": "Il lavoratore non appartiene all'azienda indicata",
			})
		}
	}

	deadline.Status = calculateStatus(deadline.ExpirationDate)

	if err := database.DB.Create(&deadline).Error; err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Errore creazione scadenza",
		})
	}

	return c.Status(201).JSON(deadline)
}

func UpdateDeadline(c *fiber.Ctx) error {
	id := c.Params("id")

	var deadline models.Deadline

	if err := database.DB.First(&deadline, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "Scadenza non trovata",
		})
	}

	if err := c.BodyParser(&deadline); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Dati non validi",
		})
	}

	var company models.Company
	if err := database.DB.First(&company, deadline.CompanyID).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "Azienda collegata non trovata",
		})
	}

	if deadline.WorkerID != nil {
		var worker models.Worker
		if err := database.DB.First(&worker, *deadline.WorkerID).Error; err != nil {
			return c.Status(404).JSON(fiber.Map{
				"error": "Lavoratore collegato non trovato",
			})
		}

		if worker.CompanyID != deadline.CompanyID {
			return c.Status(400).JSON(fiber.Map{
				"error": "Il lavoratore non appartiene all'azienda indicata",
			})
		}
	}

	deadline.Status = calculateStatus(deadline.ExpirationDate)

	if err := database.DB.Save(&deadline).Error; err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Errore aggiornamento scadenza",
		})
	}

	return c.JSON(deadline)
}

func DeleteDeadline(c *fiber.Ctx) error {
	id := c.Params("id")

	var deadline models.Deadline

	if err := database.DB.First(&deadline, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "Scadenza non trovata",
		})
	}

	if err := database.DB.Delete(&deadline).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Errore eliminazione scadenza",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Scadenza eliminata correttamente",
	})
}
