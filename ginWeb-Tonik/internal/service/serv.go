package service

import (
	"github.com/MentalMentos/ginWeb-Tonik/ginWeb/internal/repository"
	"github.com/MentalMentos/ginWeb-Tonik/ginWeb/pkg/logger"
)

type Service struct {
	*AuthService
}

func New(repo repository.Repository, logger logger.Logger) *Service {
	return &Service{
		NewAuthService(repo, logger),
	}
}
