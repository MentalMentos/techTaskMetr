package request

type CreateTaskRequest struct {
	Title       string `json:"title" binding:"required"`       // Название задачи (обязательно)
	Description string `json:"description" binding:"required"` // Описание задачи (обязательно)
}
