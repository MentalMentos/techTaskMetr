package models

import "time"

type Task struct {
	ID          int64     `gorm:"primaryKey" json:"id"`
	Title       string    `gorm:"title" json:"title"`
	Description string    `gorm:"description" json:"description"`
	CreatedAt   time.Time `gorm:"created_at" json:"created_at"`
	Status      bool      `gorm:"status" json:"status"`
}
