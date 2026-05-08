package scheduler

import (
	"log"
	"time"

	"github.com/go-co-op/gocron/v2"

	"github.com/antoniogisondi/mgl-safety-cloud-backend/internal/database"
	"github.com/antoniogisondi/mgl-safety-cloud-backend/internal/models"
)

func StartScheduler() {
	s, err := gocron.NewScheduler()
	if err != nil {
		log.Println("Errore creazione scheduler:", err)
		return
	}

	_, err = s.NewJob(
		gocron.DurationJob(24*time.Hour),
		gocron.NewTask(generateDeadlineAlerts),
	)

	if err != nil {
		log.Println("Errore creazione job notifiche:", err)
		return
	}

	s.Start()

	log.Println("Scheduler notifiche avviato")
}

func generateDeadlineAlerts() {
	var deadlines []models.Deadline

	if err := database.DB.Preload("Category").Find(&deadlines).Error; err != nil {
		log.Println("Errore recupero scadenze scheduler:", err)
		return
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
			Message:    "La scadenza \"" + deadline.Title + "\" richiede attenzione.",
			Type:       status,
			IsRead:     false,
			DeadlineID: &deadline.ID,
		}

		if err := database.DB.Create(&notification).Error; err == nil {
			created++
		}
	}

	log.Println("Scheduler notifiche completato. Create:", created)
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
