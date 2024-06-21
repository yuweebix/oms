package repository

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
	"gitlab.ozon.dev/yuweebix/homework-1/internal/models"
	"gitlab.ozon.dev/yuweebix/homework-1/internal/repository/schemas"
)

// GetPackaging пересылает полный объект упаковки
func (r *Repository) GetPackaging(p *models.Packaging) (result *models.Packaging, err error) {
	// начинаем транзакцию
	tx, err := r.db.BeginTx(r.ctx, pgx.TxOptions{IsoLevel: pgx.ReadCommitted})
	if err != nil {
		return nil, err
	}
	// если закоммититься, то откатить ничего не получится
	defer tx.Rollback(r.ctx)

	// создаем sql запрос
	query := sq.Select(packagingColumns...).
		From(packagingTable).
		Where(sq.Eq{"type": p.Type}).
		PlaceholderFormat(sq.Dollar)

	// преобразуем в сырой вид
	rawQuery, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	packaging := schemas.Packaging{}
	err = pgxscan.Get(r.ctx, tx, &packaging, rawQuery, args...)
	if err != nil {
		return nil, err
	}

	result = toModelsPackaging(&packaging)

	// коммитим
	err = tx.Commit(r.ctx)
	if err != nil {
		return nil, err
	}

	return result, nil
}
