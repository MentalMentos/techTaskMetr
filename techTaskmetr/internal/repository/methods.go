package repository

import (
	"context"
	"fmt"
	"github.com/MentalMentos/techTaskMetr/techTaskmetr/internal/clients/redis"
	"github.com/MentalMentos/techTaskMetr/techTaskmetr/internal/data/request"
	"github.com/MentalMentos/techTaskMetr/techTaskmetr/internal/data/response"
	"github.com/MentalMentos/techTaskMetr/techTaskmetr/internal/models"
	"github.com/MentalMentos/techTaskMetr/techTaskmetr/pkg/helpers"
	"github.com/MentalMentos/techTaskMetr/techTaskmetr/pkg/logger"
	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"time"
)

type RepoImpl struct {
	DB          *gorm.DB
	redisClient redis.IRedis
	logger      logger.Logger
}

func NewTaskRepo(DB *gorm.DB, logger logger.Logger, redis redis.IRedis) *RepoImpl {
	return &RepoImpl{
		DB:          DB,
		redisClient: redis,
		logger:      logger,
	}
}

func (r *RepoImpl) Create(ctx context.Context, tx pgx.Tx, m *models.Task) error {
	_, err := tx.Exec(ctx, "INSERT INTO tasks (title, description, user_id) VALUES ($1, $2);",
		m.Title, m.Description)
	if err != nil {
		// Логирование ошибки при создании транзакции в базе данных и откат транзакции
		r.logger.Info(helpers.InfoPrefix, helpers.FailedToCreateElement)
		return errors.Wrap(err, helpers)
	}

	// Создание ключа и объекта для сохранения в Redis
	transactionKey := fmt.Sprintf("transaction:%d:%d", m.UserID, time.Now().Unix())
	transactionData := map[string]interface{}{
		"title":       m.Title,
		"description": m.Description,
		"status":      "false",
		"user_id":     m.UserID,
	}

	err = r.redisClient.SetObject(ctx, transactionKey, transactionData, time.Hour*24)
	if err != nil {
		// Логирование ошибки при сохранении транзакции в кэш
		r.logger.Info(helpers.RepoPrefix, helpers.RepoCacheTransactionError)
		return errors.Wrap(err, helpers.RepoCacheTransactionError)
	}
	r.logger.Info("[  Repository  ]", helpers.Success)
	return nil
}

func (r *RepoImpl) Delete(ctx context.Context, m *models.Task) error {
	if err := r.DB.WithContext(ctx).Delete(&m).Error; err != nil {
		r.logger.Debug("[  Repository  ]", helpers.RepoDeleteError)
		return err
	}
	r.logger.Info("[  Repository  ]", helpers.Success)
	return nil
}

func (r *RepoImpl) List(ctx context.Context, user_id int) ([]response.TaskResponse, error) {
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

func (r *RepoImpl) Update(ctx context.Context, m *models.Task) error {
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

func (r *RepoImpl) GetByTitle(ctx context.Context, title string, user_id int) (models.Task, error) {
	var m models.Task
	if err := r.DB.WithContext(ctx).First(&m, "title = ? AND user_id = ?", title, user_id).Error; err != nil {
		r.logger.Debug("[  REPOSITORY  ]", helpers.FailedToGetElements)
		return m, err
	}
	r.logger.Info("[  Repository  ]", helpers.Success)
	return m, nil
}

func (r *RepoImpl) GetByID(ctx context.Context, id string, user_id int) (models.Task, error) {
	var m models.Task
	if err := r.DB.WithContext(ctx).First(&m, "id = ? AND user_id = ?", id, user_id).Error; err != nil {
		r.logger.Debug("[  REPOSITORY  ]", helpers.FailedToGetElements)
		return m, err
	}
	r.logger.Info("[  Repository  ]", helpers.Success)
	return m, nil
}
