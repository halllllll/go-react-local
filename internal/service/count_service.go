package service

import (
	"context"
	"log/slog"
	"sample/go-react-local-app/internal/models"
	"sample/go-react-local-app/internal/repository"
)

type CountServicer interface {
	Set(context.Context, int) error
	Get(context.Context, int) (*models.Count, error)
}

type countService struct {
	repo   repository.Counter
	logger *slog.Logger
}

func NewCountSerivce(repo repository.Counter, logger *slog.Logger) CountServicer {
	return &countService{repo: repo, logger: logger}
}

func (cs *countService) Set(ctx context.Context, count int) error {
	return cs.repo.Add(ctx, models.CountValue(count))
}

func (cs *countService) Get(ctx context.Context, id int) (*models.Count, error) {
	return cs.repo.FindById(ctx, models.CountId(id))
}
