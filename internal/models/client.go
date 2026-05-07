package models

import "time"

type Client struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	Name       string    `gorm:"not null" json:"name"`
	VatNumber  string    `gorm:"unique" json:"vat_number"`
	FiscalCode string    `json:"fiscal_code"`
	Email      string    `json:"email"`
	Phone      string    `json:"phone"`
	Address    string    `json:"address"`
	City       string    `json:"city"`
	Province   string    `json:"province"`
	PostalCode string    `json:"postal_code"`
	Notes      string    `json:"notes"`
	Companies  []Company `json:"companies"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
