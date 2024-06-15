package repository

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/v2/pgxscan"
	"gitlab.ozon.dev/yuweebix/homework-1/internal/models"
	"gitlab.ozon.dev/yuweebix/homework-1/internal/repository/schemas"
	"golang.org/x/exp/maps"
)

// CreateOrder добавляет заказ в бд
func (r *Repository) CreateOrder(o *models.Order) (err error) {

	query := sq.Insert(ordersTable).
		Columns(ordersColumns...).
		Values(o.Unzip()).
		PlaceholderFormat(sq.Dollar)

	rawQuery, args, err := query.ToSql()
	if err != nil {
		return err
	}

	_, err = r.db.Exec(r.ctx, rawQuery, args...)
	if err != nil {
		return err
	}

	return nil
}

// DeleteOrder удаляет заказ из бд
func (r *Repository) DeleteOrder(o *models.Order) (err error) {
	query := sq.Delete(ordersTable).
		Where(sq.Eq{"id": o.ID}).
		PlaceholderFormat(sq.Dollar)

	rawQuery, args, err := query.ToSql()
	if err != nil {
		return err
	}

	_, err = r.db.Exec(r.ctx, rawQuery, args...)
	if err != nil {
		return err
	}

	return nil
}

// GetOrders возвращает список заказов клиента
func (r *Repository) GetOrders(userID int) (list []*models.Order, err error) {
	query := sq.Select(ordersColumns...).
		From(ordersTable).
		Where(sq.Eq{"user_id": userID}).
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

// GetOrdersForDelivery возваращает список заказов клиенту на выдачу
func (r *Repository) GetOrdersForDelivery(orderIDs map[int]struct{}) (list []*models.Order, err error) {
	ids := maps.Keys(orderIDs)

	query := sq.Select(ordersColumns...).
		From(ordersTable).
		Where(sq.Eq{"id": ids}).
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

	for _, o := range orders {
		list = append(list, toModelsOrder(&o))
	}

	return list, nil
}

// GetOrder пересылает полный объект заказа
func (r *Repository) GetOrder(o *models.Order) (result *models.Order, err error) {
	query := sq.Select(ordersColumns...).
		From(ordersTable).
		Where(sq.Eq{"id": o.ID}).
		PlaceholderFormat(sq.Dollar)

	rawQuery, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	order := schemas.Order{}
	err = pgxscan.Get(r.ctx, r.db, &order, rawQuery, args...)
	if err != nil {
		return nil, err
	}

	result = toModelsOrder(&order)
	return result, nil
}
