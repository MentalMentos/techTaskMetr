package repository

import (
	"github.com/MentalMentos/techTaskMetr.git/internal/models"
	"github.com/MentalMentos/techTaskMetr.git/pkg/logger"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Task interface {
	Create(ctx *gin.Context, m *models.Task) error
	Update(ctx *gin.Context, m *models.Task) error
	Delete(ctx *gin.Context, m *models.Task) error
	List(ctx *gin.Context, user_id int64) ([]models.Task, error)
	GetByTitle(ctx *gin.Context, title string, user_id int64) (models.Task, error)
	GetByID(ctx *gin.Context, id string, user_id int64) (models.Task, error)
}

type Repository struct {
	Task
}

func New(db *gorm.DB, myLogger logger.Logger) *Repository {
	return &Repository{NewTaskRepo(db, myLogger)}
}
