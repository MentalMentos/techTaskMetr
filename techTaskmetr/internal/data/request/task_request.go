package request

type CreateTaskRequest struct {
	Title       string `json:"title" validate:"required"`       // Название задачи (обязательно)
	Description string `json:"description" validate:"required"` // Описание задачи (обязательно)
	Status      string `json:"status" validate:"required"`      //bool
}

type UpdateTaskRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Status      string `json:"status" binding:"required"`
}

type DeleteTaskRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Status      string `json:"status" binding:"required"`
}

type GetTaskRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Status      string `json:"status" binding:"required"`
}
