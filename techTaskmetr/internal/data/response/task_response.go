package response

import (
	"github.com/MentalMentos/techTaskMetr.git/internal/models"
	"time"
)

type TaskResponse struct {
	ID          int64     `json:"id"`          // Уникальный идентификатор задачи
	Title       string    `json:"title"`       // Название задачи
	Description string    `json:"description"` // Описание задачи
	CreatedAt   time.Time `json:"created_at"`  // Дата и время создания задачи
	Status      string    `json:"status"`      //0-незавершённая 1-завершённая
}

type AllTasksResponse struct {
	Tasks []models.Task `json:"tasks"`
}
