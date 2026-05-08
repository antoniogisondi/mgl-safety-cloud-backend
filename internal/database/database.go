package database

import (
	"log"
	"os"

	"github.com/antoniogisondi/mgl-safety-cloud-backend/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := os.Getenv("DATABASE_URL")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Errore connessione database: ", err)
	}

	DB = db

	log.Println("Database Supabase connesso correttamente")
}

func Migrate() {

	err := DB.AutoMigrate(
		&models.User{},
		&models.Client{},
		&models.Company{},
		&models.Worker{},
		&models.Deadline{},
		&models.DeadlineCategory{},
		&models.Notification{},
	)

	if err != nil {
		log.Fatal("Errore migrazione database: ", err)
	}

	log.Println("Migrazione database completata")
}
