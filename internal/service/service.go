package service

import (
	"github.com/MentalMentos/techTaskMetr.git/internal/data/request"
	"github.com/MentalMentos/techTaskMetr.git/internal/data/response"
	"github.com/MentalMentos/techTaskMetr.git/internal/models"
	"github.com/MentalMentos/techTaskMetr.git/internal/repository"
	"github.com/MentalMentos/techTaskMetr.git/pkg/logger"
	"github.com/gin-gonic/gin"
)

type Task interface {
	Create(ctx *gin.Context, req request.CreateTaskRequest) (*response.TaskResponse, error)
	Update(ctx *gin.Context, req request.UpdateTaskRequest) (*response.TaskResponse, error)
	Done(ctx *gin.Context, task models.Task) error
}

type Service struct {
	Task
}

func NewService(repo *repository.Repository, logger logger.Logger) *Service {
	return &Service{
		NewTaskService(repo, logger),
	}
}
