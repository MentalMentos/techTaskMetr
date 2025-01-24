package repository

import (
	"fmt"
	"github.com/MentalMentos/techTaskMetr/techTaskmetr/internal/data/request"
	"github.com/MentalMentos/techTaskMetr/techTaskmetr/internal/data/response"
	"github.com/MentalMentos/techTaskMetr/techTaskmetr/internal/models"
	"github.com/MentalMentos/techTaskMetr/techTaskmetr/pkg/helpers"
	"github.com/MentalMentos/techTaskMetr/techTaskmetr/pkg/logger"

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
		r.logger.Debug("[  Repository  ]", helpers.FailedToCreateElement)
		return err
	}
	r.logger.Info("[  Repository  ]", helpers.Success)
	return nil
}

func (r *RepoImpl) Delete(ctx *gin.Context, m *models.Task) error {
	if err := r.DB.WithContext(ctx).Delete(&m).Error; err != nil {
		r.logger.Debug("[  Repository  ]", helpers.FailedToDeleteElement)
		return err
	}
	r.logger.Info("[  Repository  ]", helpers.Success)
	return nil
}

// показывает все запланированные таски
func (r *RepoImpl) List(ctx *gin.Context, user_id int) ([]response.TaskResponse, error) {
	var tasks []response.TaskResponse
	if err := r.DB.WithContext(ctx).Where("user_id = ?", user_id).Find(&tasks).Error; err != nil {
		r.logger.Debug("[  Repository_list  ]", helpers.FailedToGetElements)
		return tasks, err
	}
	log_task := fmt.Sprintf("%v", tasks)
	r.logger.Info("[  Repository_list ]", log_task)
	r.logger.Info("[  Repository_list  ]", helpers.Success)
	return tasks, nil
}

func (r *RepoImpl) Update(ctx *gin.Context, m *models.Task) error {
	newTask := request.UpdateTaskRequest{
		User_id:     m.UserID,
		Title:       m.Title,
		Description: m.Description,
	}
	if err := r.DB.WithContext(ctx).Where("user_id = ?", m.UserID).Updates(newTask).Error; err != nil {
		r.logger.Debug("[ REPOSITORY_UPDATE ]", "FailedToUpdateElement")
		return err
	}
	r.logger.Info("[  Repository  ]", helpers.Success)
	return nil
}

func (r *RepoImpl) GetByTitle(ctx *gin.Context, title string, user_id int) (models.Task, error) {
	var m models.Task
	if err := r.DB.WithContext(ctx).First(&m, "title = ? AND user_id = ?", title, user_id).Error; err != nil {
		r.logger.Debug("[  REPOSITORY  ]", helpers.FailedToGetElements)
		return m, err
	}
	r.logger.Info("[  Repository  ]", helpers.Success)
	return m, nil
}

func (r *RepoImpl) GetByID(ctx *gin.Context, id string, user_id int) (models.Task, error) {
	var m models.Task
	if err := r.DB.WithContext(ctx).First(&m, "id = ? AND user_id = ?", id, user_id).Error; err != nil {
		r.logger.Debug("[  REPOSITORY  ]", helpers.FailedToGetElements)
		return m, err
	}
	r.logger.Info("[  Repository  ]", helpers.Success)
	return m, nil
}
