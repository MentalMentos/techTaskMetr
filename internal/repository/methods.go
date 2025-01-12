package repository

import (
	"github.com/MentalMentos/techTaskMetr.git/internal/data/request"
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

func (r *RepoImpl) Create(ctx *gin.Context, m *models.Task) error {
	if err := r.DB.WithContext(ctx).Create(&m).Error; err != nil {
		r.logger.Fatal("[  Repository  ]", helpers.FailedToCreateElement)
		return err
	}
	r.logger.Info("[  Repository  ]", helpers.Success)
	return nil
}

func (r *RepoImpl) Delete(ctx *gin.Context, m *models.Task) error {
	if err := r.DB.WithContext(ctx).Delete(&m).Error; err != nil {
		r.logger.Fatal("[  Repository  ]", helpers.FailedToDeleteElement)
		return err
	}
	r.logger.Info("[  Repository  ]", helpers.Success)
	return nil
}

// показывает все запланированные таски
func (r *RepoImpl) List(ctx *gin.Context) ([]models.Task, error) {
	var tasks []models.Task
	if err := r.DB.WithContext(ctx).Find(&tasks, "status = ?", true).Error; err != nil {
		r.logger.Fatal("[  Repository  ]", helpers.FailedToGetElements)
		return tasks, err
	}
	r.logger.Info("[  Repository  ]", helpers.Success)
	return tasks, nil
}

func (r *RepoImpl) Update(ctx *gin.Context, m *models.Task) error {
	newTask := request.UpdateTaskRequest{
		Id:    m.Title,
		Title: m.Description,
	}
	if err := r.DB.WithContext(ctx).Updates(newTask).Error; err != nil {
		r.logger.Fatal("[ REPOSITORY_UPDATE ]", "FailedToUpdateElement")
		return err
	}
	r.logger.Info("[  Repository  ]", helpers.Success)
	return nil
}

func (r *RepoImpl) GetById(ctx *gin.Context, id string) (models.Task, error) {
	var m models.Task
	if err := r.DB.WithContext(ctx).First(&m, "id = ?", id).Error; err != nil {
		r.logger.Fatal("[  REPOSITORY  ]", helpers.FailedToGetElements)
		return m, err
	}
	r.logger.Info("[  Repository  ]", helpers.Success)
	return m, nil
}
