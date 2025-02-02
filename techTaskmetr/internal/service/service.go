package service

import (
	"github.com/MentalMentos/techTaskMetr/techTaskmetr/internal/data/request"
	"github.com/MentalMentos/techTaskMetr/techTaskmetr/internal/data/response"
	"github.com/MentalMentos/techTaskMetr/techTaskmetr/internal/repository"
	"github.com/MentalMentos/techTaskMetr/techTaskmetr/pkg/logger"
	"github.com/gin-gonic/gin"
)

type Task interface {
	Create(ctx *gin.Context, req request.CreateTaskRequest) (*response.TaskResponse, error)
	Update(ctx *gin.Context, req request.UpdateTaskRequest) (*response.TaskResponse, error)
	Done(ctx *gin.Context, req request.DeleteTaskRequest) (*response.TaskResponse, error)
	List(ctx *gin.Context, user_id int) (*response.AllTasksResponse, error)
}

type Service struct {
	Task
}

func NewService(repo *repository.Repository, logger logger.Logger) *Service {
	return &Service{
		Task: NewTaskService(repo, logger),
	}
}
