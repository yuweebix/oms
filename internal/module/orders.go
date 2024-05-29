package module

import (
	"errors"

	"gitlab.ozon.dev/yuweebix/homework-1/internal/models"
)

// AcceptOrder принимает заказ от курьера
func (m Module) AcceptOrder(o *models.Order) error {
	// не указали флаг order_id
	if o.ID == -1 {
		return errors.New("no order_id")
	}

	// не указали флаг order_id
	if o.User.ID == -1 {
		return errors.New("no user_id")
	}

	// Запись в хранилище
	return m.Storage.AddOrder(o)
}

// AcceptOrder возвращает заказ курьеру
func (m Module) ReturnOrder(o *models.Order) error {
	// не указали флаг order_id
	if o.ID == -1 {
		return errors.New("no order_id")
	}

	// Запись в хранилище
	return m.Storage.DeleteOrder(o)
}

// ListOrders выводит список заказов
func (m Module) ListOrders(limit int) ([]*models.Order, error) {
	// стандартное значение
	if limit < 1 {
		limit = 10
	}

	list, err := m.Storage.ListOrders(limit)
	if err != nil {
		return nil, err
	}

	return list, nil
}

// DeliverOrders принимает список заказов, переводит их в форму для обработки в хранилище
func (m Module) DeliverOrders(oIDSlice []int) error {
	// создаем структуру сет, для быстрого поиска
	oSet := make(map[int]struct{}, len(oIDSlice))
	for _, v := range oIDSlice {
		oSet[v] = struct{}{}
	}

	return m.Storage.CheckOrdersForDelivery(oSet)
}
