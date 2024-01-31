package service

import (
	"context"
	"log/slog"
	"sample/go-react-local-app/internal/models"
	"sample/go-react-local-app/internal/repository"
	"sample/go-react-local-app/internal/transaction"
)

type CountServicer interface {
	Set(context.Context, int) error
	Get(context.Context, int) (*models.Count, error)
}

type countService struct {
	repo repository.Counter
	tx   transaction.Transaction
	log  *slog.Logger
}

func NewCountSerivce(repo repository.Counter, tx transaction.Transaction, logger *slog.Logger) CountServicer {
	return &countService{repo: repo, tx: tx, log: logger}
}

func (cs *countService) Set(ctx context.Context, count int) error {
	err := cs.tx.DoTx(ctx, func(ctx context.Context) error {
		return cs.repo.Add(ctx, models.CountValue(count))
	})
	if err != nil {
		return err
	}
	return nil
}

func (cs *countService) Get(ctx context.Context, id int) (*models.Count, error) {
	return cs.repo.FindById(ctx, models.CountId(id))
}
