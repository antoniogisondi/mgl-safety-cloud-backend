package models

import "time"

type Deadline struct {
	ID         uint  `gorm:"primaryKey" json:"id"`
	CompanyID  uint  `gorm:"not null" json:"company_id"`
	WorkerID   *uint `json:"worker_id"`
	CategoryID uint  `gorm:"not null" json:"category_id"`

	Title          string     `gorm:"not null" json:"title"`
	Description    string     `json:"description"`
	IssueDate      *time.Time `json:"issue_date"`
	ExpirationDate *time.Time `json:"expiration_date"`
	Status         string     `gorm:"default:valid" json:"status"`
	Notes          string     `json:"notes"`

	Company  Company          `gorm:"foreignKey:CompanyID" json:"-"`
	Worker   *Worker          `gorm:"foreignKey:WorkerID" json:"worker,omitempty"`
	Category DeadlineCategory `gorm:"foreignKey:CategoryID" json:"category,omitempty"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
