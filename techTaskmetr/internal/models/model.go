package models

import "time"

type Task struct {
	ID          int       `gorm:"id" json:"id"`
	Title       string    `gorm:"title" json:"title"`
	Description string    `gorm:"description" json:"description"`
	CreatedAt   time.Time `gorm:"created_at" json:"created_at"`
	Status      string    `gorm:"status" json:"status"`
	UserID      int64     `gorm:"user_id" json:"user_id"`
}
