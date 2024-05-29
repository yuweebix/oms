package module

import (
	"errors"
	"math"

	"gitlab.ozon.dev/yuweebix/homework-1/internal/models"
)

// AcceptReturn принимает заказ от курьера
func (m Module) AcceptReturn(o *models.Order) error {
	// не указали флаг order_id
	if o.ID == -1 {
		return errors.New("no order_id")
	}

	// не указали флаг order_id
	if o.User.ID == -1 {
		return errors.New("no user_id")
	}

	// Запись в хранилище
	return m.Storage.AddReturn(o)
}

// ListReturns выводит список заказов
func (m Module) ListReturns(start, finish int) ([]*models.Order, error) {
	// стандартное значение
	if start < 0 {
		start = 1
	}
	if finish < 0 {
		finish = math.MaxInt
	}

	list, err := m.Storage.ListReturns(start, finish)
	if err != nil {
		return nil, err
	}

	return list, nil
}
