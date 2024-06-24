package schemas

import (
	"time"

	"gitlab.ozon.dev/yuweebix/homework-1/internal/models"
)

type Order struct {
	ID        uint64        `db:"id"`
	UserID    uint64        `db:"user_id"`
	Expiry    time.Time     `db:"stored_until"`
	ReturnBy  time.Time     `db:"return_by"`
	Status    models.Status `db:"status"`
	Hash      string        `db:"hash"`
	CreatedAt time.Time     `db:"created_at"`
	Cost      float64       `db:"cost"`
	Weight    float64       `db:"weight"`
	Packaging string        `db:"packaging"`
}
