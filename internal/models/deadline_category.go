package models

import "time"

type DeadlineCategory struct {
	ID             uint      `gorm:"primaryKey" json:"id"`
	Name           string    `gorm:"unique;not null" json:"name"`
	Group          string    `gorm:"not null" json:"group"`
	Description    string    `json:"description"`
	ValidityMonths int       `json:"validity_months"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
