package service

import (
	"github.com/MentalMentos/techTaskMetr.git/internal/data/request"
	"github.com/MentalMentos/techTaskMetr.git/internal/data/response"
	"github.com/MentalMentos/techTaskMetr.git/internal/models"
	"github.com/MentalMentos/techTaskMetr.git/internal/repository"
	"github.com/MentalMentos/techTaskMetr.git/pkg/helpers"
	"github.com/MentalMentos/techTaskMetr.git/pkg/logger"
	"github.com/gin-gonic/gin"
	"time"
)

type TaskService struct {
	repo   *repository.Repository
	logger logger.Logger
}

func NewTaskService(repo *repository.Repository, logger logger.Logger) *TaskService {
	return &TaskService{
		repo:   repo,
		logger: logger,
	}
}

func (s *TaskService) Create(ctx gin.Context, req request.CreateTaskRequest, logger logger.Logger) (*response.TaskResponse, error) {
	task := &models.Task{
		Title:       req.Title,
		Description: req.Description,
		CreatedAt:   time.Now(),
	}
	err := s.repo.Create(ctx, task, logger)
	if err != nil {
		logger.Fatal(helpers.FailedToCreateElement, "failed to create element in service")
		return nil, err
	}
	logger.Info(helpers.InfoPrefix, "Service created new user")
	return &response.TaskResponse{
		task.ID,
		task.Title,
		task.Description,
		task.CreatedAt,
	}, err
}
