package schemas

import (
	"time"

	"gitlab.ozon.dev/yuweebix/homework-1/internal/models"
)

type Order struct {
	ID        int           `db:"order_id"`
	UserID    int           `db:"user_id"`
	Expiry    time.Time     `db:"stored_until"`
	ReturnBy  time.Time     `db:"return_by"`
	Status    models.Status `db:"status"`
	Hash      string        `db:"hash"`
	CreatedAt time.Time     `db:"created_at"`
}
