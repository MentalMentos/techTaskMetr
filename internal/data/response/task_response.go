package response

import "time"

type TaskResponse struct {
	ID          int64     `json:"id"`          // Уникальный идентификатор задачи
	Title       string    `json:"title"`       // Название задачи
	Description string    `json:"description"` // Описание задачи
	CreatedAt   time.Time `json:"created_at"`  // Дата и время создания задачи
	Status      bool      `json:"status"`      //0-незавершённая 1-завершённая
}

type AllTasksResponse struct {
	Tasks []TaskResponse `json:"tasks"`
}
