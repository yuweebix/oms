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
	ordersColumns = []string{"id", "user_id", "stored_until", "return_by", "status", "hash", "created_at"}
)

type Repository struct {
	ctx context.Context
	db  *pgx.Conn
}

func NewRepository(ctx context.Context, db *pgx.Conn) *Repository {
	return &Repository{ctx: ctx, db: db}
}

func toModelsOrder(so *schemas.Order) (mo *models.Order) {
	mo = &models.Order{
		ID:        so.ID,
		User:      &models.User{ID: *so.UserID},
		Expiry:    so.Expiry,
		ReturnBy:  so.ReturnBy,
		Status:    so.Status,
		Hash:      so.Hash,
		CreatedAt: so.CreatedAt,
	}
	return mo
}
