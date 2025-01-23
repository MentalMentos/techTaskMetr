package service

import (
	"context"
	"errors"
	"github.com/MentalMentos/techTaskMetr/techTaskmetr/auth/internal/data/request"
	"github.com/MentalMentos/techTaskMetr/techTaskmetr/auth/internal/data/response"
	"github.com/MentalMentos/techTaskMetr/techTaskmetr/auth/internal/model"
	"github.com/MentalMentos/techTaskMetr/techTaskmetr/auth/internal/repository"
	"github.com/MentalMentos/techTaskMetr/techTaskmetr/auth/pkg/helpers"
	"github.com/MentalMentos/techTaskMetr/techTaskmetr/pkg/logger"
	"github.com/MentalMentos/techTaskMetr/techTaskmetr/pkg/utils"

	"golang.org/x/crypto/bcrypt"
)

type Auth interface {
	Register(ctx context.Context, req request.RegisterUserRequest) (*response.AuthResponse, error)
	Login(ctx context.Context, req request.LoginRequest) (*response.AuthResponse, error)
	RefreshToken(ctx context.Context, refreshToken string) (*response.AuthResponse, error)
	UpdatePassword(ctx context.Context, req request.UpdateUserRequest) (*response.AuthResponse, error)
}

type AuthService struct {
	repo   repository.Repository
	logger logger.Logger
}

func NewAuthService(repo repository.Repository, logger logger.Logger) *AuthService {
	return &AuthService{
		repo,
		logger,
	}
}

func (s *AuthService) Register(ctx context.Context, req request.RegisterUserRequest) (*response.AuthResponse, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		s.logger.Fatal("[ SERVICE_REGISTER ]", helpers.FailedToHashPass)
		return nil, err
	}

	user := model.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: string(hashedPassword),
		Role:     "user",
		IP:       req.IP,
	}
	if user.Name == "" {
		s.logger.Fatal("[ SERVICE_REGISTER ]", "null name")
	} else if user.Email == "" {
		s.logger.Fatal("[ SERVICE_REGISTER ]", "null email")
	}

	user_id, err := s.repo.Create(ctx, user, s.logger)
	if err != nil {
		s.logger.Fatal("[ SERVICE_REGISTER ]", helpers.FailedToCreateUser)
		return nil, err
	}

	accessToken, refreshToken, err := utils.GenerateJWT(user.ID, user.Role)
	if err != nil {
		s.logger.Fatal("[ SERVICE_REGISTER ]", helpers.FailedToGenJWT)
		return nil, err
	}

	return &response.AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		UserID:       user_id,
	}, nil
}

func (s *AuthService) Login(ctx context.Context, req request.LoginRequest) (*response.AuthResponse, error) {
	user, err := s.repo.GetByEmail(ctx, req.Email, s.logger)
	if err != nil {
		s.logger.Fatal("[ SERVICE_LOGIN ]", helpers.FailedToGetUser)
		return nil, errors.New("user not found")
	}

	if user.IP != req.IP {
		_, err := s.repo.UpdateIP(ctx, user, req.IP, s.logger)
		if err != nil {
			s.logger.Fatal("[ SERVICE_LOGIN ]", "failed to update ip")
			return nil, errors.New("cannot update ip with login")
		}
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		s.logger.Fatal("[ SERVICE_LOGIN ]", helpers.FailedToHashPass)
		return nil, errors.New("invalid password")
	}

	accessToken, refreshToken, err := utils.GenerateJWT(user.ID, user.Role)
	if err != nil {
		s.logger.Fatal("[ SERVICE_LOGIN ]", helpers.FailedToGenJWT)
		return nil, err
	}

	return &response.AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		UserID:       user.ID,
	}, nil
}

func (s *AuthService) UpdatePassword(ctx context.Context, req request.UpdateUserRequest) (*response.AuthResponse, error) {
	user, err := s.repo.GetByEmail(ctx, req.Email, s.logger)
	if err != nil {
		s.logger.Fatal("[ SERVICE_UPDATE_PASSWORD ]", helpers.FailedToGetUser)
		return nil, errors.New("user not found")
	}

	if user.IP != req.IP {
		_, err := s.repo.UpdateIP(ctx, user, req.IP, s.logger)
		if err != nil {
			s.logger.Fatal("[ SERVICE_UPDATE_PASSWORD ]", "failed to update ip with login")
			return nil, errors.New("cannot update ip with login")
		}
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		s.logger.Fatal("[ SERVICE_UPDATE_PASSWORD ]", "invalid password")
		return nil, errors.New("invalid password")
	}
	accessToken, refreshToken, err := utils.GenerateJWT(user.ID, user.Role)
	if err != nil {
		s.logger.Fatal("[ SERVICE_UPDATE_PASSWORD ]", "failed to generate access token")
		return nil, err
	}
	return &response.AuthResponse{
		accessToken,
		refreshToken,
		user.ID,
	}, nil
}

func (s *AuthService) RefreshToken(ctx context.Context, refreshToken string) (*response.AuthResponse, error) {
	// Валидация refresh token
	claims, err := utils.ValidateJWT(refreshToken)
	if err != nil {
		s.logger.Fatal("[ SERVICE_GET_ACCESS_TOKEN ]", "failed to validate tokens")
		return nil, errors.New("invalid refresh token")
	}

	// Генерация нового набора токенов
	newAccessToken, newRefreshToken, err := utils.GenerateJWT(claims.UserID, claims.Role)
	if err != nil {
		s.logger.Fatal("[ SERVICE_GET_ACCESS_TOKEN ]", "failed to generate access tokens")
		return nil, err
	}

	return &response.AuthResponse{
		AccessToken:  newAccessToken,
		RefreshToken: newRefreshToken,
	}, nil
}
