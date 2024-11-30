package repository

import (
	"errors"
	"github.com/MentalMentos/techTaskMetr.git/internal/models"
	"github.com/gin-gonic/gin"
)

type Repository interface {
	Create(ctx *gin.Context, m *models.Task) error
}

// Create добавляет нового пользователя в базу данных
func (r *RepoImpl) Create(ctx *gin.Context, m models.Task) error {
	if err := r.DB.WithContext(ctx).Create(&m).Error; err != nil {
		return errors.New("cannot create new user")
	}
	return nil
}
