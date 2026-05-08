package models

import "time"

type Notification struct {
	ID      uint   `gorm:"primaryKey" json:"id"`
	Title   string `gorm:"not null" json:"title"`
	Message string `gorm:"not null" json:"message"`
	Type    string `json:"type"`
	IsRead  bool   `gorm:"default:false" json:"is_read"`

	DeadlineID *uint     `json:"deadline_id"`
	Deadline   *Deadline `gorm:"foreignKey:DeadlineID" json:"-"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
