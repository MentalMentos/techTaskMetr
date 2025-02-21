package service

import (
	"context"
	"github.com/MentalMentos/techTaskMetr/techTaskmetr/internal/data/request"
	"github.com/MentalMentos/techTaskMetr/techTaskmetr/internal/data/response"
	"github.com/MentalMentos/techTaskMetr/techTaskmetr/internal/repository"
	"github.com/MentalMentos/techTaskMetr/techTaskmetr/pkg/logger"
)

type Task interface {
	Create(ctx context.Context, req request.CreateTaskRequest) (*response.TaskResponse, error)
	Update(ctx context.Context, req request.UpdateTaskRequest) (*response.TaskResponse, error)
	Done(ctx context.Context, req request.DeleteTaskRequest) (*response.TaskResponse, error)
	List(ctx context.Context, user_id int) (*response.AllTasksResponse, error)
}

type Service struct {
	Task
}

func NewService(repo *repository.Repository, logger logger.Logger) *Service {
	return &Service{
		Task: NewTaskService(repo, logger),
	}
}
