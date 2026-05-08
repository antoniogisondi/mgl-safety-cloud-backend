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

	if err := database.DB.Preload("Category").Find(&deadlines).Error; err != nil {
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

	if err := database.DB.Preload("Category").First(&deadline, id).Error; err != nil {
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

	// Controllo azienda
	var company models.Company

	if err := database.DB.First(&company, deadline.CompanyID).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "Azienda collegata non trovata",
		})
	}

	// Controllo categoria
	var category models.DeadlineCategory

	if err := database.DB.First(&category, deadline.CategoryID).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "Categoria scadenza non trovata",
		})
	}

	// Controllo lavoratore
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

	// Calcolo stato
	deadline.Status = calculateStatus(deadline.ExpirationDate)

	// Creazione scadenza
	if err := database.DB.Create(&deadline).Error; err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Errore creazione scadenza",
		})
	}

	// Reload con categoria
	database.DB.Preload("Category").First(&deadline, deadline.ID)

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

func GetExpiredDeadlines(c *fiber.Ctx) error {
	var deadlines []models.Deadline
	now := time.Now()

	if err := database.DB.
		Where("expiration_date < ?", now).
		Find(&deadlines).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Errore recupero scadenze scadute",
		})
	}

	for i := range deadlines {
		deadlines[i].Status = calculateStatus(deadlines[i].ExpirationDate)
	}

	return c.JSON(deadlines)
}

func GetExpiringDeadlines(c *fiber.Ctx) error {
	var deadlines []models.Deadline
	now := time.Now()
	limit := now.AddDate(0, 0, 30)

	if err := database.DB.
		Where("expiration_date >= ? AND expiration_date <= ?", now, limit).
		Find(&deadlines).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Errore recupero scadenze in scadenza",
		})
	}

	for i := range deadlines {
		deadlines[i].Status = calculateStatus(deadlines[i].ExpirationDate)
	}

	return c.JSON(deadlines)
}

func GetDeadlinesByCompany(c *fiber.Ctx) error {
	companyID := c.Params("companyId")

	var deadlines []models.Deadline

	if err := database.DB.
		Where("company_id = ?", companyID).
		Find(&deadlines).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Errore recupero scadenze azienda",
		})
	}

	for i := range deadlines {
		deadlines[i].Status = calculateStatus(deadlines[i].ExpirationDate)
	}

	return c.JSON(deadlines)
}

func GetDeadlinesByWorker(c *fiber.Ctx) error {
	workerID := c.Params("workerId")

	var deadlines []models.Deadline

	if err := database.DB.
		Where("worker_id = ?", workerID).
		Find(&deadlines).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Errore recupero scadenze lavoratore",
		})
	}

	for i := range deadlines {
		deadlines[i].Status = calculateStatus(deadlines[i].ExpirationDate)
	}

	return c.JSON(deadlines)
}
