package clients

import (
	"github.com/gofiber/fiber/v2"

	"github.com/antoniogisondi/mgl-safety-cloud-backend/internal/database"
	"github.com/antoniogisondi/mgl-safety-cloud-backend/internal/models"
)

func GetClients(c *fiber.Ctx) error {
	var clients []models.Client

	if err := database.DB.Find(&clients).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Errore recupero clienti",
		})
	}

	return c.JSON(clients)
}

func GetClient(c *fiber.Ctx) error {
	id := c.Params("id")

	var client models.Client

	if err := database.DB.First(&client, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "Cliente non trovato",
		})
	}

	return c.JSON(client)
}

func CreateClient(c *fiber.Ctx) error {
	var client models.Client

	if err := c.BodyParser(&client); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Dati non validi",
		})
	}

	if err := database.DB.Create(&client).Error; err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Errore creazione cliente",
		})
	}

	return c.Status(201).JSON(client)
}

func UpdateClient(c *fiber.Ctx) error {
	id := c.Params("id")

	var client models.Client

	if err := database.DB.First(&client, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "Cliente non trovato",
		})
	}

	if err := c.BodyParser(&client); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Dati non validi",
		})
	}

	if err := database.DB.Save(&client).Error; err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Errore aggiornamento cliente",
		})
	}

	return c.JSON(client)
}

func DeleteClient(c *fiber.Ctx) error {
	id := c.Params("id")

	var client models.Client

	if err := database.DB.First(&client, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "Cliente non trovato",
		})
	}

	if err := database.DB.Delete(&client).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Errore eliminazione cliente",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Cliente eliminato correttamente",
	})
}
