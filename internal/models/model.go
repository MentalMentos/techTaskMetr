package models

import "time"

type Task struct {
	ID          int64     `gorm:"primaryKey" json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	Status      bool      `json:"status"`
}
