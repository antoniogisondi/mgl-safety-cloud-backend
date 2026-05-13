package models

import "time"

type Company struct {
	ID             uint   `gorm:"primaryKey" json:"id"`
	UserID         uint   `gorm:"not null" json:"user_id"`
	RagioneSociale string `gorm:"not null" json:"ragione_sociale"`
	VatNumber      string `json:"vat_number"`
	FiscalCode     string `json:"fiscal_code"`
	Address        string `json:"address"`
	City           string `json:"city"`
	Province       string `json:"province"`
	PostalCode     string `json:"postal_code"`
	ATECO          string `json:"ateco"`
	Attivita       string `json:"attivita"`
	Notes          string `json:"notes"`

	User      User       `gorm:"foreignKey:UserID" json:"-"`
	Workers   []Worker   `gorm:"foreignKey:CompanyID" json:"workers,omitempty"`
	Deadlines []Deadline `gorm:"foreignKey:CompanyID" json:"deadlines,omitempty"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
