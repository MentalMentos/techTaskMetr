package request

type CreateTaskRequest struct {
	User_id     int    `json:"user_id" binding:"required"`
	Title       string `json:"title" validate:"required"`       // Название задачи (обязательно)
	Description string `json:"description" validate:"required"` // Описание задачи (обязательно)
}

type UpdateTaskRequest struct {
	User_id     int    `json:"user_id" binding:"required"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type DeleteTaskRequest struct {
	User_id     int    `json:"user_id" binding:"required"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type GetTaskRequest struct {
	User_id int `json:"user_id" binding:"required"`
}
