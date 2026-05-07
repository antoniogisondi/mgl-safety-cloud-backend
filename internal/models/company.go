package models

import "time"

type Company struct {
	ID         uint   `gorm:"primaryKey" json:"id"`
	Name       string `gorm:"not null" json:"name"`
	VatNumber  string `json:"vat_number"`
	FiscalCode string `json:"fiscal_code"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	Address    string `json:"address"`
	City       string `json:"city"`
	Province   string `json:"province"`
	PostalCode string `json:"postal_code"`
	AtecoCode  string `json:"ateco_code"`
	Activity   string `json:"activity"`
	Notes      string `json:"notes"`

	ClientID uint   `gorm:"not null" json:"client_id"`
	Client   Client `gorm:"foreignKey:ClientID" json:"-"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
