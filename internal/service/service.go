package service

import "ship/internal/repository"

type Service struct{}

func New(repo *repository.Repository) *Service {
	return &Service{}
}
