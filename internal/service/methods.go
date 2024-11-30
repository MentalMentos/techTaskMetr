package service

import (
	"github.com/MentalMentos/techTaskMetr.git/internal/data/request"
	"github.com/MentalMentos/techTaskMetr.git/internal/data/response"
	"github.com/MentalMentos/techTaskMetr.git/internal/models"
	"github.com/gin-gonic/gin"
	"time"
)

type TaskService interface {
	Create(ctx gin.Context, req request.CreateTaskRequest) (*response.Response, error)
}

func (s *Service) Create(ctx *gin.Context, req request.CreateTaskRequest) (*response.TaskResponse, error) {
	task := &models.Task{
		Title:       req.Title,
		Description: req.Description,
		CreatedAt:   time.Now(),
	}
	err := s.repo.Create(ctx, task)
	return &response.TaskResponse{
		task.ID,
		task.Title,
		task.Description,
		task.CreatedAt,
	}, err
}
