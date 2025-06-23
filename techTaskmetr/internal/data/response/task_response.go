package response

import (
	"time"
)

type TaskResponse struct {
	ID          int       `gorm:"id" json:"id"`                   // Уникальный идентификатор задачи
	Title       string    `gorm:"title" json:"title"`             // Название задачи
	Description string    `gorm:"description" json:"description"` // Описание задачи
	CreatedAt   time.Time `gorm:"created_at" json:"created_at"`   // Дата и время создания задачи
	UserID      int       `gorm:"user_id" json:"user_id"`
}

// Указываем GORM, что эта структура соответствует таблице tasks
func (TaskResponse) TableName() string {
	return "tasks"
}

type AllTasksResponse struct {
	Tasks []TaskResponse `json:"tasks"`
}
