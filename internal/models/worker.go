package models

import "time"

type Worker struct {
	ID         uint       `gorm:"primaryKey" json:"id"`
	CompanyID  uint       `gorm:"not null" json:"company_id"`
	FirstName  string     `gorm:"not null" json:"first_name"`
	LastName   string     `gorm:"not null" json:"last_name"`
	FiscalCode string     `json:"fiscal_code"`
	Email      string     `json:"email"`
	Phone      string     `json:"phone"`
	BirthDate  *time.Time `json:"birth_date"`
	Role       string     `json:"role"`
	Department string     `json:"department"`
	JobTitle   string     `json:"job_title"`
	HireDate   *time.Time `json:"hire_date"`
	Notes      string     `json:"notes"`

	Company   Company    `gorm:"foreignKey:CompanyID" json:"-"`
	Deadlines []Deadline `gorm:"foreignKey:WorkerID" json:"deadlines,omitempty"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
