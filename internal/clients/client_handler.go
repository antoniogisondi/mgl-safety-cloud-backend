package clients

import (
	"github.com/gofiber/fiber/v2"

	"github.com/antoniogisondi/mgl-safety-cloud-backend/internal/database"
	"github.com/antoniogisondi/mgl-safety-cloud-backend/internal/models"
)

func GetClients(c *fiber.Ctx) error {
	var clients []models.Client

	if err := database.DB.Preload("Companies.Workers").Find(&clients).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Errore recupero clienti",
		})
	}

	for i := range clients {
		if clients[i].Companies == nil {
			clients[i].Companies = []models.Company{}
		}

		for j := range clients[i].Companies {
			if clients[i].Companies[j].Workers == nil {
				clients[i].Companies[j].Workers = []models.Worker{}
			}
		}
	}

	return c.JSON(clients)
}

func GetClient(c *fiber.Ctx) error {
	id := c.Params("id")

	var client models.Client

	if err := database.DB.Preload("Companies.Workers").First(&client, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "Cliente non trovato",
		})
	}

	if client.Companies == nil {
		client.Companies = []models.Company{}
	}

	for i := range client.Companies {
		if client.Companies[i].Workers == nil {
			client.Companies[i].Workers = []models.Worker{}
		}
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

	tx := database.DB.Begin()

	if err := tx.Where("company_id IN (?)",
		tx.Model(&models.Company{}).Select("id").Where("client_id = ?", id),
	).Delete(&models.Worker{}).Error; err != nil {
		tx.Rollback()
		return c.Status(500).JSON(fiber.Map{
			"error": "Errore eliminazione lavoratori collegati",
		})
	}

	if err := tx.Where("client_id = ?", id).Delete(&models.Company{}).Error; err != nil {
		tx.Rollback()
		return c.Status(500).JSON(fiber.Map{
			"error": "Errore eliminazione aziende collegate",
		})
	}

	if err := tx.Delete(&client).Error; err != nil {
		tx.Rollback()
		return c.Status(500).JSON(fiber.Map{
			"error": "Errore eliminazione cliente",
		})
	}

	tx.Commit()

	return c.JSON(fiber.Map{
		"message": "Cliente, aziende e lavoratori collegati eliminati correttamente",
	})
}
