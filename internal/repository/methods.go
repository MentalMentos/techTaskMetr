package repository

import (
	"github.com/MentalMentos/techTaskMetr.git/internal/models"
	"github.com/MentalMentos/techTaskMetr.git/pkg/helpers"
	"github.com/MentalMentos/techTaskMetr.git/pkg/logger"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RepoImpl struct {
	DB     *gorm.DB
	logger logger.Logger
}

func NewTaskRepo(DB *gorm.DB, logger logger.Logger) *RepoImpl {
	return &RepoImpl{
		DB:     DB,
		logger: logger,
	}
}

func (r *RepoImpl) Create(ctx *gin.Context, m *models.Task, logger logger.Logger) error {
	if err := r.DB.WithContext(ctx).Create(&m).Error; err != nil {
		logger.Fatal("[  Repository  ]", helpers.FailedToCreateElement)
		return err
	}
	logger.Info("[  Repository  ]", helpers.Success)
	return nil
}
