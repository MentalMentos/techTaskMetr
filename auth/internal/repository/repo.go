package repository

import (
	"context"
	"github.com/MentalMentos/techTaskMetr/auth/internal/model"
	"github.com/MentalMentos/techTaskMetr/techTaskmetr/pkg/logger"

	"gorm.io/gorm"
)

type Repository interface {
	Create(ctx context.Context, us model.User, logger logger.Logger) (int, error)
	Update(ctx context.Context, us model.User, logger logger.Logger) (int, error)
	Delete(ctx context.Context, usId int, logger logger.Logger) error
	UpdatePassword(ctx context.Context, us model.User, hashPassword string, logger logger.Logger) (model.User, error)
	UpdateIP(ctx context.Context, us model.User, ip string, logger logger.Logger) (model.User, error)
	GetByEmail(ctx context.Context, email string, logger logger.Logger) (model.User, error)
	GetByID(ctx context.Context, userID int, logger logger.Logger) (model.User, error)
	GetAll(ctx context.Context, logger logger.Logger) ([]model.User, error)
}

type Repo struct {
	Repository
}

func NewRepository(db *gorm.DB) *Repo {
	return &Repo{NewRepo(db)}
}
