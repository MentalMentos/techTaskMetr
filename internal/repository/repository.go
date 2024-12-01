package repository

import (
	"github.com/MentalMentos/techTaskMetr.git/internal/models"
	"github.com/MentalMentos/techTaskMetr.git/pkg/logger"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Task interface {
	Create(ctx gin.Context, m *models.Task, logger logger.Logger) error
}

type Repository struct {
	Task
}

func New(db *gorm.DB, myLogger logger.Logger) *Repository {
	return &Repository{NewTaskRepo(db, myLogger)}
}
