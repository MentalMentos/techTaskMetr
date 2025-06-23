package repository

import (
	"context"
	"github.com/MentalMentos/techTaskMetr/techTaskmetr/internal/clients/redis"
	"github.com/MentalMentos/techTaskMetr/techTaskmetr/internal/data/response"
	"github.com/MentalMentos/techTaskMetr/techTaskmetr/internal/models"
	"github.com/MentalMentos/techTaskMetr/techTaskmetr/pkg/logger"
	"github.com/jackc/pgx/v4"
	"gorm.io/gorm"
)

type Task interface {
	Create(ctx context.Context, tx pgx.Tx, m *models.Task) error
	Update(ctx context.Context, m *models.Task) error
	Delete(ctx context.Context, m *models.Task) error
	List(ctx context.Context, user_id int) ([]response.TaskResponse, error)
	GetByTitle(ctx context.Context, title string, user_id int) (models.Task, error)
	GetByID(ctx context.Context, id string, user_id int) (models.Task, error)
}

type Repository struct {
	Task
}

func New(db *gorm.DB, myLogger logger.Logger, redisClient redis.IRedis) *Repository {
	return &Repository{NewTaskRepo(db, myLogger, redisClient)}
}
