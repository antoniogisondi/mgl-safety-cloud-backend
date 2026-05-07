package workers

import (
	"github.com/gofiber/fiber/v2"

	"github.com/antoniogisondi/mgl-safety-cloud-backend/internal/database"
	"github.com/antoniogisondi/mgl-safety-cloud-backend/internal/models"
)

func GetWorkers(c *fiber.Ctx) error {
	var workers []models.Worker

	if err := database.DB.Preload("Deadlines").Find(&workers).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Errore recupero lavoratori",
		})
	}

	for i := range workers {
		if workers[i].Deadlines == nil {
			workers[i].Deadlines = []models.Deadline{}
		}
	}

	return c.JSON(workers)
}

func GetWorker(c *fiber.Ctx) error {
	id := c.Params("id")

	var worker models.Worker

	if err := database.DB.Preload("Deadlines").First(&worker, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "Lavoratore non trovato",
		})
	}

	if worker.Deadlines == nil {
		worker.Deadlines = []models.Deadline{}
	}

	return c.JSON(worker)
}

func CreateWorker(c *fiber.Ctx) error {
	var worker models.Worker

	if err := c.BodyParser(&worker); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Dati non validi",
		})
	}

	var company models.Company

	if err := database.DB.First(&company, worker.CompanyID).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "Azienda collegata non trovata",
		})
	}

	if err := database.DB.Create(&worker).Error; err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Errore creazione lavoratore",
		})
	}

	return c.Status(201).JSON(worker)
}

func UpdateWorker(c *fiber.Ctx) error {
	id := c.Params("id")

	var worker models.Worker

	if err := database.DB.First(&worker, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "Lavoratore non trovato",
		})
	}

	if err := c.BodyParser(&worker); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Dati non validi",
		})
	}

	if err := database.DB.Save(&worker).Error; err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Errore aggiornamento lavoratore",
		})
	}

	return c.JSON(worker)
}

func DeleteWorker(c *fiber.Ctx) error {
	id := c.Params("id")

	var worker models.Worker

	if err := database.DB.First(&worker, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "Lavoratore non trovato",
		})
	}

	if err := database.DB.Delete(&worker).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Errore eliminazione lavoratore",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Lavoratore eliminato correttamente",
	})
}
