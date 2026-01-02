package order

import (
	"context"
	"log/slog"
)

type Service interface {
	createOrder(ctx context.Context) error
}

type ServiceImpl struct {
	r   Repository
	log *slog.Logger
}

func NewService(r Repository, log *slog.Logger) Service {
	return &ServiceImpl{r: r, log: log.With("layer", "service")}
}

func (srv *ServiceImpl) createOrder(ctx context.Context) error {

	return nil
}
