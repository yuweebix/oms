package repository

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/v2/pgxscan"
	"gitlab.ozon.dev/yuweebix/homework-1/internal/models"
	"gitlab.ozon.dev/yuweebix/homework-1/internal/repository/schemas"
)

// GetReturns передает список возвратов
func (r *Repository) GetReturns() (list []*models.Order, err error) {
	query := sq.Select(ordersColumns...).
		From(ordersTable).
		Where(sq.Eq{"status": models.StatusReturned}).
		PlaceholderFormat(sq.Dollar)

	rawQuery, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	orders := []schemas.Order{}
	err = pgxscan.Select(r.ctx, r.db, &orders, rawQuery, args...)
	if err != nil {
		return nil, err
	}

	list = make([]*models.Order, 0, len(orders))
	for _, o := range orders {
		list = append(list, toModelsOrder(&o))
	}

	return list, nil
}
