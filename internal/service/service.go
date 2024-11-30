package service

import "github.com/MentalMentos/techTaskMetr.git/internal/repository"

type Service struct {
	repo repository.Repository
}

func NewService(repo repository.Repository) *Service {
	return &Service{repo: repo}
}
