package dashboard

import (
	"time"

	"github.com/gofiber/fiber/v2"

	"github.com/antoniogisondi/mgl-safety-cloud-backend/internal/database"
	"github.com/antoniogisondi/mgl-safety-cloud-backend/internal/models"
)

func GetStats(c *fiber.Ctx) error {
	var clientsCount int64
	var companiesCount int64
	var workersCount int64
	var deadlinesCount int64
	var expiredCount int64
	var expiringCount int64
	var validCount int64

	now := time.Now()
	limit := now.AddDate(0, 0, 30)

	database.DB.Model(&models.Company{}).Count(&companiesCount)
	database.DB.Model(&models.Worker{}).Count(&workersCount)
	database.DB.Model(&models.Deadline{}).Count(&deadlinesCount)

	database.DB.Model(&models.Deadline{}).
		Where("expiration_date < ?", now).
		Count(&expiredCount)

	database.DB.Model(&models.Deadline{}).
		Where("expiration_date >= ? AND expiration_date <= ?", now, limit).
		Count(&expiringCount)

	database.DB.Model(&models.Deadline{}).
		Where("expiration_date > ?", limit).
		Count(&validCount)

	return c.JSON(fiber.Map{
		"clients":            clientsCount,
		"companies":          companiesCount,
		"workers":            workersCount,
		"deadlines":          deadlinesCount,
		"expired_deadlines":  expiredCount,
		"expiring_deadlines": expiringCount,
		"valid_deadlines":    validCount,
	})
}
