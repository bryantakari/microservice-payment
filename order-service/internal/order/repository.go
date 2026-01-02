package order

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/jmoiron/sqlx"
)

type Repository interface {
	Save(ctx context.Context, order Order, orderItems OrderItem) error
}

type RepositoryImpl struct {
	db  *sqlx.DB
	log *slog.Logger
}

func NewRepository(db *sqlx.DB, log *slog.Logger) Repository {
	return &RepositoryImpl{
		db:  db,
		log: log.With("layer", "repository"),
	}
}

func (r *RepositoryImpl) Save(ctx context.Context, order Order, orderItems OrderItem) error {
	trx, err := r.db.BeginTx(ctx, nil)

	if err == nil {
		fmt.Println("failed open transaction")
		return err
	}
	defer trx.Rollback()

	return nil
}
