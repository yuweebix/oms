package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"gitlab.ozon.dev/yuweebix/homework-1/internal/models"
)

// querier интерфейс, исполняющий sql запросы. В качестве возможных куриеров можно использовать пул или же например транзакцию.
type querier interface {
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
}

// Getquerier проверяет, была ли создана транзакция и возвращает её. В противном случае возвращается пул.
func (r *Repository) Getquerier(ctx context.Context) querier {
	if tx, ok := ctx.Value(txKey).(querier); ok && tx != nil {
		return tx
	}
	return r.pool
}

type txKeyType string

const txKey txKeyType = "tx"

// RunTx используется для начала транзакции с предоставленными опциями
func (r *Repository) RunTx(ctx context.Context, opts models.TxOptions, fn func(ctxTX context.Context) error) error {
	pgxOpts := convertTxOptions(opts)
	tx, err := r.pool.BeginTx(ctx, pgxOpts)
	if err != nil {
		return err
	}

	ctxTX := context.WithValue(ctx, txKey, tx)

	if err := fn(ctxTX); err != nil {
		if err := tx.Rollback(ctx); err != nil {
			return err
		}
		return err
	}

	if err := tx.Commit(ctx); err != nil {
		if err := tx.Rollback(ctx); err != nil {
			return err
		}
		return err
	}

	return nil
}

// RunTxWithRollback используется для начала транзакции с предоставленными опциями с откатом
func (r *Repository) RunTxWithRollback(ctx context.Context, opts models.TxOptions, fn func(ctxTX context.Context) error) error {
	pgxOpts := convertTxOptions(opts)
	tx, err := r.pool.BeginTx(ctx, pgxOpts)
	if err != nil {
		return err
	}

	ctxTX := context.WithValue(ctx, txKey, tx)

	if err := fn(ctxTX); err != nil {
		if err := tx.Rollback(ctx); err != nil {
			return err
		}
		return err
	}

	if err := tx.Rollback(ctx); err != nil {
		return err
	}

	return nil
}
