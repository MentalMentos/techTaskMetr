package service

import (
	"github.com/MentalMentos/techTaskMetr/auth/internal/repository"
	"github.com/MentalMentos/techTaskMetr/techTaskmetr/pkg/logger"
)

type Service struct {
	*AuthService
}

func New(repo repository.Repository, logger logger.Logger) *Service {
	return &Service{
		NewAuthService(repo, logger),
	}
}
