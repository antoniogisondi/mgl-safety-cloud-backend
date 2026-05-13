package models

import "time"

type DeadlineCategory struct {
	ID             uint   `gorm:"primaryKey" json:"id"`
	Name           string `gorm:"not null" json:"name"`
	Group          string `json:"group"`
	Description    string `json:"description"`
	ValidityMonths *int   `json:"validity_months"` // nil = nessuna scadenza

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
