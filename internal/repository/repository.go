package repository

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"gitlab.ozon.dev/yuweebix/homework-1/internal/models"
	"gitlab.ozon.dev/yuweebix/homework-1/internal/repository/schemas"
)

const (
	ordersTable = "orders"
)

var (
	ordersColumns = []string{"id", "user_id", "stored_until", "return_by", "status", "hash", "created_at", "cost", "weight", "packaging"}
)

type Repository struct {
	pool *pgxpool.Pool
}

func NewRepository(ctx context.Context, connString string) (*Repository, error) {
	config, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, err
	}

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, err
	}

	err = pool.AcquireFunc(ctx, func(conn *pgxpool.Conn) error {
		return conn.Ping(ctx)
	})
	if err != nil {
		pool.Close()
		return nil, err
	}

	return &Repository{pool: pool}, nil
}

func (r *Repository) Close() {
	r.pool.Close()
}

func toModelsOrder(so *schemas.Order) *models.Order {
	return &models.Order{
		ID:        uint64(so.ID),
		User:      &models.User{ID: uint64(so.UserID)},
		Expiry:    so.Expiry,
		ReturnBy:  so.ReturnBy,
		Status:    so.Status,
		Hash:      so.Hash,
		CreatedAt: so.CreatedAt,
		Cost:      uint64(so.Cost),
		Weight:    so.Weight,
		Packaging: models.PackagingType(so.Packaging),
	}
}

func convertTxOptions(opts models.TxOptions) pgx.TxOptions {
	return pgx.TxOptions{
		IsoLevel:   pgx.TxIsoLevel(opts.IsoLevel),
		AccessMode: pgx.TxAccessMode(opts.AccessMode),
	}
}

func (r *Repository) DeleteAllOrders(ctx context.Context) error {
	// создаем sql запрос
	query := squirrel.Delete(ordersTable).PlaceholderFormat(squirrel.Dollar)

	// преобразуем в сырой вид
	rawQuery, args, err := query.ToSql()
	if err != nil {
		return err
	}

	// выполняем запрос
	_, err = r.pool.Exec(ctx, rawQuery, args...)
	return err
}
