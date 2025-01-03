package request

type CreateTaskRequest struct {
	Title       string `json:"title" validate:"required"`       // Название задачи (обязательно)
	Description string `json:"description" validate:"required"` // Описание задачи (обязательно)
}

type UpdateTaskRequest struct {
	Id          string `json:"id" binding:"required"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type DeleteTaskRequest struct {
	Id          string `json:"id" binding:"required"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Status      string `json:"status" binding:"required"`
}

type GetTaskRequest struct {
	Id          string `json:"id" binding:"required"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Status      string `json:"status" binding:"required"`
}
