package response

import (
	"time"
)

type TaskResponse struct {
	ID          int       `gorm:"id" json:"id"`                   // Уникальный идентификатор задачи
	Title       string    `gorm:"title" json:"title"`             // Название задачи
	Description string    `gorm:"description" json:"description"` // Описание задачи
	CreatedAt   time.Time `gorm:"created_at" json:"created_at"`   // Дата и время создания задачи
	Status      string    `gorm:"status" json:"status"`           //0-незавершённая 1-завершённая
	UserID      int       `gorm:"user_id" json:"user_id"`
}

type AllTasksResponse struct {
	Tasks []TaskResponse `json:"tasks"`
}
