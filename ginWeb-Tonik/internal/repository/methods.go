package repository

import (
	"context"
	"errors"
	"github.com/MentalMentos/ginWeb-Tonik/ginWeb/internal/data/request"
	"github.com/MentalMentos/ginWeb-Tonik/ginWeb/pkg/logger"

	"github.com/MentalMentos/ginWeb-Tonik/ginWeb/internal/model"
	"gorm.io/gorm"
)

type RepoImpl struct {
	DB *gorm.DB
}

func NewRepo(db *gorm.DB) *RepoImpl {
	return &RepoImpl{
		db,
	}
}

func (r *RepoImpl) Create(ctx context.Context, us model.User, logger logger.Logger) (int64, error) {
	if err := r.DB.WithContext(ctx).Create(&us).Error; err != nil {
		logger.Fatal("[ REPO_CREATE ]", err.Error())
		return 0, errors.New("cannot create new user")
	}
	logger.Info("[ REPO_CREATE ]", "user created successful")
	return us.ID, nil
}

func (r *RepoImpl) Update(ctx context.Context, us model.User, logger logger.Logger) (int64, error) {
	updateData := request.UpdateUserRequest{
		Name:  us.Name,
		Email: us.Email,
	}

	if err := r.DB.WithContext(ctx).Model(&model.User{}).Where("id = ?", us.ID).Updates(updateData).Error; err != nil {
		logger.Fatal("[ REPO_UPDATE ]", err.Error())
		return 0, errors.New("cannot update user")
	}
	logger.Info("[ REPO_UPDATE ]", "user's email/name updated successful")
	return us.ID, nil
}

func (r *RepoImpl) Delete(ctx context.Context, usId int64, logger logger.Logger) error {
	if err := r.DB.WithContext(ctx).Delete(&model.User{}, usId).Error; err != nil {
		logger.Fatal("[ REPO_DELETE ]", err.Error())
		return errors.New("cannot delete user")
	}
	logger.Info("[ REPO_DELETE ]", "user deleted successful")
	return nil
}

func (r *RepoImpl) UpdatePassword(ctx context.Context, us model.User, hashPassword string, logger logger.Logger) (model.User, error) {
	updateUser := request.UpdateUserRequest{
		Name:     us.Name,
		Email:    us.Email,
		Password: hashPassword,
	}
	if err := r.DB.WithContext(ctx).Updates(&updateUser).Error; err != nil {
		logger.Fatal("[ REPO_UPDATE_PASSWORD ]", err.Error())
		return model.User{}, errors.New("cannot update password")
	}
	logger.Fatal("[ REPO_UPDATE_PASSWORD ]", "password updated successful")
	return us, nil
}

func (r *RepoImpl) GetByEmail(ctx context.Context, email string, logger logger.Logger) (model.User, error) {
	var user model.User
	if err := r.DB.WithContext(ctx).Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return user, errors.New("user not found")
		}
		logger.Fatal("[ REPO_GET_BY_EMAIL ]", err.Error())
		return user, err
	}
	logger.Info("[ REPO_GET_BY_EMAIL ]", "get by email successful")
	return user, nil
}

func (r *RepoImpl) GetByID(ctx context.Context, userID int64, logger logger.Logger) (model.User, error) {
	var user model.User
	if err := r.DB.WithContext(ctx).First(&user, userID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return user, errors.New("user not found")
		}
		logger.Fatal("[ REPO_GET_BY_ID ]", err.Error())
		return user, err
	}
	logger.Info("[ REPO_GET_BY_ID ]", "get by id successful")
	return user, nil
}

func (r *RepoImpl) GetAll(ctx context.Context, logger logger.Logger) ([]model.User, error) {
	var users []model.User
	if err := r.DB.WithContext(ctx).Find(&users).Error; err != nil {
		logger.Fatal("[ REPO_ALL ]", err.Error())
		return nil, errors.New("users not found")
	}
	logger.Info("[ REPO_ALL ]", "list all successful")
	return users, nil
}

func (r *RepoImpl) UpdateIP(ctx context.Context, us model.User, ip string, logger logger.Logger) (model.User, error) {
	updateUser := request.UpdateUserRequest{
		Name:     us.Name,
		Email:    us.Email,
		Password: us.Password,
		IP:       ip,
	}

	if err := r.DB.WithContext(ctx).Updates(&updateUser).Error; err != nil {
		logger.Fatal("[ REPO_UPDATE_IP ]", err.Error())
		return us, errors.New("cannot update password")
	}
	logger.Info("[ REPO_UPD_IP ]", "ip updated successful")
	return us, nil
}
