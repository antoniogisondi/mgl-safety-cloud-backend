package models

import "time"

type User struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Name     string `gorm:"not null" json:"name"`
	Surname  string `json:"surname"`
	Email    string `gorm:"unique;not null" json:"email"`
	Password string `gorm:"not null" json:"-"`
	Role     string `gorm:"default:user" json:"role"`
	UserType string `gorm:"not null" json:"user_type"` // persona_fisica | consulente | admin
	Phone    string `json:"phone"`
	Notes    string `json:"notes"`

	Companies []Company `gorm:"foreignKey:UserID" json:"companies,omitempty"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
