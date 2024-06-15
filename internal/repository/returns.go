package repository

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
	"gitlab.ozon.dev/yuweebix/homework-1/internal/models"
	"gitlab.ozon.dev/yuweebix/homework-1/internal/repository/schemas"
)

// GetReturns передает список возвратов
func (r *Repository) GetReturns() (list []*models.Order, err error) {
	// начинаем транзакцию
	tx, err := r.db.BeginTx(r.ctx, pgx.TxOptions{IsoLevel: pgx.ReadCommitted})
	if err != nil {
		return nil, err
	}
	// если закоммититься, то откатить ничего не получится
	defer tx.Rollback(r.ctx)

	// создаем sql запрос
	query := sq.Select(ordersColumns...).
		From(ordersTable).
		Where(sq.Eq{"status": models.StatusReturned}).
		PlaceholderFormat(sq.Dollar)

	// преобразуем в сырой вид
	rawQuery, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	orders := []schemas.Order{}
	err = pgxscan.Select(r.ctx, tx, &orders, rawQuery, args...)
	if err != nil {
		return nil, err
	}

	list = make([]*models.Order, 0, len(orders))
	for _, o := range orders {
		list = append(list, toModelsOrder(&o))
	}

	// коммитим
	err = tx.Commit(r.ctx)
	if err != nil {
		return nil, err
	}

	return list, nil
}
