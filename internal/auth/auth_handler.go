package auth

import (
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"

	"github.com/antoniogisondi/mgl-safety-cloud-backend/internal/database"
	"github.com/antoniogisondi/mgl-safety-cloud-backend/internal/models"
)

type RegisterRequest struct {
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
	UserType string `json:"user_type"`
	Phone    string `json:"phone"`
	Notes    string `json:"notes"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Register(c *fiber.Ctx) error {
	var body RegisterRequest

	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Dati non validi"})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Errore creazione password"})
	}

	user := models.User{
		Name:     body.Name,
		Surname:  body.Surname,
		Email:    body.Email,
		Password: string(hashedPassword),
		UserType: body.UserType,
		Phone:    body.Phone,
		Notes:    body.Notes,
		Role:     "clients",
	}

	if err := database.DB.Create(&user).Error; err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Utente già esistente o dati non validi"})
	}

	return c.Status(201).JSON(fiber.Map{
		"message": "Utente creato correttamente",
		"user":    user,
	})
}

func Login(c *fiber.Ctx) error {
	var body LoginRequest
	var user models.User

	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Dati non validi"})
	}

	if body.Email == "" || body.Password == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Email e password sono obbligatorie"})
	}

	if err := database.DB.Where("email = ?", body.Email).First(&user).Error; err != nil {
		return c.Status(401).JSON(fiber.Map{"error": "Credenziali non valide"})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)); err != nil {
		return c.Status(401).JSON(fiber.Map{"error": "Credenziali non valide"})
	}

	claims := jwt.MapClaims{
		"user_id":   user.ID,
		"email":     user.Email,
		"role":      user.Role,
		"user_type": user.UserType,
		"exp":       time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Errore generazione token"})
	}

	return c.JSON(fiber.Map{
		"message": "Login effettuato",
		"token":   tokenString,
		"user": fiber.Map{
			"id":        user.ID,
			"name":      user.Name,
			"surname":   user.Surname,
			"email":     user.Email,
			"role":      user.Role,
			"user_type": user.UserType,
			"phone":     user.Phone,
			"notes":     user.Notes,
		},
	})
}
