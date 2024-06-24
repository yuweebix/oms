package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
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
	ctx context.Context
	db  *pgx.Conn
}

func NewRepository(ctx context.Context, connString string) (*Repository, error) {
	conn, err := pgx.Connect(ctx, connString)
	if err != nil {
		return nil, err
	}
	return &Repository{ctx: ctx, db: conn}, nil
}

func toModelsOrder(so *schemas.Order) (mo *models.Order) {
	mo = &models.Order{
		ID:        uint64(so.ID),
		User:      &models.User{ID: uint64(so.UserID)},
		Expiry:    so.Expiry,
		ReturnBy:  so.ReturnBy,
		Status:    so.Status,
		Hash:      so.Hash,
		CreatedAt: so.CreatedAt,
		Cost:      uint64(so.Cost),
		Weight:    so.Weight,
	}
	return mo
}
