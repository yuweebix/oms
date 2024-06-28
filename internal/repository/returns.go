package repository

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/v2/pgxscan"
	"gitlab.ozon.dev/yuweebix/homework-1/internal/models"
	"gitlab.ozon.dev/yuweebix/homework-1/internal/repository/schemas"
)

// GetReturns возвращает список возвратов
func (r *Repository) GetReturns(ctx context.Context, limit uint64, offset uint64) (list []*models.Order, err error) {
	qr := r.GetQuerier(ctx)

	// создаем sql запрос
	query := sq.Select(ordersColumns...).
		From(ordersTable).
		Where(sq.Eq{"status": models.StatusReturned}).
		Limit(limit).
		Offset(offset).
		PlaceholderFormat(sq.Dollar)

	// преобразуем в сырой вид
	rawQuery, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	orders := []schemas.Order{}
	err = pgxscan.Select(ctx, qr, &orders, rawQuery, args...)
	if err != nil {
		return nil, err
	}

	list = make([]*models.Order, 0, len(orders))
	for _, o := range orders {
		list = append(list, toModelsOrder(&o))
	}

	return list, nil
}
