package service

import (
	"context"

	"yourproject/internal/model"
	"yourproject/internal/repository"
	"yourproject/pkg/logger"
)

type Service struct {
	repo   repository.Repository
	logger *logger.Logger
}

func New(repo repository.Repository, logger *logger.Logger) *Service {
	return &Service{
		repo:   repo,
		logger: logger,
	}
}

type ExampleRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (s *Service) GetExamples(ctx context.Context) ([]model.Example, error) {
	return s.repo.GetExamples(ctx)
}

func (s *Service) CreateExample(ctx context.Context, req ExampleRequest) (*model.Example, error) {
	example := &model.Example{
		Name:        req.Name,
		Description: req.Description,
	}

	return s.repo.CreateExample(ctx, example)
}
