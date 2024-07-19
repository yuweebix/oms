package schemas

import (
	"time"

	"gitlab.ozon.dev/yuweebix/homework-1/internal/models"
)

type Order struct {
	ID        uint64    `redis:"id"`
	UserID    uint64    `redis:"user_id"`
	Expiry    time.Time `redis:"stored_until"`
	ReturnBy  time.Time `redis:"return_by"`
	Status    string    `redis:"status"`
	Hash      string    `redis:"hash"`
	CreatedAt time.Time `redis:"created_at"`
	Cost      uint64    `redis:"cost"`
	Weight    float64   `redis:"weight"`
	Packaging string    `redis:"packaging"`
}

func FromModelsOrder(o *models.Order) *Order {
	return &Order{
		ID:        o.ID,
		UserID:    o.User.ID,
		Expiry:    o.Expiry,
		ReturnBy:  o.ReturnBy,
		Status:    string(o.Status),
		Hash:      o.Hash,
		CreatedAt: o.CreatedAt,
		Cost:      o.Cost,
		Weight:    o.Weight,
		Packaging: string(o.Packaging),
	}
}
