package notifications

import (
	"time"

	"github.com/gofiber/fiber/v2"

	"github.com/antoniogisondi/mgl-safety-cloud-backend/internal/database"
	"github.com/antoniogisondi/mgl-safety-cloud-backend/internal/models"
)

func GetNotifications(c *fiber.Ctx) error {
	var notifications []models.Notification

	if err := database.DB.Order("created_at DESC").Find(&notifications).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Errore recupero notifiche"})
	}

	return c.JSON(notifications)
}

func MarkAsRead(c *fiber.Ctx) error {
	id := c.Params("id")

	var notification models.Notification

	if err := database.DB.First(&notification, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Notifica non trovata"})
	}

	notification.IsRead = true

	if err := database.DB.Save(&notification).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Errore aggiornamento notifica"})
	}

	return c.JSON(notification)
}

func DeleteNotification(c *fiber.Ctx) error {
	id := c.Params("id")

	var notification models.Notification

	if err := database.DB.First(&notification, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Notifica non trovata"})
	}

	if err := database.DB.Delete(&notification).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Errore eliminazione notifica"})
	}

	return c.JSON(fiber.Map{"message": "Notifica eliminata correttamente"})
}

func GenerateDeadlineAlerts(c *fiber.Ctx) error {
	var deadlines []models.Deadline

	if err := database.DB.
		Preload("Category").
		Find(&deadlines).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Errore recupero scadenze",
		})
	}

	created := 0

	for _, deadline := range deadlines {
		status := calculateNotificationStatus(deadline.ExpirationDate)

		if status == "" {
			continue
		}

		var existing models.Notification

		err := database.DB.
			Where("deadline_id = ? AND type = ?", deadline.ID, status).
			First(&existing).Error

		if err == nil {
			continue
		}

		notification := models.Notification{
			Title:      "Avviso scadenza",
			Message:    "La scadenza \"" + deadline.Title + "\" risulta " + status,
			Type:       status,
			IsRead:     false,
			DeadlineID: &deadline.ID,
		}

		if err := database.DB.Create(&notification).Error; err == nil {
			created++
		}
	}

	return c.JSON(fiber.Map{
		"message": "Generazione notifiche completata",
		"created": created,
	})
}

func calculateNotificationStatus(expirationDate *time.Time) string {
	if expirationDate == nil {
		return ""
	}

	now := time.Now()
	daysLeft := int(expirationDate.Sub(now).Hours() / 24)

	if daysLeft < 0 {
		return "deadline_expired"
	}

	if daysLeft <= 30 {
		return "deadline_expiring"
	}

	return ""
}
