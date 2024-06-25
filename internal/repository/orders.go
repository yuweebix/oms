package repository

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/v2/pgxscan"
	"gitlab.ozon.dev/yuweebix/homework-1/internal/models"
	e "gitlab.ozon.dev/yuweebix/homework-1/internal/repository/errors"
	"gitlab.ozon.dev/yuweebix/homework-1/internal/repository/schemas"
)

// CreateOrder добавляет заказ в бд
func (r *Repository) CreateOrder(tx models.Tx, o *models.Order) (err error) {
	// создаем sql запрос
	query := sq.Insert(ordersTable).
		Columns(ordersColumns...).
		Values(o.Unzip()).
		PlaceholderFormat(sq.Dollar)

	// преобразуем в сырой вид
	rawQuery, args, err := query.ToSql()
	if err != nil {
		return err
	}

	res, err := tx.Exec(r.ctx, rawQuery, args...)
	if err != nil {
		return err
	}

	// проверка на то, изменилось ли что-то вообще
	if res.RowsAffected() > 1 {
		return e.ErrTooManyOrdersAffected
	} else if res.RowsAffected() == 0 {
		return e.ErrNoOrdersAffected
	}

	return nil
}

// DeleteOrder удаляет заказ из бд
func (r *Repository) DeleteOrder(tx models.Tx, o *models.Order) (err error) {
	// создаем sql запрос
	query := sq.Delete(ordersTable).
		Where(sq.Eq{"id": o.ID}).
		PlaceholderFormat(sq.Dollar)

	// преобразуем в сырой вид
	rawQuery, args, err := query.ToSql()
	if err != nil {
		return err
	}

	res, err := tx.Exec(r.ctx, rawQuery, args...)
	if err != nil {
		return err
	}

	// проверка на то, изменилось ли что-то вообще
	if res.RowsAffected() > 1 {
		return e.ErrTooManyOrdersAffected
	} else if res.RowsAffected() == 0 {
		return e.ErrNoOrdersAffected
	}

	return nil
}

// UpdateOrder обновляет данные заказа в бд
func (r *Repository) UpdateOrder(tx models.Tx, o *models.Order) (err error) {
	// создаем sql запрос
	query := sq.Update(ordersTable).
		Set("user_id", o.User.ID).
		Set("stored_until", o.Expiry).
		Set("return_by", o.ReturnBy).
		Set("status", o.Status).
		Set("hash", o.Hash).
		Set("created_at", o.CreatedAt).
		Where(sq.Eq{"id": o.ID}).
		PlaceholderFormat(sq.Dollar)

	// преобразуем в сырой вид
	rawQuery, args, err := query.ToSql()
	if err != nil {
		return err
	}

	res, err := tx.Exec(r.ctx, rawQuery, args...)
	if err != nil {
		return err
	}

	// проверка на то, изменилось ли что-то вообще
	if res.RowsAffected() > 1 {
		return e.ErrTooManyOrdersAffected
	} else if res.RowsAffected() == 0 {
		return e.ErrNoOrdersAffected
	}

	return nil
}

// GetOrders возвращает список заказов клиента
func (r *Repository) GetOrders(tx models.Tx, userID uint64, limit uint64, offset uint64, isStored bool) (list []*models.Order, err error) {
	// создаем sql запрос
	query := sq.Select(ordersColumns...).
		From(ordersTable).
		Where(sq.Eq{"user_id": userID}).
		OrderBy("created_at DESC") // сортировка по времени создания в убывающем порядке

	// применяем фильтрацию по isStored, если необходимо
	if isStored {
		query = query.Where(sq.Or{
			sq.Eq{"status": models.StatusAccepted},
			sq.Eq{"status": models.StatusReturned},
		})
	}

	// добавляем Limit и Offset
	query = query.Limit(limit).
		Offset(offset).
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

	return list, nil
}

// GetOrdersForDelivery возвращает список заказов клиенту на выдачу
func (r *Repository) GetOrdersForDelivery(tx models.Tx, orderIDs []uint64) (list []*models.Order, err error) {
	// создаем sql запрос
	query := sq.Select(ordersColumns...).
		From(ordersTable).
		Where(sq.Eq{"id": orderIDs}).
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

	for _, o := range orders {
		list = append(list, toModelsOrder(&o))
	}

	return list, nil
}

// GetOrder пересылает полный объект заказа
func (r *Repository) GetOrder(tx models.Tx, o *models.Order) (result *models.Order, err error) {
	// создаем sql запрос
	query := sq.Select(ordersColumns...).
		From(ordersTable).
		Where(sq.Eq{"id": o.ID}).
		PlaceholderFormat(sq.Dollar)

	// преобразуем в сырой вид
	rawQuery, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	order := schemas.Order{}
	err = pgxscan.Get(r.ctx, tx, &order, rawQuery, args...)
	if err != nil {
		return nil, err
	}

	result = toModelsOrder(&order)

	return result, nil
}
