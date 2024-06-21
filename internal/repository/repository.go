package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
	"gitlab.ozon.dev/yuweebix/homework-1/internal/models"
	"gitlab.ozon.dev/yuweebix/homework-1/internal/repository/schemas"
)

const (
	ordersTable    = "orders"
	packagingTable = "packaging"
)

var (
	ordersColumns    = []string{"id", "user_id", "stored_until", "return_by", "status", "hash", "created_at", "cost", "weight", "packaging"}
	packagingColumns = []string{"type", "cost", "weight_limit"}
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
		ID:        so.ID,
		User:      &models.User{ID: so.UserID},
		Expiry:    so.Expiry,
		ReturnBy:  so.ReturnBy,
		Status:    so.Status,
		Hash:      so.Hash,
		CreatedAt: so.CreatedAt,
		Cost:      so.Cost,
		Weight:    so.Weight,
	}
	return mo
}

func toModelsPackaging(sp *schemas.Packaging) (mp *models.Packaging) {
	mp = &models.Packaging{
		Type:        sp.Type,
		Cost:        sp.Cost,
		WeightLimit: sp.WeightLimit,
	}
	return mp
}
