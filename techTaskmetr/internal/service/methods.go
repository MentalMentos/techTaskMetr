package service

import (
	"context"
	"github.com/MentalMentos/techTaskMetr/api_gateway/pkg/helpers"
	"github.com/MentalMentos/techTaskMetr/techTaskmetr/internal/data/request"
	"github.com/MentalMentos/techTaskMetr/techTaskmetr/internal/data/response"
	"github.com/MentalMentos/techTaskMetr/techTaskmetr/internal/models"
	"github.com/MentalMentos/techTaskMetr/techTaskmetr/internal/repository"
	"github.com/MentalMentos/techTaskMetr/techTaskmetr/pkg/logger"
	"github.com/pkg/errors"
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

func (s *TaskService) Create(ctx context.Context, req request.CreateTaskRequest) (*response.TaskResponse, error) {
	task := &models.Task{
		UserID:      req.User_id,
		Title:       req.Title,
		Description: req.Description,
		CreatedAt:   time.Now(),
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
		task.UserID,
	}, nil
}

func (s *TaskService) Update(ctx context.Context, req request.UpdateTaskRequest) (*response.TaskResponse, error) {
	task, err := s.repo.GetByTitle(ctx, req.Title, req.User_id)
	if err != nil {
		s.logger.Fatal(helpers.FailedToUpdateElement, "failed to update element in service(get by title)")
	}
	if task.UserID != req.User_id {
		return nil, errors.New("user id not match")
	}
	if task.Title != req.Title {
		return nil, errors.New("title not match")
	}
	if task.Title == "" || task.Description == "" {
		return nil, errors.New("title or description is null")
	}
	task2 := &models.Task{
		UserID:      task.UserID,
		Title:       task.Title,
		Description: task.Description,
	}

	err = s.repo.Update(ctx, task2)
	if err != nil {
		s.logger.Fatal("[ SERVICE_UPDATE ]", "failed to update element in service")
		return nil, err
	}

	s.logger.Info(helpers.InfoPrefix, "Service updated new user")
	return &response.TaskResponse{
		Title:       task2.Title,
		Description: task2.Description,
		CreatedAt:   task2.CreatedAt,
	}, nil
}

func (s *TaskService) Done(ctx context.Context, req request.DeleteTaskRequest) (*response.TaskResponse, error) {
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
	}, nil
}

func (s *TaskService) List(ctx context.Context, user_id int) (*response.AllTasksResponse, error) {
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
