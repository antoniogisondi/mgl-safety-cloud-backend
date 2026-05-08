package models

import "time"

type Deadline struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	Title       string `gorm:"not null" json:"title"`
	Description string `json:"description"`

	CategoryID uint             `gorm:"not null" json:"category_id"`
	Category   DeadlineCategory `gorm:"foreignKey:CategoryID" json:"category"`

	IssueDate      *time.Time `json:"issue_date"`
	ExpirationDate *time.Time `json:"expiration_date"`
	Status         string     `gorm:"default:valid" json:"status"`
	Notes          string     `json:"notes"`

	CompanyID uint    `gorm:"not null" json:"company_id"`
	Company   Company `gorm:"foreignKey:CompanyID" json:"-"`

	WorkerID *uint   `json:"worker_id"`
	Worker   *Worker `gorm:"foreignKey:WorkerID" json:"-"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
