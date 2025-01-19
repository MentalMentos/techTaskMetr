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

func (s *TaskService) Create(ctx *gin.Context, req request.CreateTaskRequest) (*response.TaskResponse, error) {
	task := &models.Task{
		UserID:      req.User_id,
		Title:       req.Title,
		Description: req.Description,
		CreatedAt:   time.Now(),
		Status:      "false",
	}
	err := s.repo.Create(ctx, task)
	if err != nil {
		s.logger.Fatal(helpers.FailedToCreateElement, "failed to create element in service")
		return nil, err
	}
	s.logger.Info(helpers.InfoPrefix, "Service created new user")
	return &response.TaskResponse{
		task.ID,
		task.Title,
		task.Description,
		task.CreatedAt,
		task.Status,
	}, nil
}

func (s *TaskService) Update(ctx *gin.Context, req request.UpdateTaskRequest) (*response.TaskResponse, error) {
	task, err := s.repo.GetByTitle(ctx, req.Title, req.User_id)
	if err != nil {
		s.logger.Fatal(helpers.FailedToUpdateElement, "failed to update element in service(get by title)")
	}
	task2 := &models.Task{
		UserID:      req.User_id,
		Title:       req.Title,
		Description: req.Description,
	}

	err = s.repo.Update(ctx, task2)
	if err != nil {
		s.logger.Fatal("[ SERVICE_UPDATE ]", "failed to update element in service")
		return nil, err
	}

	s.logger.Info(helpers.InfoPrefix, "Service updated new user")
	return &response.TaskResponse{
		Title:       task.Title,
		Description: task.Description,
		CreatedAt:   task.CreatedAt,
	}, nil
}

func (s *TaskService) Done(ctx *gin.Context, req request.DeleteTaskRequest) (*response.TaskResponse, error) {
	task, err := s.repo.GetByTitle(ctx, req.Title, req.User_id)
	if err != nil {
		s.logger.Fatal("[ SERVICE_DONE ]", "failed to get element in service by id")
		return nil, err
	}
	err = s.repo.Delete(ctx, &task)
	if err != nil {
		s.logger.Fatal(helpers.FailedToDeleteElement, "failed to delete element in service")
		return nil, err
	}
	s.logger.Info(helpers.InfoPrefix, "task done")
	return &response.TaskResponse{
		ID:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		CreatedAt:   task.CreatedAt,
		Status:      "true",
	}, nil
}

func (s *TaskService) List(ctx *gin.Context, user_id int64) (*response.AllTasksResponse, error) {
	tasks, err := s.repo.List(ctx, user_id)
	if err != nil {
		s.logger.Fatal("[ SERVICE_LIST ]", helpers.FailedToGetElements)
		return nil, err
	}
	s.logger.Info(helpers.InfoPrefix, "service list")
	return &response.AllTasksResponse{
		tasks,
	}, nil

}
